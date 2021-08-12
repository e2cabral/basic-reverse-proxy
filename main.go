package main

import (
	"net/http"
	"reverse-proxy-test/proxy"
)

func main() {
	parser, err := proxy.NewURIParser("https://jsonplaceholder.typicode.com")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", parser.Handle)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
