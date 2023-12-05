package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepo(db)
	user.ID, error = repo.CreateUser(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

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
