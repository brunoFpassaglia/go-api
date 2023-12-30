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
func (f followers) GetFollowers(users models.User) []models.User {
	return nil
}
