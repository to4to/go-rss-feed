package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {


	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port Not Bound ..failed to connect from env")

	}


	fmt.Printf("PORT :",port)
}
