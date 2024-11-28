package handlers

import (
	"encoding/json"
	"net/http"
	"webapp/models"
	"webapp/utils"

	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Nickname string `json:"nickname"`
			Password string `json:"password"`
		}

		// Čitanje podataka iz zahteva
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Neispravan zahtev", http.StatusBadRequest)
			return
		}

		// Pronalaženje korisnika u bazi
		var user models.User
		if err := db.Where("nickname = ?", creds.Nickname).First(&user).Error; err != nil {
			http.Error(w, "Korisnik ne postoji", http.StatusUnauthorized)
			return
		}

		// Provera lozinke
		if !utils.CheckPasswordHash(creds.Password, user.Password) {
			http.Error(w, "Neispravna lozinka", http.StatusUnauthorized)
			return
		}

		// Generisanje JWT tokena
		token, err := utils.GenerateJWT(user.Nickname)
		if err != nil {
			http.Error(w, "Greška pri generisanju tokena", http.StatusInternalServerError)
			return
		}

		// Odgovor sa tokenom
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
