package main

import (
	"fmt"
	"inibackend/config"
	_ "inibackend/docs"
	"inibackend/router"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	// Load file .env saat program dijalankan
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Gagal memuat file .env")
	// }
	if _,err=os.Stat(".env");err ==nil{
		fmt.Println("File .env tidak ditemukan")
	} else {
		fmt.Println("File .env ditemukan")
	}
}

// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /
// @schemes http https
func main() {
	app := fiber.New()

	// Logging request
	app.Use(logger.New())

	// Basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setup router
	router.SetupRoutes(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Endpoint not found",
		})
	})
	// Baca PORT dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port kalau tidak ada
	}

	log.Printf("Server is running at http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
