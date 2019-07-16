## A download manager written in Go.
App allows to configure number of thread that should be used to download any type of file.

## Pre-requisites
* Go v1.12 or higher

## Setup
Run bash command (mentioned below) from root of the app directory.
```
$ bash bin/setup.sh <!-- Install the dependencies and run server-->
```

## To Manually setup the project
```
$ go get -u github.com/gorilla/mux <!-- Install dependencies -->

$ go build main.go <!-- Complie the project -->

$ go test -v <!-- Run test cases -->

$ ./main <!-- Run the server -->
```
Api endpoint download('/) expects two params `url` (string) and `threads` (int).
```
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
```

## Postman collection
Postman collection is located in `/test/postman/Go_DL_MANAGER_API.postman_collection`
Import this collection to test the Apis via Postman. It contains two POST requests.
First with valid url and Second with invalid url.

## Download location
The downloaded files will be inside `go_dl_manager/downloads/` folder.

## Logs
Logs are saved in `go_dl_manager/development.log` file. It contains the informations
like process ID(to differentiate between calls), IP address, params, request method
and time taken by server to process the incoming request.
Ex:
```
pid:[32045]		[::1]:62774 params: {"url":"example.com","threads":4}
pid:[32045]		POST		/		[::1]:62774		2.499650223s
```
