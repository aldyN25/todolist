package main

import (
	"fmt"
	"log"

	"github.com/aldyN25/todolist/app/configs"
	apiv1 "github.com/aldyN25/todolist/routers/api/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// api route : api/v1
	apiv1.ApiRoutes(app)

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	log.Fatal(app.Listen(listen))
}
