package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	// ID   int    `json:"id" form:"id" gorm:"primaryKey"`
	Title       string   `json:"title" form:"title"`
	Description string   `json:"description" form:"description"`
	Isbn        string   `json:"isbn" form:"isbn"`
	Genre       string   `json:"genre" form:"genre"`
	Rating      float64  `json:"rating" form:"rating"`
	DirectorID  uint     `json:"director_id" form:"director_id"`
	Director    Director `json:"director"`
	// Director DirectorResponse `json:"director"`
}

type MovieResponse struct {
	gorm.Model
	// ID          int     `json:"id" gorm:"primary_key"`
	Title       string  `json:"title" form:"title"`
	Description string  `json:"description" form:"description"`
	Isbn        string  `json:"isbn" form:"isbn"`
	Genre       string  `json:"genre" form:"genre"`
	Rating      float64 `json:"rating" form:"rating"`
	DirectorID  uint    `json:"-" form:"director_id"`
}

func (MovieResponse) TableName() string {
	return "movies"
}
