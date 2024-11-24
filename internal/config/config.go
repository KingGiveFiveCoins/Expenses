package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/KingGiveFiveCoins/Expenses/internal/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Config struct {
	Port string
}

func Load() *Config {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
		os.Exit(1)
	}
	DB = db

	utils.InitUtils([]byte(os.Getenv("JWT_SECRET")))
	return &Config{
		Port: os.Getenv("PORT"),
	}
}
