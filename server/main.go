package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	logger := NewLogger()

	// create new fiber instance  and use across whole app
	app := fiber.New()

	// middleware to allow all clients to communicate using http and allow cors
	app.Use(cors.New())

	// serve  images from images directory prefixed with /images
	// i.e http://localhost:4000/images/someimage.webp
	app.Static("/images", "./images")

	// handle image uploading using post request
	app.Post("/", handleFileupload)

	err := app.Listen(":4000")
	if err != nil {
		logger.log.Fatal().Err(err)
	}

}
