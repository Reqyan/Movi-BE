package controllers

import (
	"Movi-BE/models"
	"Movi-BE/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// get all movies
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movies []models.Movie
	models.DB.Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

// get movies by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var movie models.Movie

	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Movie not Found")
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// validator
var validate *validator.Validate

// create director
// type DirectorInput struct {
// 	gorm.Model
// 	Firstname string `json:"firstname" validate:"required"`
// 	Lastname  string `json:"lastname" validate:"required"`
// }

// func CreateDirector(w http.ResponseWriter, r *http.Request) {
// 	var input DirectorInput
// 	body, _ := ioutil.ReadAll(r.Body)
// 	_ = json.Unmarshal(body, &input)

// 	validate = validator.New()
// 	err := validate.Struct(input)

// 	if err != nil {
// 		utils.RespondWithError(w, http.StatusBadRequest, "Validation error")
// 		return
// 	}

// 	director := &models.Director{
// 		Firstname: input.Firstname, Lastname: input.Lastname,
// 	}

// 	models.DB.Create(director)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(director)
// }

// create movie
type MovieInput struct {
	gorm.Model
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Isbn        string  `json:"isbn" validate:"required"`
	Genre       string  `json:"genre" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Director    struct {
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required"`
	} `json:"director"`
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var input MovieInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "validation Error")
		return
	}

	var director models.Director
	if err := models.DB.Where("firstname = ? AND lastname = ?", input.Director.Firstname, input.Director.Lastname).First(&director).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika direktor tidak ditemukan, buat direktor baru
			director = models.Director{
				Firstname: input.Director.Firstname,
				Lastname:  input.Director.Lastname,
			}
			if err := models.DB.Create(&director).Error; err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create director")
				return
			}
		} else {
			// Jika ada error lain selain direktor tidak ditemukan, kembalikan error
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	movie := &models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Isbn:        input.Isbn,
		Genre:       input.Genre,
		Rating:      input.Rating,
		DirectorID:  director.ID,
	}

	if err := models.DB.Create(movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create movie")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// Update Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var movie models.Movie

	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	var input MovieInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	var director models.Director
	if err := models.DB.Where("firstname = ? AND lastname = ?", input.Director.Firstname, input.Director.Lastname).First(&director).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika direktor tidak ditemukan, buat direktor baru
			director = models.Director{
				Firstname: input.Director.Firstname,
				Lastname:  input.Director.Lastname,
			}
			if err := models.DB.Create(&director).Error; err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create director")
				return
			}
		} else {
			// Jika ada error lain selain direktor tidak ditemukan, kembalikan error
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	movie.Title = input.Title
	movie.Description = input.Description
	movie.Isbn = input.Isbn
	movie.Genre = input.Genre
	movie.Rating = input.Rating
	movie.DirectorID = director.ID

	if err := models.DB.Save(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update movie")
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var movie models.Movie
	if err := models.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	models.DB.Delete(&movie)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(movie)
}
