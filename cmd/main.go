package main

import (
	"log"

	"github.com/DhansAL/Dokacloud/app/router"
	"github.com/DhansAL/Dokacloud/internal/database"
	"github.com/DhansAL/Dokacloud/internal/server"
	"github.com/DhansAL/Dokacloud/lib/config"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	client, err := database.NewDatabase(c)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	log.Println("connected to database", client)
	app := server.NewFiberServer(c)
	router.Register(app)
	app.Listen(c.App.Port)

}
