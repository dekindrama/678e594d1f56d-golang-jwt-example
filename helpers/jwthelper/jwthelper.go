package jwthelper

import (
	"time"

	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/config"
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.UserId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return t, err
	}

	return t, nil
}
