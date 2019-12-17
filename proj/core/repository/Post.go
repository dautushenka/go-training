package repository

import (
	"database/sql"
	"go-training/proj/core/model"
)

type PostRepository interface {
	GetPostList() ([]model.Post, error)
	GetPostById(id int32) (*model.Post, error)
}

type postRepository struct {
	db *sql.DB
}

func CreatePostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetPostList() ([]model.Post, error) {
	posts := []model.Post{}
	rows, err := r.db.Query("SELECT Id, CategoryId, Title, Body, Date FROM posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		err = rows.Scan(
			&post.Id,
			&post.CategoryId,
			&post.Title,
			&post.Body,
			&post.Date,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *postRepository) GetPostById(id int32) (*model.Post, error) {
	return &model.Post{}, nil
}
