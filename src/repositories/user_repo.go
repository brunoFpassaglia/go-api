package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
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

func (u users) ShowUser(id uint64) (models.User, error) {
	result, error := u.db.Query(
		"SELECT ID, NAME, NICK, EMAIL FROM USERS WHERE ID = ?", id,
	)

	if error != nil {
		return models.User{}, error
	}

	var user models.User
	if result.Next() {
		if error = result.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email,
		); error != nil {
			return models.User{}, error
		} else {
			return user, nil
		}
	}
	return models.User{}, errors.New("Not found")
}

func (u users) UpdateUser(user models.User) error {
	statement, error := u.db.Prepare("UPDATE USERS set NAME = ?, NICK=?, EMAIL=? where ID = ?")

	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.ID)
	if error != nil {
		return error
	}

	_, error = result.RowsAffected()
	if error != nil {
		return error
	}
	return nil
}
func (u users) DeleteUser(id uint64) error {
	statement, error := u.db.Prepare(
		"DELETE FROM USERS WHERE ID = ?",
	)
	defer statement.Close()
	result, error := statement.Exec(id)
	if error != nil {
		return error
	}

	_, error = result.RowsAffected()
	if error != nil {
		return error
	}
	return nil
}
