package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCategoryRepository(t *testing.T) {
	Convey("Testing Post model", t, func() {
		Convey("model is invalid if category is 0", func() {
			post := Post{
				CategoryId: 0,
				Title:      "Title",
				Body:       "Body",
			}
			So(post.IsValid(), ShouldBeFalse)
		})
		Convey("model is invalid if title is empty", func() {
			post := Post{
				CategoryId: 1,
				Title:      "",
				Body:       "Body",
			}
			So(post.IsValid(), ShouldBeFalse)
		})
		Convey("model is invalid if body is empty", func() {
			post := Post{
				CategoryId: 1,
				Title:      "Title",
				Body:       "",
			}
			So(post.IsValid(), ShouldBeFalse)
		})
		Convey("model is valid if category, title and body is persist", func() {
			post := Post{
				CategoryId: 1,
				Title:      "Title",
				Body:       "Body",
			}
			So(post.IsValid(), ShouldBeTrue)
		})
	})
}
