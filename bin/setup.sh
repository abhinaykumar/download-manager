#!/bin/bash -e

# Install Go packages
go list github.com/gorilla/mux
if [ $? -eq 0 ]; then
    echo "mux is already installed"
else
	echo "Installing package mux..."
    go get -u github.com/gorilla/mux
fi

# Build the project
go build main.go

echo `pwd`
# Run test case
go test -v
