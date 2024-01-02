package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
)

type posts struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *posts {
	return &posts{db}
}

func (p *posts) GetPosts() {}
func (p *posts) CreatePost(post models.Post) (uint64, error) {
	statement, error := p.db.Prepare("INSERT INTO POSTS (TITLE, CONTENT, AUTHOR_ID) VALUES (?, ?, ?)")

	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(post.Title, post.Content, post.AuthorId)
	if error != nil {
		return 0, error
	}

	id, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(id), nil
}
func (p *posts) ShowPost(id uint64) (models.Post, error) {
	result, error := p.db.Query(
		"SELECT p.*, u.NICK FROM POSTS p join USERS u on p.AUTHOR_ID = u.ID WHERE p.ID = ?", id,
	)

	if error != nil {
		return models.Post{}, error
	}
	defer result.Close()

	var post models.Post
	if result.Next() {
		if error = result.Scan(
			&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.Likes, &post.CreatedAt, &post.AuthorNick,
		); error != nil {
			return models.Post{}, error
		} else {
			return post, nil
		}
	}
	return models.Post{}, errors.New("Not found")
}
func (p *posts) UpdatePost() {}
func (p *posts) DeletePost() {}
