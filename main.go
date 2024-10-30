// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
    "os"
)

func main() {
    // Database connection parameters from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

    // Open a connection to the database
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }
    log.Println("Connected to the database successfully!")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}

