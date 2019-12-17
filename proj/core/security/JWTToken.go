package security

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"go-training/proj/core/model"
	"strconv"
	"time"
)

const SECRET_KEY = "ergermm325423fsdfsdfs"

func GetToken(user *model.User) string {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Id:        strconv.Itoa(int(user.Id)),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Go-Proj",
		ExpiresAt: time.Now().Unix() + 7*24*3600,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		panic(err)
	}

	return ss
}

func ValidateToken(tokenStr string) (int32, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if token.Valid {
		claims := token.Claims.(*jwt.StandardClaims)
		fmt.Println(claims.Issuer)
		id, _ := strconv.ParseInt(claims.Id, 10, 32)
		return int32(id), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return 0, errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, errors.New("Timing is everything")
		} else {
			return 0, errors.New("Couldn't handle this token:")
		}
	} else {
		return 0, errors.New("Couldn't handle this token:")
	}
}
