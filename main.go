package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf(w, "Hello, my name is %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

}