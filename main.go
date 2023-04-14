package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const responseEnv = "RESPONSE_TEXT"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got request\n")
		io.WriteString(w, os.Getenv(responseEnv)+"\n")
	})

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}
