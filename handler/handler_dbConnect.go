package handler

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/to4to/go-rss-feed/internal/db"
	"github.com/to4to/go-rss-feed/models"
)

// DBConnect establishes a connection to the database using the DB_URL environment variable.
// It returns an instance of models.ApiConfig containing the database connection and an error, if any.
func DBConnect() (models.ApiConfig, error) {
    // Retrieve the database URL from the environment variables
    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("DB_URL not found in the Environment")
    }

    // Open a connection to the PostgreSQL database using the retrieved URL
    conn, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Database Connection failed:", err)
    }

    // Create an instance of models.ApiConfig with the database connection
    apiCfg := models.ApiConfig{
        DB: db.New(conn),
    }

    return apiCfg, nil
}
