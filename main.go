package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/to4to/go-rss-feed/router"
)

// main is the entry point of the program.
func main() {
	r := router.Router()

	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port Not Bound ..failed to connect from env")
	}

	fmt.Println("PORT :", port)

	server := &http.Server{

		Handler: r,
		Addr:    ":" + port,
	}

	// Start the HTTP server and listen for incoming requests.
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in the Environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Database Connection failed:", err)
	}
}
