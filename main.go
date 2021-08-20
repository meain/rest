package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/imroc/req"
)

type requestMethod int

const (
	GET requestMethod = iota
	POST
)

type requestObject struct {
	url    string
	method requestMethod
}

func parseInput(input string) (requestObject, error) {
	lines := strings.Split(input, "\n")
	parseStarted := false
	var parsedMethod requestMethod
	var parsedURL string
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if !parseStarted {
			if len(line) == 0 {
				continue
			} else if tokens[0] == "GET" {
				parsedMethod = GET
				parsedURL = tokens[1]
			} else if tokens[0] == "POST" {
				parsedMethod = POST
				parsedURL = tokens[1]
			}
		}
	}
	if parsedURL != "" {
		return requestObject{url: parsedURL, method: parsedMethod}, nil
	} else {
		return requestObject{}, errors.New("Unable to parse the thing")
	}
}

func makeRequest(rO requestObject) {
	r, err := req.Get(rO.url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.String())
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
			fmt.Println("input:", input)
		}
	}
	fmt.Println("input:", input)
	rm, err := parseInput(input)
	if err != nil {
		fmt.Println("Unable to parse input")
		return
	}
	makeRequest(rm)
}
