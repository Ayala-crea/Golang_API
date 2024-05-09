package hargaHandler

import (
	"strconv"

	"github.com/Ayala-Crea/toko-bersama/database"
	model "github.com/Ayala-Crea/toko-bersama/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetHarga(c *fiber.Ctx) error {
	{
		var harga []model.Harga

		// Ambil data pengguna dari database
		result := database.DB.Find(&harga)

		if result.Error != nil {
			// Jika terjadi kesalahan, kembalikan respons dengan status Bad Request dan pesan kesalahan
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}

		// Jika berhasil, kembalikan respons dengan status OK dan data pengguna
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Data User Berhasil Ditampilkan!",
			"data":    harga,
		})
	}
}

func CreateHarga(c *fiber.Ctx) error {
	var harga model.Harga
	if err := c.BodyParser(&harga); err != nil {
		return err
	}

	result := database.DB.Create(&harga)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Berhasil Ditambahkan",
		"data":    harga,
	})
}

func GetMakananById(c *fiber.Ctx) error {
	//get id from url
	Id := c.Params("Id")

	var harga model.Harga
	result := database.DB.First(&harga, Id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user yang dicari tidak ada",
		})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil dicari",
		"data":    harga,
	})
}

func DeleteHarga(c *fiber.Ctx) error {
    idStr := c.Params("Id")
    Id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "ID pengguna tidak valid",
        })
    }

    var harga model.Harga
    result := database.DB.First(&harga, Id)

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Pengguna tidak ditemukan",
        })
    }

    if err := database.DB.Delete(&harga).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pengguna berhasil dihapus",
		"data": harga,
    })
}

func UpdateDataMakanan(c *fiber.Ctx) error {
	Id := c.Params("Id")

	var harga model.Harga
	result := database.DB.First(&harga, Id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data yang ingin di update tidak ada",
		})
	}

	var newMakanan model.Harga
	if err := c.BodyParser(&newMakanan); err != nil {
		return err
	}
	result = database.DB.Model(&harga).Updates(newMakanan)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil dicari",
		"data":    harga,
	})
}