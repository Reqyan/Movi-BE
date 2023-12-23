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

// get all director
func getAllDirectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var directors []models.Director
	models.ConnectDatabase().Preload("Movies").Preload("Director").Find(&directors)
	json.NewEncoder(w).Encode(directors)
}

// get directors by id
func getDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var director models.Director

	if err := models.ConnectDatabase().Where("id = ?", id).First(&director).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Director not Found")
		return
	}

	json.NewEncoder(w).Encode(director)
}

// validator
// var validateDirector *validator.Validate

// create Director
type DirectorInput struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
}

func createDirector(w http.ResponseWriter, r *http.Request) {
	var input DirectorInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "validation Error")
		return
	}

	director := &models.Director{
		Name: input.Name,
	}

	if err := models.ConnectDatabase().Create(director).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create director")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}

// Update Director
func updateDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var director models.Director

	if err := models.ConnectDatabase().Where("id = ?", id).First(&director).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Director not found")
		return
	}

	var input DirectorInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	director.Name = input.Name

	if err := models.ConnectDatabase().Save(&director).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update director")
		return
	}

	json.NewEncoder(w).Encode(director)
}

// Delete Director
func deleteDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var director models.Director
	if err := models.ConnectDatabase().Where("id = ?", id).First(&director).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Director not found")
		return
	}

	models.ConnectDatabase().Delete(&director)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(director)
}
