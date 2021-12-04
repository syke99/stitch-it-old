package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// create new fiber instance  and use across whole app
	app := fiber.New()

	// middleware to allow all clients to communicate using http and allow cors
	app.Use(cors.New())

	// serve  images from images directory prefixed with /images
	// i.e http://localhost:4000/images/someimage.webp
	app.Static("/images", "./images")

	// handle image uploading using post request
	app.Post("/", handleFileupload)

	// handle image processing using put request and provinding
	// image name
	app.Put("/:imageName", handleImageProcessing)

	// delete uploaded image by providing unique image name
	app.Delete("/:imageName", handleDeleteImage)
}

func handleFileupload(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	err = c.SaveFile(file, fmt.Sprintf("./images/%s", file.Filename))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully"})
}

func handleImageProcessing(c *fiber.Ctx) error {

	logger := NewLogger()

	imageName := c.Params("imageName")

	f, err := os.Open(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		logger.log.Err(err)
	}

	defer f.Close()

	// img, _, err := image.Decode(f)
	// if err != nil {
	// 	logger.log.Err(err)
	// }

	// bounds := img.Bounds()
	// w, h := bounds.Max.X, bounds.Max.Y

	// var dmcs [][]string

	// for y := 0; y < h; y++ {
	// 	var row []string
	// 	for x := 0; x < w; x++ {
	// 		col = img.At(x, y)

	// 	}
	// }

	return nil
}

func handleDeleteImage(c *fiber.Ctx) error {
	// extract image name from params
	imageName := c.Params("imageName")

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
