package server

import (
	"github.com/DhansAL/Dokacloud/lib/config"
	"github.com/gofiber/fiber/v2"
)

func NewFiberServer(config *config.Config) *fiber.App {
	// setup fiber webserver
	app := fiber.New()
	return app
}
