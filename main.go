package main

import (
	"net/http"
	// "github.com/black9/GO-WEB/web"
)

func main() {

	http.ListenAndServe(":3000", goweb.NewHttpHandler())
}
