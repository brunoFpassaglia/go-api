package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (u users) GetUsers(query string) ([]models.User, error) {
	query = fmt.Sprintf("%%%s%%", query)
	result, error := u.db.Query("SELECT NAME, NICK, EMAIL FROM USERS WHERE NAME LIKE ? OR NICK LIKE ?", query, query)

	if error != nil {
		return nil, error
	}
	defer result.Close()
	var users []models.User
	for result.Next() {
		var user models.User
		if error = result.Scan(&user.Name, &user.Nick, &user.Email); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}
