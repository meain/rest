package main

import (
	"bufio"
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

func parseInput(input string) requestObject {
	lines := strings.Split(input, "\n")
	var filteredLines []string
	for _, line := range lines {
		if len(line) != 0 {
			filteredLines = append(filteredLines, line)
		}
	}
	fmt.Println("filteredLines:", filteredLines)
	firstLine := filteredLines[0] // TODO: do bounds check
	tokens := strings.Split(firstLine, " ")
	var rM requestMethod = GET
	if tokens[0] == "GET" {
		rM = GET
	} else if tokens[0] == "POST" {
		rM = POST
	} else {
		panic("Dude, what the hell is going on?")
	}
	return requestObject{url: tokens[1], method: rM}
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
	rm := parseInput(input)
	makeRequest(rm)
}
