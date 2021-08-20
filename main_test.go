package main

import (
	"fmt"
	"testing"
)

func TestInputParseInvalid(t *testing.T) {
	tests := []struct {
		input string
		want  requestObject
	}{
		{"", requestObject{}},
		{"GE https://meain.io", requestObject{}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Logf("Parsing was supposed to fail and it failed")
			} else {
				t.Fatalf("got %v; want nothing", got)
			}

		})
	}
}

func TestInputParseBasic(t *testing.T) {
	tests := []struct {
		input string
		want  requestObject
	}{
		{"GET https://meain.io", requestObject{url: "https://meain.io", method: GET}},
		{"POST https://meain.io", requestObject{url: "https://meain.io", method: POST}},
		{"POST meain.io", requestObject{url: "meain.io", method: POST}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Fatalf("parsing failed")
			}
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}

		})
	}
}
