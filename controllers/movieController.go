package controllers

import (
	"Movi-BE/models"
	"Movi-BE/utils"
	"encoding/json"
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
	models.ConnectDatabase().Preload("Movies").Preload("Director").Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

// get movies by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var movie models.Movie

	if err := models.ConnectDatabase().Preload("Movies").Preload("Director").Where("id = ?", id).First(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Movie not Found")
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// validator
var validate *validator.Validate

// create movie
type MovieInput struct {
	gorm.Model
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Isbn        string  `json:"isbn" validate:"required"`
	Genre       string  `json:"genre" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	DirectorID  uint    `json:"director_id"`
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

	// Set DirectorID di MovieInput

	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Isbn:        input.Isbn,
		Genre:       input.Genre,
		Rating:      input.Rating,
		DirectorID:  input.DirectorID,
	}

	if err := models.ConnectDatabase().Preload("Director").Create(&movie).Error; err != nil {
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

	if err := models.ConnectDatabase().Where("id = ?", id).First(&movie).Error; err != nil {
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

	// var director models.Director
	// if err := models.ConnectDatabase().Where("first_name = ? AND last_name = ?", input.Director.Firstname, input.Director.Lastname).First(&director).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		// Jika direktor tidak ditemukan, buat direktor baru
	// 		director = models.Director{
	// 			FirstName: input.Director.Firstname,
	// 			LastName:  input.Director.Lastname,
	// 		}
	// 		if err := models.ConnectDatabase().Create(&director).Error; err != nil {
	// 			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create director")
	// 			return
	// 		}
	// 	} else {
	// 		// Jika ada error lain selain direktor tidak ditemukan, kembalikan error
	// 		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// }

	movie.Title = input.Title
	movie.Description = input.Description
	movie.Isbn = input.Isbn
	movie.Genre = input.Genre
	movie.Rating = input.Rating
	movie.DirectorID = input.DirectorID
	if err := models.ConnectDatabase().Save(&movie).Error; err != nil {
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
	if err := models.ConnectDatabase().Where("id = ?", id).First(&movie).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	models.ConnectDatabase().Delete(&movie)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(movie)
}
