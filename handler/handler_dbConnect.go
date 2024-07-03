package handler

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/to4to/go-rss-feed/internal/db"
	"github.com/to4to/go-rss-feed/models"
)

func DBConnect() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in the Environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Database Connection failed:", err)
	}

	apiCfg := models.ApiConfig{
		DB: db.New(conn),
	}
}
