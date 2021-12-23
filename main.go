package main

import (
	"fmt"
	"log"
	"netbotript/markets"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
        markets.GetSpots()
	fmt.Println("App started")
}
