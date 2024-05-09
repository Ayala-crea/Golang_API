package belajarRoutes

import (
	hargaHandler "github.com/Ayala-Crea/toko-bersama/internal/handlers/harga"

	"github.com/gofiber/fiber/v2"
)

func SetupBelajarRoutes(router fiber.Router) {
	harga := router.Group("/makanan")
	harga.Get("/", hargaHandler.GetHarga)
	harga.Post("/", hargaHandler.CreateHarga)
	harga.Get("/:Id", hargaHandler.GetMakananById)
	harga.Delete("/:Id", hargaHandler.DeleteHarga)
	harga.Put("/:Id", hargaHandler.UpdateDataMakanan)
}
