package main

import (
	"fmt"

	"github.com/emregocer/golang_project/config"
	"github.com/emregocer/golang_project/internal/database"
	"github.com/emregocer/golang_project/internal/server"

	_ "github.com/lib/pq"
)

func main() {
	config.InitializeEnvironment()
	conf := config.NewConfig()

	server := server.Server{}

	db, err := database.NewDB(conf.Database)
	if err != nil {
		panic("could not connect to the database.")
	}
	defer db.Close()
	server.DB = &db

	server.Initialize()

	fmt.Println("Server runs at:", 8080)
	server.Listen(8080)
}
