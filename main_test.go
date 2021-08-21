package main

import (
	"fmt"
	"reflect"
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
		{"GET https://meain.io", requestObject{url: "https://meain.io", method: "GET"}},
		{"POST https://meain.io", requestObject{url: "https://meain.io", method: "POST"}},
		{"POST meain.io", requestObject{url: "meain.io", method: "POST"}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Fatalf("parsing failed")
			}
			if got.url != tc.want.url || got.method != tc.want.method {
				t.Fatalf("got %v:%v; want %v:%v", got.url, got.method, tc.want.url, tc.want.method)
			} else {
				t.Logf("Success !")
			}

		})
	}
}

func TestInputParseWithComments(t *testing.T) {
	tests := []struct {
		input string
		want  requestObject
	}{
		{"# Sample comment\nGET https://meain.io", requestObject{url: "https://meain.io", method: "GET"}},
		{"# Sample comment\n#More comments\nGET https://meain.io", requestObject{url: "https://meain.io", method: "GET"}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Fatalf("parsing failed")
			}
			if got.url != tc.want.url || got.method != tc.want.method {
				t.Fatalf("got %v:%v; want %v:%v", got.url, got.method, tc.want.url, tc.want.method)
			} else {
				t.Logf("Success !")
			}

		})
	}
}

func TestInputParseWithHeaders(t *testing.T) {
	tests := []struct {
		input string
		want  requestObject
	}{
		{"GET https://meain.io\nContent-Type: application/json", requestObject{url: "https://meain.io", method: "GET", headers: map[string]string{"Content-Type": "application/json"}}},
		{"GET https://meain.io\nContent-Type: application/json\nKeep-Alive:300", requestObject{url: "https://meain.io", method: "GET", headers: map[string]string{"Content-Type": "application/json", "Keep-Alive": "300"}}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Fatalf("parsing failed")
			}
			if !reflect.DeepEqual(got.headers, tc.want.headers) {
				t.Fatalf("got %v; want %v", got.headers, tc.want.headers)
			} else {
				t.Logf("Success !")
			}

		})
	}
}

func TestInputParseWithData(t *testing.T) {
	tests := []struct {
		input string
		want  requestObject
	}{
		{"GET https://meain.io\n\nHello World!", requestObject{url: "https://meain.io", method: "GET", data: "Hello World!"}},
		{"GET https://meain.io\n\n{\n\"key\":\"value\"\n}", requestObject{url: "https://meain.io", method: "GET", data: "{\n\"key\":\"value\"\n}"}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("InputParse=%d", i), func(t *testing.T) {
			got, e := parseInput(tc.input)
			if e != nil {
				t.Fatalf("parsing failed")
			}
			if got.data != tc.want.data {
				t.Fatalf("got %v; want %v", got.data, tc.want.data)
			} else {
				t.Logf("Success !")
			}

		})
	}
}
