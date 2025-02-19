package jwthelper

import (
	"fmt"
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

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	//* Parse dan verifikasi token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Memastikan bahwa metode signing yang digunakan adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config("SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	//* Pastikan token valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	//* Ambil klaim dari token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	//* return claims
	return *claims, nil
}
