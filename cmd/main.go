package main

import (
	"golfy-api-go/cmd/models"

	"github.com/gofiber/fiber/v2"
)

type (
	Player = models.Player
	Course = models.Course
	UUID   = models.UUID
)

func main() {
	app := fiber.New()

	players := make(map[UUID]Player)
	courses := make(map[UUID]Course)

	app.Get("/players", getPlayers)
	app.Get("/players/:playerID")
	app.Post("/players")
	app.Put("/players/:playerID")
	app.Delete("/players/:playerID")
}

func getPlayers(c *fiber.Ctx) error {
	return nil
}
