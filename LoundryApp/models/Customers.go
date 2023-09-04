package models

import "gorm.io/gorm"

type Customers struct {
	ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama    string  `json:"nama"`
	NomorHP string  `json:"nomor_hp"`
	Email   string  `json:"email"`
	Point   float64 `json:"point"`
}

// TambahCustomer digunakan untuk menambah pelanggan ke database
func TambahCustomer(customer *Customers) error {
	if err := DB.Create(customer).Error; err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat menambah pelanggan
	}
	return nil
}

// GetCustomerByID mengambil data pelanggan berdasarkan ID
func GetCustomerByID(id uint) (Customers, error) {
	var customer Customers
	if err := DB.Where("ID = ?", id).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return customer, err // Pelanggan tidak ditemukan
		}
		return customer, err // Terjadi kesalahan lain saat mengambil data
	}
	return customer, nil // Mengembalikan data pelanggan jika berhasil ditemukan
}

// UpdateCustomer digunakan untuk mengupdate data pelanggan di database
func UpdateCustomer(customer *Customers) error {
	if err := DB.Save(customer).Error; err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat menyimpan perubahan
	}
	return nil
}

// DeleteCustomer digunakan untuk menghapus pelanggan dari database berdasarkan ID
func DeleteCustomer(id uint) error {
	if err := DB.Delete(&Customers{}, id).Error; err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat menghapus
	}
	return nil
}
