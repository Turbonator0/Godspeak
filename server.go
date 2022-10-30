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
	mode := req.URL.Query().Get("mode")

	s := speak(countInt)
	var Speakstr string
	switch mode {
	case "json":
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		for v := range s {
			if v > countInt-2 {
				Speakstr += "\"" + s[v] + "\""
			} else {
				Speakstr += ("\"" + s[v] + "\"" + ",")
			}
		}
		fmt.Fprintf(w, "[%s]\n", Speakstr)
		break
	case "text":
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		for v := range s {
			Speakstr += s[v] + " "
		}
		fmt.Fprintf(w, "%s\n", Speakstr)
		break
	default: // I don't know why I added this here but whatever
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		for v := range s {
			Speakstr += s[v] + " "
		}
		fmt.Fprintf(w, "%s\n", Speakstr)
		break
	}
}

func server() {
	http.HandleFunc("/god", nHandler)
	fmt.Printf("Now serving on http://localhost:8080/god\n")
	http.ListenAndServe(":8080", nil)
}
func main() {
	server()
}
