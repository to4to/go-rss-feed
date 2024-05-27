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
	"github.com/pressly/goose/v3/database"
//"github.com/to4to/go-rss-feed/internal"
	"github.com/to4to/go-rss-feed/handler"

)



type apiConfig struct{
DB *sql.DB
}
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

	database.New(conn)
	apiCfg := apiConfig{DB: conn}

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{

			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerErr)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server Starting At Port: %v ", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
