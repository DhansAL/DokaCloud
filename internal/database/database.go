package database

import (
	"log"

	"github.com/DhansAL/Dokacloud/internal/ent"
	"github.com/DhansAL/Dokacloud/lib/config"
	_ "github.com/lib/pq"
)

func NewDatabase(config *config.Config) (*ent.Client, error) {
	client, err := ent.Open("postgres", config.Db.PostgresDSN)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	log.Println("connected to database")
	defer client.Close()
	return client, err

}
