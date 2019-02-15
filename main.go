package main

import (
	"io"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "It works!")
}

func main() {
	http.HandleFunc("/", root)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
