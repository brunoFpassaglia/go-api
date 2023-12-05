package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

func (u users) CreateUser(user models.User) (uint64, error) {
	statement, error := u.db.Prepare("INSERT INTO USERS (NAME, NICK, EMAIL, PASSWORD) VALUES (?, ?, ?, ?)")

	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	id, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(id), nil
}
