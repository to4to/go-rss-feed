package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	//"github.com/pressly/goose/v3/database"
	//"github.com/to4to/go-rss-feed/handler"
	"github.com/to4to/go-rss-feed/handler"
	"github.com/to4to/go-rss-feed/internal/db"
	"github.com/to4to/go-rss-feed/router"
)



func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT Not Found")
	}
	fmt.Println("Running At Port: ", portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL Not Found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't Connect To Data Base")
	}

	//queries:=db.New(conn)

	// apiCfg := ApiConfig{DB: db.New(conn)}

	apiCfg:=ApiCofig{DB: db.New(conn)}

	r:=router.
	

	fmt.Printf("Server Starting At Port: %v ", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
