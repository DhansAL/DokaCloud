package database

import "github.com/DhansAL/Dokacloud/internal/ent"

type Database struct {
	Ent *ent.Client
}
type Seeder interface {
	Seed(*ent.Client) error
	Count() (int, error)
}
