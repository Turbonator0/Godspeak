# Godsay API in Go

Requires: 
- Go
- Make (Optional, compile with `go build -o godspeak server.go speak.go`)


## Usage 

`curl 0.0.0.0:8080/god` (Defaults to 10 words)

`curl 0.0.0.0:8080/god?count=N` Where N is the desired number of words