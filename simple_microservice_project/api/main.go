package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request))
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main() {
	fmt.Printf("server")
	handleRequests()
}
