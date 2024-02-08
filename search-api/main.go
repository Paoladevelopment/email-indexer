package main

import (
	"fmt"
	"net/http"

	"github.com/Paoladevelopment/search-api/routes"
)

func main() {
	r := routes.SetUpRouter()

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
