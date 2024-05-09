package main

import (
	"github.com/Ayala-Crea/toko-bersama/database"
	"github.com/Ayala-Crea/toko-bersama/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Tambahkan middleware CORS ke aplikasi Fiber
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Content-Type",
	}))

	// Membuat koneksi ke database
	database.ConnectDB()

	// Menyiapkan rute
	router.SetupRoutes(app)

	// Mendengarkan permintaan pada port 3000
	app.Listen(":3000")
}
