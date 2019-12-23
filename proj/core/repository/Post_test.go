package repository

import (
	"database/sql/driver"
	"errors"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"go-training/proj/core/model"
	"testing"
)

type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func TestPostRepository(t *testing.T) {
	Convey("Testing Post repository", t, func() {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repository := CreatePostRepository(db)
		Convey("Get all posts", func() {
			columns := []string{"Id", "CategoryId", "UserId", "Title", "Body", "Date"}
			mock.ExpectQuery("SELECT (.+) FROM posts").
				WillReturnRows(sqlmock.NewRows(columns).
					AddRow(1, 1, 1, "Title1", "Body1", "2019-12-20").
					AddRow(2, 1, 1, "Title2", "Body2", "2019-12-20").
					AddRow(3, 1, 1, "Title3", "Body3", "2019-12-20"))
			posts, err := repository.GetPostList(0)

			So(err, ShouldBeNil)
			So(len(*posts), ShouldEqual, 3)
		})
		Convey("Get category posts", func() {
			columns := []string{"Id", "CategoryId", "UserId", "Title", "Body", "Date"}
			mock.ExpectQuery("SELECT (.+) WHERE CategoryId=(.+)").
				WithArgs(1).
				WillReturnRows(sqlmock.NewRows(columns).
					AddRow(1, 1, 1, "Title1", "Body1", "2019-12-20").
					AddRow(2, 1, 1, "Title2", "Body2", "2019-12-20").
					AddRow(3, 1, 1, "Title3", "Body3", "2019-12-20"))
			posts, err := repository.GetPostList(1)

			So(err, ShouldBeNil)
			So(len(*posts), ShouldEqual, 3)
		})
		Convey("Add post", func() {
			userId := int32(1)
			post := &model.Post{
				UserId:     userId,
				CategoryId: 1,
				Title:      "Title",
				Body:       "Body",
			}

			mock.ExpectExec("INSERT INTO posts (.+) VALUES (.+)").
				WithArgs(post.CategoryId, userId, post.Title, post.Body, AnyString{}).
				WillReturnResult(sqlmock.NewResult(1, 1))
			err := repository.AddPost(post, userId)

			So(err, ShouldBeNil)
		})
		Convey("Add post returns error", func() {
			userId := int32(1)
			post := &model.Post{
				UserId:     userId,
				CategoryId: 1,
				Title:      "Title",
				Body:       "Body",
			}

			mock.ExpectExec("INSERT INTO posts (.+) VALUES (.+)").
				WillReturnError(errors.New("test"))
			err := repository.AddPost(post, userId)

			So(err, ShouldNotBeNil)
		})
	})
}
