package models

type Admin struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password" db:"password_hash"`
	Level int `json:"level"`
}

type AdminSignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}