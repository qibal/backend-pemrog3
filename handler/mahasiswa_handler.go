package handler

import (
	"fmt"
	"inibackend/model"
	"inibackend/repository"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Hello API sudah jalan")
}

// GetAllMahasiswa godoc
// @Summary Get All Data Mahasiswa.
// @Description Mengambil semua data mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.Mahasiswa
// @Failure 401 "Unauthorized"
// @Failure 500
// @Router /api/mahasiswa [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data dari database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diambil",
		"data":    data,
		"status":  fiber.StatusOK,
	})
}

// GetMahasiswaByNPM godoc
// @Summary Get By NPM Data Mahasiswa.
// @Description Ambil per NPM data mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path int true "Masukan NPM"
// @Success 200 {object} model.Mahasiswa
// @Failure 400 "NPM harus berupa angka"
// @Failure 401 "Unauthorized"
// @Failure 404 "Data tidak ditemukan"
// @Router /api/mahasiswa/{npm} [get]
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung ambil sebagai string

	// Dapatkan mahasiswa dari repository
	mhs := repository.GetMahasiswaByNPM(c.Context(), npm)

	// Periksa apakah mahasiswa ditemukan dengan memeriksa NPM-nya
	// (atau field lain yang menandakan bahwa data valid)
	if mhs.NPM == 0 { // Sesuaikan dengan struktur model.Mahasiswa Anda
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
			"status":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data mahasiswa ditemukan",
		"data":    mhs,
		"status":  fiber.StatusOK,
	})
}

// InsertMahasiswa godoc
// @Summary Insert data Mahasiswa.
// @Description Input data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.Mahasiswa true "Payload Body [RAW]"
// @Success 201 {object} model.Mahasiswa
// @Failure 400 "Invalid request data"
// @Failure 401 "Unauthorized"
// @Failure 409 "Gagal menambahkan mahasiswa"
// @Router /api/mahasiswa [post]
func InsertMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	insertedID, err := repository.InsertMahasiswa(c.Context(), mhs)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan mahasiswa: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Mahasiswa berhasil ditambahkan",
		"id":      insertedID,
		"status":  fiber.StatusCreated,
	})
}

// UpdateMahasiswa godoc
// @Summary Update data Mahasiswa.
// @Description Ubah data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path integer true "Masukan NPM"
// @Param request body model.MahasiswaRequest true "Payload Body [RAW]"
// @Success 200 {object} model.Mahasiswa
// @Failure 400 "Invalid request data or Invalid NPM format"
// @Failure 401 "Unauthorized"
// @Failure 404 "Error Update Data Mahasiswa"
// @Router /api/mahasiswa/{npm} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung gunakan npm sebagai string
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	updatedNPM, err := repository.UpdateMahasiswa(c.Context(), npm, mhs) // Gunakan npm string
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Error Update Data Mahasiswa %s : %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data mahasiswa berhasil diupdate",
		"npm":     updatedNPM,
		"status":  fiber.StatusOK,
	})
}

// DeleteMahasiswa godoc
// @Summary Delete data Mahasiswa.
// @Description Hapus data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path integer true "Masukan NPM"
// @Success 200 "Mahasiswa berhasil dihapus"
// @Failure 400 "Invalid NPM format"
// @Failure 401 "Unauthorized"
// @Failure 404 "Mahasiswa tidak ditemukan"
// @Router /api/mahasiswa/{npm} [delete]
func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm") // Langsung gunakan npm sebagai string

	deletedNPM, err := repository.DeleteMahasiswa(c.Context(), npm) // Gunakan npm string
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Mahasiswa dengan NPM %s tidak ditemukan: %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mahasiswa berhasil dihapus",
		"npm":     deletedNPM,
		"status":  fiber.StatusOK,
	})
}
