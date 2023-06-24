package models

type TransactionBook struct {
	Id     int `json:"id" gorm:"primary_key:auto_increment"`
	IdTransaction int  `json:"idTransaction" form:"idTransaction"`
	Transaction   Transaction `json:"transaction" form:"transaction" gorm:"foreignKey:IdTransaction"`
	IdBook int  `json:"idBook" form:"idBook"`
	Book   Book `json:"book" form:"book" gorm:"foreignKey:IdBook"`
}
