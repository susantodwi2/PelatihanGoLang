package main

import (
	"log"

	"LoundryApp/controllers/CustomersControllers"
	"LoundryApp/controllers/TransaksiControllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route untuk customers
	r.POST("/customers", CustomersControllers.TambahCustomer)
	r.PUT("/customers/:id", CustomersControllers.UpdateCustomer)
	r.DELETE("/customers/:id", CustomersControllers.HapusCustomer)

	// Route untuk transactions
	r.POST("/transactions", TransaksiControllers.TambahTransaksi)

	// Jalankan server pada port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
