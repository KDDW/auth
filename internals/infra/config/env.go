package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot read .env file: ", err)
	}
}
