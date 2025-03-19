package main

import (
	"log"
	"net/http"
	"rest-project/internal/controllers"
)

func main() {
	controllers.New()

	server := &http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
