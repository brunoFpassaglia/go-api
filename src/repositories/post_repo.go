package repositories

import (
	"api/src/models"
	"database/sql"
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
func (p *posts) ShowPost()   {}
func (p *posts) UpdatePost() {}
func (p *posts) DeletePost() {}
