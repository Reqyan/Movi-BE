package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint     `json:"id" gorm:"primary_key"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Isbn        string   `json:"isbn"`
	Genre       string   `json:"genre"`
	Rating      float64  `json:"rating"`
	DirectorID  uint     `json:"director_id"`
	Director    Director `gorm:"references:DirectorID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"director"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}
