package controllers

import (
	"LoundryApp/models/"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TambahTransaksi(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan CustomerID dari inputan
	customerIDStr := c.PostForm("CustomerID")
	customerID, err := strconv.Atoi(customerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CustomerID harus berupa angka"})
		return
	}

	// Mendapatkan Customer berdasarkan CustomerID
	customer, err := models.GetCustomerByID(uint(customerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer tidak ditemukan"})
		return
	}

	// Menghitung TotalTransaksi
	beratCucianStr := c.PostForm("BeratCucian")
	beratCucian, err := strconv.ParseFloat(beratCucianStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BeratCucian harus berupa angka"})
		return
	}
	transaction.TotalTransaksi = beratCucian * 5000

	// Menghitung PerolehanPoint
	transaction.PerolehanPoint = transaction.TotalTransaksi / 50

	// Mengatur TanggalTransaksi ke waktu sekarang
	transaction.TanggalTransaksi = time.Now()

	// Menyimpan Transaksi ke database
	err = models.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menambahkan PerolehanPoint ke saldo pelanggan
	customer.Point += transaction.PerolehanPoint
	err = models.UpdateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
