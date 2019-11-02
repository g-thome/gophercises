package main

import (
	"fmt"
	"net/http"
	"urlshortner/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/go": "https://gophercises.com/exercises/urlshort",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
