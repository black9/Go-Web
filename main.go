package main

import (
	"fmt"
	"net/http"
)

type guHandler struct{}

func (f *guHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GU Handler")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, GO Web")
	})

	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is Page1")
	})

	http.Handle("/gu", &guHandler{})

	http.ListenAndServe(":3000", nil)
}
