package main

import (
	"log"

	"github.com/DhansAL/Dokacloud/internal/database"
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
	defer client.Close()
}
