package models

type Transaction struct {
	Id     int `json:"id" gorm:"primary_key:auto_increment"`
	IdUser int `json:"-"`
	IdBook int `json:"-"`
}
