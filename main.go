package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// logger := NewLogger()

	// create new fiber instance  and use across whole app
	app := fiber.New()

	// middleware to allow all clients to communicate using http and allow cors
	app.Use(cors.New())

	// serve  react app from stitch-it/stitch-it directory
	app.Static("/stitch-it", "./stitch-it")

	// serve index.html at root path
	app.Get("/", func(c *fiber.Ctx) error {
		if err := c.SendFile("./stitch-it/public/index.html"); err != nil {
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": err})
		}

		return c.Next()
	})

	// handle image uploading using post request
	app.Post("/images", handleFileupload)

	// handle image processing using put request
	app.Put("/images/:image", handleImageProcessing)

	// handle excel download using get request
	app.Get("/patterns/:image", func(c *fiber.Ctx) error {

		f := c.Params("image")

		patNm := strings.Split(f, ".")[0]

		if err := c.SendFile("./public/patterns/" + patNm + ".xlsx"); err != nil {
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": err})
		}

		handleDeleteImage(c, f)

		return nil
	})

	// attempt to serve app on localhost:4000
	err := app.Listen(":4000")
	if err != nil {
		// if it errors, try once more and log if
		// it errors again
		log.Fatal(app.Listen(":4000"))
	}

}
