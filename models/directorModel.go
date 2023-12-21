package models

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	Movies    []Movie `gorm:"foreignKey:DirectorID"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
}
