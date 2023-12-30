package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	r := c.Request
	w := c.Writer

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
	if error = user.Prepare("create"); error != nil {
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
func GetUsers(c *gin.Context) {
	r := c.Request
	w := c.Writer
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepo(db)
	users, error := repo.GetUsers(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}
func ShowUser(c *gin.Context) {
	w := c.Writer
	paramValue, exists := c.Params.Get("id")
	id, error := strconv.ParseUint(paramValue, 10, 64)
	if error != nil || !exists {
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
	user, error := repo.ShowUser(id)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, user)

}
func UpdateUser(c *gin.Context) {
	r := c.Request
	w := c.Writer
	paramValue, exists := c.Params.Get("id")
	id, error := strconv.ParseUint(paramValue, 10, 64)
	if error != nil || !exists {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	idToken, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	if idToken != id {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot update another user"))
		return
	}

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
	user.ID = id

	if error = user.Prepare("edit"); error != nil {
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
	if error = repo.UpdateUser(user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
func DeleteUser(c *gin.Context) {
	r := c.Request
	w := c.Writer
	paramValue, exists := c.Params.Get("id")
	id, error := strconv.ParseUint(paramValue, 10, 64)

	if error != nil || !exists {
		responses.Error(c.Writer, http.StatusBadRequest, error)
	}

	idToken, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	if idToken != id {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot delete another user"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepo(db)
	error = repo.DeleteUser(id)
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(c.Writer, http.StatusNoContent, nil)

	// w.Write([]byte("delete usuarios"))
}
