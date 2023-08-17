package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/csv"
	"encoding/json"

)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my name is %s!", r.URL.Path[1:])
}

func csvHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm( 30 >> 20)
	file, _, err  = r.FormFile("csv_file")
	if err != nill {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records,err = reader.ReadAll()
	if err != nill {
		http.Error(w, "Error reading CSV", http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(w, "Empty CSV", http.StatusBadRequest)
	}


}

func main() {
	fmt.Printf("Starting server at port 8081")
	http.HandleFunc("/", handler)
	http.HandleFunc("/csv", csvHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))



}
