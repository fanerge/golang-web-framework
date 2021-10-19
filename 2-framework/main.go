package main

// curl http://localhost:9999/
// curl http://localhost:9999/hello
// curl http://localhost:9999/world

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New() // engine
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
