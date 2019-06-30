Api endpoint download('/) expects two params `url`(string) and `threads`(int).

Endpoint: http://localhost:3000 
Request Method: POST
Content-Type: json
Sample request:
{
    "url": "https://file-examples.com/wp-content/uploads/2017/11/file_example_MP3_700KB.mp3",
    "threads": 3
}

Sample response:
File is downloaded from url: https://file-examples.com/wp-content/uploads/2017/11/file_example_MP3_700KB.mp3

<!-- To install the dependencies and run server-->
$ bash bin/setup.sh

###Manually setting up the project
<!-- Installing dependencies -->
$ go get -u github.com/gorilla/mux

<!-- Build the project -->
$ go build main.go

<!-- Run test case -->
$ go test -v

<!-- Run the go server -->
$ ./main