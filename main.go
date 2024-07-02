package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/to4to/go-rss-feed/router"
)



func main() {


r:=router.Router()

	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port Not Bound ..failed to connect from env")

	}

	fmt.Println("PORT :", port)

	server := &http.Server{

		Handler: *r,
		Addr:    ":" + port,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
