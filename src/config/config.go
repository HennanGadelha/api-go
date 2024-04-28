package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Conn  = ""
	Porta = 0
)

func Carregar() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Porta, error = strconv.Atoi(os.Getenv("API_PORT"))

	if error != nil {
		Porta = 9000
	}

	Conn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
