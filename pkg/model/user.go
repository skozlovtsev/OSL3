package model

type User struct {
	ID       int    `json:"id" gorm:"id;unique;autoIncrement"`
	Username string `json:"username" gorm:"primaryKey" form:"username"`
	Password string `json:"password" form:"password"`
}
