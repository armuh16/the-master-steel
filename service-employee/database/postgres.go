package database

import (
	"database/sql"
	"fmt"
	"log"
	"service-user/config"
	"time"

	_ "github.com/lib/pq"
)

// DB is a global database connection pool.
var DB *sql.DB

func init() {
	cfg := config.Get() // Fetching the configuration

	// Constructing the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Optional: Configure the database connection pool
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Hour)
	DB.SetConnMaxIdleTime(15 * time.Minute)

	fmt.Println("Successfully connected to PostgreSQL")
}

// GetDB returns the database connection pool.
func GetDB() *sql.DB {
	return DB
}
