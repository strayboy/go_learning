package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"test\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}
