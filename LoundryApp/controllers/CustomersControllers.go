package controllers

import (
	"LoundryApp/models/"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TambahCustomer digunakan untuk menambah pelanggan ke database
func TambahCustomer(c *gin.Context) {
	nama := c.PostForm("nama")
	nomorHP := c.PostForm("nomor_hp")
	email := c.PostForm("email")

	// Validasi data input jika diperlukan
	if nama == "" || nomorHP == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	customer := models.Customers{
		Nama:    nama,
		NomorHP: nomorHP,
		Email:   email,
	}

	if err := models.TambahCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer berhasil ditambahkan", "data": customer})
}

// UpdateCustomer digunakan untuk mengupdate pelanggan di database berdasarkan ID
func UpdateCustomer(c *gin.Context) {
	// Ambil id dari parameter URL
	idParam := c.Param("id")

	// Konversi idParam menjadi tipe data uint
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	// Ambil data customer dari database berdasarkan id
	customer, err := models.GetCustomerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer tidak ditemukan"})
		return
	}

	// Baca data yang akan di-update dari form
	nama := c.PostForm("nama")
	nomorHP := c.PostForm("nomor_hp")
	email := c.PostForm("email")

	// Validasi data input jika diperlukan
	if nama == "" || nomorHP == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	// Update data customer
	customer.Nama = nama
	customer.NomorHP = nomorHP
	customer.Email = email

	if err := models.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer berhasil di-update", "data": customer})
}

// HapusCustomer digunakan untuk menghapus pelanggan dari database berdasarkan ID
func HapusCustomer(c *gin.Context) {
	// Ambil id dari parameter URL
	idParam := c.Param("id")

	// Konversi idParam menjadi tipe data uint
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	// Cek apakah customer dengan ID tersebut ada di database
	if _, err := models.GetCustomerByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer tidak ditemukan"})
		return
	}

	// Hapus customer dari database berdasarkan ID
	if err := models.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer berhasil dihapus"})
}
