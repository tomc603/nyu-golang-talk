package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, 世界</h1>")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
