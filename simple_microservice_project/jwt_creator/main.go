package main

import (
	"fmt"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Println(w, string(validToken))
}
func handleRequests() {
	http.HandleFunc("/", Index)
}

func main() {

}
