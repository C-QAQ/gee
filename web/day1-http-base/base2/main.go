package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		fmt.Fprintf(w, "这是来自http-base2的ServeHttp的case/")
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "这是来自http-base2的ServeHttp的case/hello")
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		fmt.Fprintf(w, "这是来自http-base2的ServeHttp的default")
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8080", engine))
}
