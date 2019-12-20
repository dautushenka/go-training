package handlers

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"go-training/proj/core/model"
	"go-training/proj/core/repository/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCategoryHandler(t *testing.T) {
	repositoryMock := mocks.CreateCategoryRepositoryMock()
	handler := CreateCategoryHandler(repositoryMock)
	Convey("Get categories handler", t, func() {
		req, err := http.NewRequest("GET", "/api/v1/category", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.GetList)

		repositoryMock.On("GetCategoryList").Return(&[]model.Category{}, nil).Once()

		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusOK)
		So(strings.Trim(rr.Body.String(), "\n"), ShouldEqual, "[]")
	})
	Convey("Category handler panic if repository return error", t, func() {
		req, err := http.NewRequest("GET", "/api/v1/category", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.GetList)

		repositoryMock.On("GetCategoryList").Return(&[]model.Category{}, errors.New("test")).Once()

		So(func() { handler.ServeHTTP(rr, req) }, ShouldPanic)
	})
}
