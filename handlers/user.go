package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp-backend/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Dobavljanje svih korisnika
func GetUsersHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			http.Error(w, "Greška prilikom dobavljanja korisnika", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// Dodavanje novog korisnika
func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Neispravan zahtev", http.StatusBadRequest)
			return
		}
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Greška prilikom dodavanja korisnika", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// Ažuriranje korisnika
func UpdateUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Neispravan ID korisnika", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			http.Error(w, "Korisnik nije pronađen", http.StatusNotFound)
			return
		}

		var updates models.User
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, "Neispravan zahtev", http.StatusBadRequest)
			return
		}

		if err := db.Model(&user).Updates(updates).Error; err != nil {
			http.Error(w, "Greška prilikom ažuriranja korisnika", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Brisanje korisnika
func DeleteUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Neispravan ID korisnika", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			http.Error(w, "Korisnik nije pronađen", http.StatusNotFound)
			return
		}

		var userCount int64
		db.Model(&models.User{}).Count(&userCount)
		if userCount <= 1 {
			http.Error(w, "Ne možete obrisati poslednjeg korisnika", http.StatusForbidden)
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			http.Error(w, "Greška prilikom brisanja korisnika", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
