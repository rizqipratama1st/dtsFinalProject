package database

import (
	"qms/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	//for Production
	// dsn := os.Getenv("CONNSTRING")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	For LocalDB Developmen
	dsn := "root:74712331@tcp(localhost:3306)/project"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.SlotBooking{})

	models.InsertBank(db)

	return db
}
