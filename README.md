Solution has two parts, first is the download('/') API, where it expects a url 
and thread counts as params and downloads the file inside `downloads` folder.
Second is stats('/stats') API, this page return stats such as uniquq IPs, Server
latency etc.

<!-- Getting started -->
<!-- Installing dependencies -->
$ go get -u github.com/gorilla/mux

<!-- Build the project -->
$ go build main.go

<!-- Run the go server -->
$ ./main

<!-- Run test case -->
$ go test ./...