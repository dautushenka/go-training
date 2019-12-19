package core

import (
	"database/sql"
	"github.com/icrowley/fake"
	_ "github.com/mattn/go-sqlite3"
	"go-training/proj/core/model/fixtures"
	"go-training/proj/core/security"
	"math/rand"
	"time"
)

func InitializeDB() *sql.DB {
	//db, err := sql.Open("sqlite3", "file:data.db?cache=shared&mode=memory")
	db, err := sql.Open("sqlite3", "file:data.db?cache=shared")
	if err != nil {
		panic(err)
	}

	//createSchema(db)
	//applyFixtures(db)

	return db
}

func createSchema(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
Id INTEGER PRIMARY KEY,
Login TEXT NOT NULL,
PasswordHash TEXT NOT NULL,
FirstName TEXT NOT NULL,
LastName TEXT NOT NULL
)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
Id INTEGER PRIMARY KEY,
Name TEXT NOT NULL
)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
Id INTEGER PRIMARY KEY,
CategoryId INTEGER,
UserId INTEGER,
Title TEXT NOT NULL,
Body TEXT NOT NULL,
Date TEXT NOT NULL,
   FOREIGN KEY (CategoryId) 
      REFERENCES categories (Id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
Id INTEGER PRIMARY KEY,
PostId INTEGER,
UserId INTEGER,
Body TEXT NOT NULL,
   FOREIGN KEY (PostId) 
      REFERENCES posts (Id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION,
   FOREIGN KEY (UserId) 
      REFERENCES users (Id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
)`)
	if err != nil {
		panic(err)
	}

}

func applyFixtures(db *sql.DB) {
	db.Exec("DELETE FROM users")
	for _, user := range fixtures.Users {
		_, err := db.Exec(
			"INSERT INTO users (Id, Login, PasswordHash, FirstName, LastName) VALUES ($1, $2, $3, $4, $5)",
			user.Id,
			user.Login,
			security.EncodePassword(user.Login),
			user.FirstName,
			user.LastName,
		)
		if err != nil {
			panic(err)
		}
	}

	db.Exec("DELETE FROM categories")
	for _, category := range fixtures.Categories {
		_, err := db.Exec(
			"INSERT INTO categories (Id, Name) VALUES ($1, $2)",
			category.Id,
			category.Name,
		)
		if err != nil {
			panic(err)
		}
	}

	db.Exec("DELETE FROM posts")
	insertPost := func(Id int32, CategoryId int32, Title string, Body string, Date string) {
		var postId interface{}
		if Id != 0 {
			postId = Id
		}
		_, err := db.Exec(
			"INSERT INTO posts (Id, CategoryId, UserId, Title, Body, Date) VALUES ($1, $2, $3, $4, $5, $6)",
			postId,
			CategoryId,
			rand.Intn(2)+1,
			Title,
			Body,
			Date,
		)

		if err != nil {
			panic(err)
		}
	}

	for _, post := range fixtures.Posts {
		insertPost(post.Id,
			post.CategoryId,
			post.Title,
			post.Body,
			time.Now().String(),
		)
	}

	for i := 0; i < 100; i++ {
		insertPost(0,
			int32(rand.Intn(3)+1),
			fake.Title(),
			fake.Paragraphs(),
			time.Now().Add(time.Hour*24*10*-1).String(),
		)
	}
}
