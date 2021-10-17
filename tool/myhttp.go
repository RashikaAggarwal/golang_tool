package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Function to make Http request and prints the respective response
func makeHttpRequest(req <-chan string, res chan<- bool) {
	for url := range req {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}
		text := string(body)

		//Generating md5 hash of the response
		hash := fmt.Sprintf("%x", md5.Sum([]byte(text)))
		fmt.Println(url + " " + hash)
		defer resp.Body.Close()

		//Sending message over the channel
		res <- true
	}
}

//Function to begin code execution
func main() {
	var parallel int
	flag.IntVar(&parallel, "parallel", 10, "Number of parallel requests to be executed")
	flag.Parse()
	args := flag.Args()

	request := make(chan string, len(args))
	response := make(chan bool, len(args))

	if parallel > len(args) {
		parallel = len(args)
	}

	//Requests to be executed via go routines
	for p := 0; p < parallel; p++ {
		go makeHttpRequest(request, response)
	}

	//Sending message over the channel
	for _, url := range args {
		request <- url
	}
	close(request)

	//Recieving message over the channel
	for i := 0; i < len(args); i++ {
		<-response
	}
}
