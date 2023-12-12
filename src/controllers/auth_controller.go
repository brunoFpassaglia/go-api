package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	r := c.Request
	w := c.Writer

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var userReq models.User

	if error = json.Unmarshal(body, &userReq); error != nil {
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
	user, error := repo.GetUserByEmail(userReq.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.CheckPasswd(userReq.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	token, error := auth.MakeToken(user.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, token)

}
