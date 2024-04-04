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
	// app.Post("/players")
	// app.Put("/players/:playerID")
	// app.Delete("/players/:playerID")

	app.Get("/courses", getCourses)
	// app.Post("/courses")
	// app.Put("/courses/:coursesID")
	// app.Delete("/courses/:coursesID")
}

func getPlayers(c *fiber.Ctx) error {
	return nil
}

func getCourses(c *fiber.Ctx) error {
	return nil
}
