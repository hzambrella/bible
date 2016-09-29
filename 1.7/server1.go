package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "url.path=%q\n", r.URL.Path)
}
