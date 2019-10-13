package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello from Go img geter")
	log.Println("This is text for log")

	http.HandleFunc("/api/imgs", GetImage)

	err := http.ListenAndServe(":37213", nil)
	if err != nil {
		log.Fatal("ListenServer: ", err)
	}
}

// GetImage geting an image from mongo and gives it as a file via http
func GetImage(w http.ResponseWriter, r *http.Request) {

	// This is an excample how to get and take URL parameters with Golang
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url param is missing")
		return
	}

	key := keys[0]

	fmt.Fprintf(w, "Url param is = "+string(key))
}

// check is a basic check for error
func check(e error) {
	if e != nil {
		log.Fatal("Unexpected error: ", e)
	}
}
