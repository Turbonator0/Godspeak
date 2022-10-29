package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func nHandler(w http.ResponseWriter, req *http.Request) {
	count := req.URL.Query().Get("count")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Printf("Missing count parameter\n")
		countInt = 10 // Default to 10, no reason
	}
	s := speak(countInt)
	var Speakstr string
	for v := range s {
		if v > countInt-2 {
			Speakstr += "\"" + s[v] + "\""
		} else {
			Speakstr += ("\"" + s[v] + "\"" + ",")
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "[%s]", Speakstr)
}

func server() {
	http.HandleFunc("/god", nHandler)
	http.ListenAndServe(":8080", nil)
	// Handle request (paramter N [/godspeak?words=<N>])
	// Return the JSON to server the client
}
func main() {
	server()
}
