package models

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	// ID        int     `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
	// LastName  string  `json:"last_name" form:"last_name"`
	Movies []Movie `json:"movies"`
	// Movies []MovieResponse `json:"movies"`
}

type DirectorResponse struct {
	// ID   int    `json:"id" form:"id"`
	gorm.Model
	Name string `json:"name" form:"name"`
}
