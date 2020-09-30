package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count = 0
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s", r.URL.Path,)
	fmt.Println("handler: ", count)
}


func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "%d", count)
	mu.Unlock()
	fmt.Println("count: ", count)
}

/*
为什么每次求情最后都会调用默认的“/”handler函数
 */