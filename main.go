package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/exit", exit)
	http.HandleFunc("/terminate", terminate)

	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Path: %s", r.URL.Path)
}

func exit(w http.ResponseWriter, r *http.Request) {

	log.Fatal("Exiting 0...")
	os.Exit(0)
}

func terminate(w http.ResponseWriter, r *http.Request) {

	log.Fatal("Terminating 1...")
	os.Exit(1)
}
