package repository

import (
	"database/sql"
	"go-training/proj/core/model"
	"time"
)

type PostRepository interface {
	GetPostList(categoryId int32) (*[]model.Post, error)
	GetPostById(id int32) (*model.Post, error)
	DeletePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	AddPost(post *model.Post, userId int32) error
}

type postRepository struct {
	db *sql.DB
}

func CreatePostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetPostList(categoryId int32) (*[]model.Post, error) {
	if categoryId == 0 {
		return r.retrievePosts("SELECT Id, CategoryId, UserId, Title, Body, Date FROM posts")
	}
	return r.retrievePosts("SELECT Id, CategoryId, UserId, Title, Body, Date FROM posts WHERE CategoryId=$1", categoryId)
}

func (r *postRepository) GetPostById(id int32) (*model.Post, error) {
	posts, err := r.retrievePosts("SELECT Id, CategoryId, UserId, Title, Body, Date FROM posts WHERE Id=$1", id)
	if err != nil {
		return nil, err
	}
	if len(*posts) > 0 {
		return &(*posts)[0], nil
	}

	return nil, nil
}

func (r *postRepository) AddPost(post *model.Post, userId int32) error {
	date := time.Now().Format(time.RFC1123)
	result, err := r.db.Exec(
		"INSERT INTO posts (CategoryId, UserId, Title, Body, Date) VALUES ($1, $2, $3, $4, $5)",
		post.CategoryId,
		userId,
		post.Title,
		post.Body,
		date,
	)

	if err != nil {
		return err
	}
	postId, _ := result.LastInsertId()
	post.Id = int32(postId)
	post.Date = date

	return nil
}

func (r *postRepository) DeletePost(post *model.Post) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE Id=$1", post.Id)

	return err
}

func (r *postRepository) UpdatePost(post *model.Post) error {
	_, err := r.db.Exec(
		"REPLACE INTO posts (Id, CategoryId, UserId, Title, Body, Date) VALUES ($1, $2, $3, $4, $5, $6)",
		post.Id,
		post.CategoryId,
		post.UserId,
		post.Title,
		post.Body,
		post.Date,
	)

	return err
}

func (r *postRepository) retrievePosts(query string, args ...interface{}) (*[]model.Post, error) {
	var posts []model.Post
	rows, err := r.db.Query(query, args...)

	if err != nil {
		return &posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err = rows.Scan(
			&post.Id,
			&post.CategoryId,
			&post.UserId,
			&post.Title,
			&post.Body,
			&post.Date,
		)
		if err != nil {
			return &posts, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}
