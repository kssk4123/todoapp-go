package config 

import (
        "database/sql"
        "log"
        "os"
        "fmt"

        "github.com/joho/godotenv"
        _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
        err := godotenv.Load()
        if err != nil {
                log.Fatal("Error loading .env file")
        }

        user := os.Getenv("DB_USER")
        password := os.Getenv("DB_PASS")
        host := os.Getenv("DB_HOST")
        port := os.Getenv("DB_PORT")
        dbName := os.Getenv("DB_NAME")

        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify the connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	log.Println("Database connected successfully!")
}
