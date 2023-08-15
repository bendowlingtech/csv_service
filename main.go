package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/csv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my name is %s!", r.URL.Path[1:])
}

func csvHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm( 30 >> 20)
	file, _, err  = r.FormFile("csv_file")
}

func main() {
	fmt.Printf("Starting server at port 8081")
	http.HandleFunc("/", handler)
	http.HandleFunc("/csv", csvHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))


}
