package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
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
	defer db.Close()

	repo := repositories.NewUserRepo(db)
	id, error := repo.CreateUser(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("user created: %d", id)))

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
