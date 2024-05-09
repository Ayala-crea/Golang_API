package belajarHandler

import (
	"strconv"

	"github.com/Ayala-Crea/toko-bersama/database"
	"github.com/Ayala-Crea/toko-bersama/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetUsers adalah handler untuk mendapatkan pengguna dari database
func GetUsers(c *fiber.Ctx) error {
	{
		var users []model.User

		// Ambil data pengguna dari database
		result := database.DB.Find(&users)

		if result.Error != nil {
			// Jika terjadi kesalahan, kembalikan respons dengan status Bad Request dan pesan kesalahan
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}

		// Jika berhasil, kembalikan respons dengan status OK dan data pengguna
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Data User Berhasil Ditampilkan!",
			"data":    users,
		})
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Berhasil Ditambahkan",
		"data":    user,
	})
}

func GetUser(c *fiber.Ctx) error {
	//get id from url
	id := c.Params("user_id")

	var user model.User
	result := database.DB.First(&user, id)

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
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
    idStr := c.Params("id_user")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "ID pengguna tidak valid",
        })
    }

    var user model.User
    result := database.DB.First(&user, id)

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Pengguna tidak ditemukan",
        })
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pengguna berhasil dihapus",
    })
}

func UpdateData(c *fiber.Ctx) error {
	id := c.Params("user_id")

	var user model.User
	result := database.DB.First(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data yang ingin di update tidak ada",
		})
	}

	var newUser model.User
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}
	result = database.DB.Model(&user).Updates(newUser)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil dicari",
		"data":    user,
	})
}

// func Index(c *gin.Context){}
