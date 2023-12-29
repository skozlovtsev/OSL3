package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// https://gorm.io/docs/security.html#Query-Condition
func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func GetUserByID(id int) (User, error) {
	var user User

	result := db.First(&user, "id = ?", id)

	return user, result.Error
}

func GetUserByUsernamePassword(username string, password string) (User, error) {
	var user User

	result := db.First(&user, "username = ?", username, "password = ?", password)

	return user, result.Error
}
