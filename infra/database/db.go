package database

import (
	"database/sql"
	"fmt"
	"github.com/marcioecom/clipbot-server/helper"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func Init() error {
	conn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		helper.GetEnv("DB_HOST"),
		helper.GetEnv("DB_PORT"),
		helper.GetEnv("DB_NAME"),
		helper.GetEnv("DB_USER"),
		helper.GetEnv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	DB = db

	if err := db.Ping(); err != nil {
		return err
	}

	zap.L().Info("connected to database")
	return nil
}
