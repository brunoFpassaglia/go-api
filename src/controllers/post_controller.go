package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {}
func CreatePost(c *gin.Context) {
	r := c.Request
	w := c.Writer

	idToken, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(body, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	post.AuthorId = idToken

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repo := repositories.NewPostRepo(db)
	post.ID, error = repo.CreatePost(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, post)

}
func ShowPost(c *gin.Context)   {}
func UpdatePost(c *gin.Context) {}
func DeletePost(c *gin.Context) {}
