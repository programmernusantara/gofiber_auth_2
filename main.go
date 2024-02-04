package main

import (
	// Impor library yang dibutuhkan
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

// Fungsi utama untuk menjalankan aplikasi
func main() {
	// Buat instance baru Fiber untuk aplikasi
	app := fiber.New()

	// Definisikan rute untuk halaman utama
	app.Get("/", func(c *fiber.Ctx) error {
		// Kirimkan pesan "Selamat Datang" ke browser
		return c.SendString("Selamat Datang")
	})

	// Buat data pengguna untuk autentikasi
	users := map[string]string{
		"wildan": "190205", // Nama pengguna dan kata sandi
	}

	// Definisikan rute untuk pendaftaran pengguna baru
	app.Post("/daftar", func(c *fiber.Ctx) error {
		// Buat struktur untuk menerima data pengguna dari request
		type PenggunaBaru struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Parse data JSON dari request
		var newUser PenggunaBaru
		if err := c.BodyParser(&newUser); err != nil {
			// Jika gagal memproses data JSON, kembalikan pesan error
			return c.Status(fiber.StatusBadRequest).JSON("Gagal memproses data daftar")
		}

		// Tambahkan pengguna baru ke dalam map users
		users[newUser.Username] = newUser.Password

		// Kembalikan pesan sukses pendaftaran
		return c.JSON("Pengguna " + newUser.Username + " berhasil terdaftar")
	})

	// Aktifkan middleware untuk autentikasi dasar
	app.Use(basicauth.New(basicauth.Config{
		Users:           users,
		ContextUsername: "user",
	}))

	// Definisikan rute untuk login
	app.Get("/login", func(c *fiber.Ctx) error {
		// Dapatkan nama pengguna yang sudah login
		user, _ := c.Locals("user").(string)

		// Kembalikan pesan sukses login
		return c.JSON("Selamat " + user + " anda berhasil login")
	})

	// Jalankan aplikasi pada port 8080
	app.Listen(":8080")
}
