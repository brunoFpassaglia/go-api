package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("criando usuarios"))

	body, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repo := repositories.NewUserRepo(db)
	repo.CreateUser(user)

}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}
func ShowUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("show usuarios"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("update usuarios"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("delete usuarios"))
}
