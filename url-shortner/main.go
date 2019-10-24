package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func redirectToYoutube(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://youtube.com", 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/youtube", redirectToYoutube)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
