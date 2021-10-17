package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//Function to test Main method
func TestMainProgram(t *testing.T) {
	cases := []struct {
		Name string
		Args []string
	}{
		{"flags set", []string{"", "-parallel", "1", "google.com"}},
		{"flags set with default value", []string{"", "google.com", "facebook.com"}},
	}
	for _, tc := range cases {
		flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
		os.Args = tc.Args
		main()
	}
}

//Function to test makeHttpRequest method
func TestMakeHttpRequest(t *testing.T) {
	expected := "dummy data"

	//Test Server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	req := make(chan string)
	res := make(chan bool)
	go makeHttpRequest(req, res)
	req <- svr.URL
	output := <-res
	if output != true {
		t.Errorf("Unexpected Result, got %t, wanted %t", output, true)
	} else {
		t.Log("Test passed successfully for input", svr.URL)
	}
}
