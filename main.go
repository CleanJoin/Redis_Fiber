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
	//value := internal.ConnectRedis()
	database, err := internal.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	hackers, _ := database.GetLeaderboard()
	app := fiber.New()
	//dd := fmt.Sprintf("%+v", hackers.Users)
	app.Get("/json/hackers", func(c *fiber.Ctx) error {
		return c.JSON(hackers.Users)
	})
	log.Fatal(app.Listen(":8010"))

}
