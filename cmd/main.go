package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/afallenhope/go-vendor/cmd/api"
	"github.com/afallenhope/go-vendor/config"
	"github.com/afallenhope/go-vendor/db"
)

func main() {
	port, err := strconv.Atoi(config.Envs.DBPort)

	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewPostgresStorage(db.PgConfig{
		Host:     config.Envs.DBHost,
		Port:     port,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Name:     config.Envs.DBName,
	})

	initStorage(db)
	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
