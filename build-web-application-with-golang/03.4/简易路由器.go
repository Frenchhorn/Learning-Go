package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Say hello")
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
