package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    ServerAddress string
    DBUser        string
    DBPassword    string
    DBName        string
    DBHost        string
}

func (c *Config) Print() {
    log.Printf("Server Address: %s", c.ServerAddress)
    log.Printf("Database User: %s", c.DBUser)
    log.Printf("Database Name: %s", c.DBName)
    log.Printf("Database Host: %s", c.DBHost)
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return &Config{
        ServerAddress: os.Getenv("SERVER_ADDRESS"),
        DBUser:        os.Getenv("DB_USER"),
        DBPassword:    os.Getenv("DB_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
        DBHost:        os.Getenv("DB_HOST"),
    }
}

