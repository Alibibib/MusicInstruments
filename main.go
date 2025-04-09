package main

import (
	"log"
	"net/http"

	"MusicInstruments/database"
	"MusicInstruments/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.AddUserHandler).Methods("POST")
	r.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	// Запускаем HTTP сервер
	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
