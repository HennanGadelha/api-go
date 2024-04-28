package main

import (
	"api-go/src/config"
	"api-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	fmt.Println("Started API", config.Porta)

	router := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))

}
