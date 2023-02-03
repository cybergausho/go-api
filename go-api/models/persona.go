package models

import "github.com/jinzhu/gorm"

type Persona struct {
	gorm.Model

	//ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
