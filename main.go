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
	stateUrl parseState = iota
	stateHeaders
	stateData
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
	currentParseState := stateUrl
	var parsedMethod string
	var parsedURL string
	parsedHeaders := make(map[string]string)
	var parsedData string
	for _, line := range lines {
		if currentParseState == stateUrl {
			if len(line) == 0 || strings.HasPrefix(strings.Trim(line, " "), "#") {
				continue
			}
			tokens := strings.Split(line, " ")
			if isIn(tokens[0], []string{"GET", "POST", "DELETE", "HEAD", "PUT", "PATCH"}) {
				parsedMethod = tokens[0]
			} else {
				continue
			}
			currentParseState = stateHeaders
			parsedURL = tokens[1]
		} else if currentParseState == stateHeaders {
			if len(line) == 0 {
				currentParseState = stateData
				continue
			}
			tokens := strings.Split(line, ":")
			if len(tokens) == 2 {
				parsedHeaders[strings.Trim(tokens[0], " ")] = strings.Trim(tokens[1], " ")
			} else {
				fmt.Println("Unable to parse:", line)
			}
		} else if currentParseState == stateData {
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
		if parsedMethod == "" {
			return requestObject{}, errors.New("No method name found")
		}
		return requestObject{}, errors.New("Unable to parse the thing")

	}
}

func makeRequest(rO requestObject) error {
	req, err := http.NewRequest(rO.method, rO.url, strings.NewReader(rO.data))
	if err != nil {
		return errors.New("Unable to create request")
	}
	for header, value := range rO.headers {
		req.Header.Add(header, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Unable to process request")
	}
	fmt.Println(resp.Status)
	for header, value := range resp.Header {
		fmt.Println(header+":", value)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Unable to read resopnse body")
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
	return nil
}

func main() {
	var input string
	if len(os.Args) > 1 {
		// We have a filename passed in
		fileName := os.Args[1]
		if fileName == "--help" {
			fmt.Println("Usage: rest <filename>")
			return
		} else {
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println("Unable to read file")
				return
			}
			input = string(content)
		}
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
	err = makeRequest(rm)
	if err != nil {
		fmt.Println("Unable to make request:", err)
	}
}
