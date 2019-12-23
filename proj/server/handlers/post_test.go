package handlers

import (
	"context"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"go-training/proj/core/model"
	"go-training/proj/core/repository/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	repositoryMock := mocks.CreatePostRepositoryMock()
	handler := CreatePostHandler(repositoryMock)
	Convey("Add post handler successful", t, func() {
		post := model.Post{
			CategoryId: 1,
			Title:      "Title2",
			Body:       "Body",
		}
		body, err := json.Marshal(post)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "/api/v1/post", strings.NewReader(string(body)))
		if err != nil {
			t.Fatal(err)
		}
		ctx := req.Context()
		ctx = context.WithValue(ctx, "User", &model.User{Id: 1})
		req = req.WithContext(ctx)
		req.Header.Set("Content-type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CreatePost)

		repositoryMock.On("AddPost", &post, int32(1)).Return(nil).Once()

		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusOK)
		//So(strings.Trim(rr.Body.String(), "\n"), ShouldEqual, "[]")
	})
	Convey("Add post handler no body", t, func() {
		req, err := http.NewRequest("POST", "/api/v1/post", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		ctx := req.Context()
		ctx = context.WithValue(ctx, "User", &model.User{Id: 1})
		req = req.WithContext(ctx)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CreatePost)

		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusBadRequest)
		// @TODO assert error message
	})
	Convey("Add post handler invalid entity", t, func() {
		post := model.Post{
			CategoryId: 0,
			Title:      "Title2",
			Body:       "",
		}
		body, err := json.Marshal(post)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "/api/v1/post", strings.NewReader(string(body)))
		if err != nil {
			t.Fatal(err)
		}
		ctx := req.Context()
		ctx = context.WithValue(ctx, "User", &model.User{Id: 1})
		req = req.WithContext(ctx)
		req.Header.Set("Content-type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CreatePost)

		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusBadRequest)
		// @TODO assert error message
	})
}
