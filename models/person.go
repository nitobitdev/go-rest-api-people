package models

type Person struct {
	ID      uint   `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type RequestPerson struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
