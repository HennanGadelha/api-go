package main

import (
	"api-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Started API")

	router := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", router))

}
