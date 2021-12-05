package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// )

// func handleDeleteImage(c *fiber.Ctx) error {
// 	// extract image name from params
// 	imageName := c.Params("imageName")

// 	// delete image from ./images
// 	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
// 	if err != nil {
// 		log.Println(err)
// 		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
// 	}

// 	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
// }
