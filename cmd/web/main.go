package main

import (
	"fmt"
	"github.com/saibeo/website/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting web server on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
