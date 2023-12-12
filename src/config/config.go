package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Secret           []byte
	Port             = 0
)

func LoadConfig() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))

	if error != nil {
		Port = 9000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	Secret = []byte(os.Getenv("SECRET"))
}
