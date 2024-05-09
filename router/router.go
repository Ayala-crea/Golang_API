package router

import (
	belajarRoutes "github.com/Ayala-Crea/toko-bersama/internal/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	// Setup the Node Routes
	belajarRoutes.SetupBelajarRoutes(api)
}
