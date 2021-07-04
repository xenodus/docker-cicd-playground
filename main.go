package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Hello World")
	fmt.Println("Env value is:", os.Getenv("MY_ENV_VAL"))
	fmt.Println("Other env value is:", os.Getenv("MY_OTHER_ENV_VAL"))
}
