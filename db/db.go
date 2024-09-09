package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func NewPostgresStorage(cfg PgConfig) (*sql.DB, error) {
	pgsql := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)
	db, err := sql.Open("postgres", pgsql)

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
