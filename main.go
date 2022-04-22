package main

import (
	"fmt"
	"log"

	"redisFiber/internal"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/idoko/rediboard/db"
)

var (
	ListenAddr = ":8010"
	RedisAddr  = "localhost:6379"
)

func main() {
	value := internal.ConnectRedis()
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	database.GetLeaderboard()
	app := fiber.New()

	app.Get("/json/hackers", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%v", value))
	})
	log.Fatal(app.Listen(":8010"))

}
