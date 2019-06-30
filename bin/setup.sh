#!/bin/bash -e

# Install Go packages
echo "Installing dependencies..."
echo ""
go list github.com/gorilla/mux > /dev/null
if [ $? -eq 0 ]; then
    echo "mux is already installed"
    echo ""
else
	echo "Installing package mux..."
    go get -u github.com/gorilla/mux
fi

# Build the project
echo "Compiling program..."
go build main.go
echo ""

# Run test case
echo "Running the test cases..."
echo ""
go test -v
echo ""

echo "Server running on port: 3000"
./main