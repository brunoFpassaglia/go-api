package repositories

import (
	"api/src/models"
	"database/sql"
)

type followers struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) *followers {
	return &followers{db}
}

func (f followers) Follow(follower, followee uint64) error {
	statement, error := f.db.Prepare(
		"INSERT IGNORE INTO FOLLOWERS (USER_ID, FOLLOWER_ID) VALUES (?, ?)",
	)
	if error != nil {
		return error
	}

	defer statement.Close()

	_, error = statement.Exec(followee, follower)
	if error != nil {
		return error
	}

	return nil
}
func (f followers) Unfollow(follower, followee uint64) error {

	statement, error := f.db.Prepare(
		"DELETE FROM FOLLOWERS WHERE USER_ID = ? AND FOLLOWER_ID = ?",
	)
	if error != nil {
		return error
	}

	defer statement.Close()
	_, error = statement.Exec(followee, follower)
	if error != nil {
		return error
	}
	return nil

}

func (f followers) GetFollowers(userId uint64) ([]models.User, error) {
	result, error := f.db.Query(`
		SELECT U.NAME, U.NICK, U.EMAIL, U.CREATED_AT 
			FROM USERS U JOIN FOLLOWERS F ON U.ID = F.FOLLOWER_ID
			WHERE F.USER_ID = ?
	`, userId)
	if error != nil {
		return nil, error
	}
	defer result.Close()

	users := []models.User{}
	for result.Next() {
		var user models.User
		if error = result.Scan(&user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

func (f followers) GetFollowing(userId uint64) ([]models.User, error) {
	result, error := f.db.Query(`
	SELECT U.NAME, U.NICK, U.EMAIL, U.CREATED_AT
			FROM USERS U JOIN FOLLOWERS F ON U.ID = F.USER_ID 
			WHERE F.FOLLOWER_ID = ?
	`, userId)
	if error != nil {
		return nil, error
	}
	defer result.Close()

	users := []models.User{}
	for result.Next() {
		var user models.User
		if error = result.Scan(&user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}
