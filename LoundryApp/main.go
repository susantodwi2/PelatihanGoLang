package main

import (
	"LoundryApp/config"
	"LoundryApp/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inisialisasi koneksi ke database
	models.DB = config.DBInit()
	defer models.DB.Close()

	r.POST("/tambah-customer", controllers.TambahCustomer)
	r.PUT("/update-customer/:id", controllers.UpdateCustomer)
	r.DELETE("/hapus-customer/:id", controllers.HapusCustomer)

	r.Run(":8080") // Ganti port sesuai kebutuhan Anda
}
