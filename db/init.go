package db

import (
	"database/sql"
	"docs-sync-service/config"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(cfg *config.Config) *sql.DB  {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}
	return db
}
