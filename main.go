package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Docker Tutorial")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "In a World Full of Apples, Go Bananas")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
