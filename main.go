package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type parseState int

const (
	Url parseState = iota
	Headers
	Data
)

type requestObject struct {
	url     string
	method  string
	headers map[string]string
	data    string
}

func isIn(item string, items []string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func parseInput(input string) (requestObject, error) {
	lines := strings.Split(input, "\n")
	currentParseState := Url
	var parsedMethod string
	var parsedURL string
	parsedHeaders := make(map[string]string)
	var parsedData string
	for _, line := range lines {
		if len(line) == 0 {
			if currentParseState == Headers {
				currentParseState = Data
			} else {
				continue
			}
		} else if strings.HasPrefix(strings.Trim(line, " "), "#") {
			// This is a comment
			continue
		} else if currentParseState == Url {
			tokens := strings.Split(line, " ")
			if isIn(tokens[0], []string{"GET", "POST", "DELETE", "HEAD", "PUT", "PATCH"}) {
				parsedMethod = tokens[0]
			} else {
				continue
			}
			currentParseState = Headers
			parsedURL = tokens[1]
		} else if currentParseState == Headers {
			tokens := strings.Split(line, ":")
			if len(tokens) == 2 {
				parsedHeaders[strings.Trim(tokens[0], " ")] = strings.Trim(tokens[1], " ")
			} else {
				fmt.Println("Unable to parse: ", line)
			}
		} else if currentParseState == Data {
			parsedData += line + "\n"
		}
	}
	if parsedURL != "" {
		rO := requestObject{url: parsedURL, method: parsedMethod}
		if len(parsedHeaders) > 0 {
			rO.headers = parsedHeaders
		}
		if len(parsedData) > 0 {
			// Remove the last \n that was added
			parsedData = parsedData[:len(parsedData)-1]
			rO.data = parsedData
		}
		return rO, nil
	} else {
		return requestObject{}, errors.New("Unable to parse the thing")

	}
}

func makeRequest(rO requestObject) {
	req, err := http.NewRequest(rO.method, rO.url, strings.NewReader(rO.data))
	if err != nil {
		fmt.Println("Unable to create request")
	}
	for header, value := range rO.headers {
		req.Header.Add(header, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to process request")
	}
	fmt.Println(resp.Status)
	for header, value := range resp.Header {
		fmt.Println(header+":", value)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unable to read resopnse body")
	}
	contentType, ok := resp.Header["Content-Type"]
	if ok && strings.Split(contentType[0], ";")[0] == "application/json" {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, body, "", "  ")
		if err != nil {
			fmt.Println("\n", string(body))
		} else {
			fmt.Println("\n", string(prettyJSON.Bytes()))
		}
	} else {
		fmt.Println("\n", string(body))
	}
}

func main() {
	var input string
	if len(os.Args) > 1 {
		// We have a filename passed in
		fileName := os.Args[1]
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Unable to read file")
		}
		input = string(content)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = input + scanner.Text() + "\n"
		}
	}
	rm, err := parseInput(input)
	if err != nil {
		fmt.Println("Unable to parse input")
		return
	}
	makeRequest(rm)
}
