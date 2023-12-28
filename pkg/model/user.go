package model

type User struct {
	Username string `json:"username" gorm:"primaryKey" form:"username"`
	Password string `json:"password" form:"password"`
}
