package main

import (
	"log"

	"github.com/gporawetku/go-blog-api-example/config"
	"github.com/gporawetku/go-blog-api-example/database"
	"github.com/gporawetku/go-blog-api-example/migrations"
	"github.com/gporawetku/go-blog-api-example/server"
)

func main() {
	config.LoadConfig()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(db)
	server.Start()
}
