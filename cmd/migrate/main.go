package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/afallenhope/go-vendor/config"
	"github.com/afallenhope/go-vendor/db"
)

func main() {

	port, err := strconv.Atoi(config.Envs.DBPort)

	db, err := db.NewPostgresStorage(db.PgConfig{
		Host:     config.Envs.DBHost,
		Port:     port,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Name:     config.Envs.DBName,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if strings.ToLower(cmd) == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if strings.ToLower(cmd) == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
