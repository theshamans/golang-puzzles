package main

import (
	"fmt"
	"net/http"
)

type rootHandler struct{}

func (r rootHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	fmt.Println("handler hit!")
}

func main() {

	mux := http.NewServeMux()
	r := rootHandler{}

	mux.Handle("/", r)

	http.ListenAndServe(":9090", mux)
}
