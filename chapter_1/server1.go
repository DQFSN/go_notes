package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/handler1", handler1)
	http.HandleFunc("/handler2", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler1 URL.Path = %q\n", r.URL.Path)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler2 URL.Path = %q\n", r.URL.Path)
}