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
	secret := os.Getenv("LINE_CHANNEL_SECRET")
	token := os.Getenv("LINE_CHANNEL_TOKEN")

	fmt.Println(secret)
	fmt.Println(token)
}
