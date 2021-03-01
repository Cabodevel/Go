package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola Sofía"))
	})

	err := http.ListenAndServe(":5001", nil)

	if err != nil {
		log.Fatal(err)
	}
}
