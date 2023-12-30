package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/repositories"
	"api/src/responses"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	followerId, error := auth.ExtractUserId(c.Request)
	if error != nil {
		responses.Error(c.Writer, http.StatusUnauthorized, error)
		return
	}

	paramValue, exists := c.Params.Get("id")
	followee, error := strconv.ParseUint(paramValue, 10, 64)
	if error != nil || !exists {
		responses.Error(c.Writer, http.StatusBadRequest, error)
		return
	}
	if followee == followerId {
		responses.Error(c.Writer, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repo := repositories.NewFollowerRepo(db)
	if error = repo.Follow(followerId, followee); error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(c.Writer, http.StatusNoContent, nil)
}
func UnFollowUser(c *gin.Context) {
	followerId, error := auth.ExtractUserId(c.Request)
	if error != nil {
		responses.Error(c.Writer, http.StatusUnauthorized, error)
		return
	}

	paramValue, exists := c.Params.Get("id")
	followee, error := strconv.ParseUint(paramValue, 10, 64)
	if error != nil || !exists {
		responses.Error(c.Writer, http.StatusBadRequest, error)
		return
	}
	if followee == followerId {
		responses.Error(c.Writer, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repo := repositories.NewFollowerRepo(db)
	if error = repo.Unfollow(followerId, followee); error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(c.Writer, http.StatusNoContent, nil)
}

func GetFollowers(c *gin.Context) {
	paramValue, exists := c.Params.Get("id")
	userId, error := strconv.ParseUint(paramValue, 10, 64)
	if error != nil || !exists {
		responses.Error(c.Writer, http.StatusBadRequest, error)
		return
	}
	db, error := database.Connect()
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repo := repositories.NewFollowerRepo(db)
	followers, error := repo.GetFollowers(userId)
	if error != nil {
		responses.Error(c.Writer, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(c.Writer, http.StatusOK, followers)

}
