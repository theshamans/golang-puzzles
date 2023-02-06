package main

import (
	"fmt"
	"net/http"
)

type rootHandler struct{}

func main() {

	mux := http.NewServeMux()
	r := rootHandler{}

	// Make the necessary changes such that it is possible to set r as the 
	// handler for the root path. Whenever a request is sent to any path,
	// including paths like /hello, /goodbye, the code should print "handler
	// hit!"
	mux.Handle("/", r)

	http.ListenAndServe(":9090", mux)
}
