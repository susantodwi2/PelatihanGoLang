package models

import (
	"time"
)

type Transaction struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID       uint      `json:"customer_id"`
	BeratCucian      float64   `json:"berat_cucian"`
	TotalTransaksi   float64   `json:"total_transaksi"`
	PerolehanPoint   float64   `json:"perolehan_point"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
}
