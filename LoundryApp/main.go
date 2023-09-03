package main

import (
	"log"

	"LoundryApp/controllers/"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route untuk customers
	r.POST("/customers", controllers.TambahCustomer)
	r.PUT("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.HapusCustomer)

	// Route untuk transactions
	r.POST("/transactions", controllers.TambahTransaksi)

	// Jalankan server pada port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
