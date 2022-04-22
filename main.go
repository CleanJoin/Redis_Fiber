package main

import (
	"log"

	"redisFiber/internal"

	"github.com/gofiber/fiber/v2"
)

var (
	ListenAddr = ":8010"
	RedisAddr  = "localhost:6379"
)

func main() {

	database, err := internal.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	hackers, err := database.GetHackers()

	if err != nil {
		log.Fatalf("%v", err)
	}
	app := fiber.New()
	app.Get("/json/hackers", func(c *fiber.Ctx) error {
		return c.JSON(hackers.Users)
	})
	log.Fatal(app.Listen(":8010"))

}
