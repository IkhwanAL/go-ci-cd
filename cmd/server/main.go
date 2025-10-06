package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", os.Getenv("POD_NAME"))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	port := "8080"
	log.Printf("listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
