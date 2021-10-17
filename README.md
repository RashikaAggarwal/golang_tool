# Golang Tool

This repository is a golang tool which makes http requests to the external server and prints the address of the request along with the MD5 hash of the response. 

## Prerequisites

Before you start, please make sure you have Go installed in your system. If not, please use the following link to install Golang:
https://golang.org/doc/install

## Getting Started

Clone the git repository in your system and then cd into project root directory

```bash
$ git clone https://github.com/RashikaAggarwal/golang_tool.git
$ cd golang_tool
```

Build your tool by executing the following steps
```bash
$ cd tool
$ go build
```

## Sample Outputs

This tool takes address of the http requests as command line arguments. See below examples
```bash
$ ./myhttp.exe google.com
http://google.com 99dfc598be174b2d5784295c3f115a1a
```

```bash
$ ./myhttp.exe google.com http://gmail.com
http://gmail.com a0211acb7663404213fbf352c2509ccb
http://google.com b54f561b0e57ec45a0bdce916c980c96
```

This tool also takes "parallel" input parameter(integer value) to execute the http requests in parallel. The default value of "parallel" is 10.
```bash
$ ./myhttp.exe -parallel 2 google.com http://gmail.com http://www.facebook.com
http://gmail.com c3187c303a30b0aabeba65e332c0faba
http://google.com adfc9b9e273ddf921cb59eb0c3b86549
http://www.facebook.com dc484ef3a7c6accb7491c853cad3d438
```

In case of failures(say incorrect address), it prints error message and gets exited.
```bash
$ ./myhttp.exe gool
Error:  Get "http://gool": dial tcp: lookup gool: no such host
```

This repository includes dependencies only from Go standard libraries. It also contains unit test cases which provide the industry standard code coverage.