package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	webService := gee.New()
	webService.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "调用了get-/\npath: %s\n", req.URL.Path)
	})

	webService.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "调用了get-/hello\npath: %s\n", req.URL.Path)
	})

	webService.POST("/post", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "调用了post-/post\npath: %s\n", req.URL.Path)
	})

	webService.Run(8080)
}
