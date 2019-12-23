package repository

import (
	"errors"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCategoryRepository(t *testing.T) {
	Convey("Testing Category repository", t, func() {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repository := CreateCategoryRepository(db)
		Convey("Testing GetCategoryList", func() {
			columns := []string{"Id", "Name"}
			mock.ExpectQuery("SELECT Id, Name from categories").WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Name"))
			categories, err := repository.GetCategoryList()

			So(err, ShouldBeNil)
			So(len(*categories), ShouldEqual, 1)
		})
		Convey("GetCategoryList returns error", func() {
			errorMessage := "These is no category table"
			mock.ExpectQuery("SELECT Id, Name from categories").WillReturnError(errors.New(errorMessage))
			categories, err := repository.GetCategoryList()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, errorMessage)
			So(len(*categories), ShouldEqual, 0)
		})
	})
}
