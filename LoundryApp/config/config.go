package config

import (
	"LoundryApp/models/"

	"gorm.io/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "mysql://root:7PGnHAzKRrwogmEyynXQ@containers-us-west-158.railway.app:6301/railway")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Customers{}, &models.Transaksi{})
	return db
}
