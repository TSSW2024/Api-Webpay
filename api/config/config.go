package config

import (
	"fmt"
	"log"
	"os"
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	if DBHost == "" || DBUser == "" || DBPassword == "" || DBPort == "" || DBName == "" {
		log.Fatal("Missing required environment variables for database connection")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}
