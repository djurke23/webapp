package main

import (
	"log"
	"net/http"
	"webapp/database"
	"webapp/handlers"
	"webapp/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Povezivanje sa bazom podataka
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Greška prilikom povezivanja sa bazom:", err)
	}

	// Pristup sql.DB objektu za zatvaranje konekcije
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Greška prilikom pristupa sql.DB:", err)
	}
	defer sqlDB.Close()

	// Kreiranje router-a
	router := mux.NewRouter()

	// Javni endpointi (bez autentifikacije)
	router.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")

	// Zaštićeni endpointi
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/users", handlers.GetUsersHandler(db)).Methods("GET")
	api.HandleFunc("/users", handlers.CreateUserHandler(db)).Methods("POST")
	api.HandleFunc("/users/{id}", handlers.UpdateUserHandler(db)).Methods("PUT")
	api.HandleFunc("/users/{id}", handlers.DeleteUserHandler(db)).Methods("DELETE")

	// Pokretanje servera
	log.Println("Server je pokrenut na http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
