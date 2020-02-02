package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go! %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
