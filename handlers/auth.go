package handlers

import (
	"fmt"
	"time"

	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/helpers/hashhelper"
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/helpers/jwthelper"
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/models"
	"github.com/gofiber/fiber/v2"
)

func getUserByUsername(u string) (*models.User, error) {
	if u != "user1" {
		return nil, fmt.Errorf("user with same username is not found")
	}

	user := &models.User{
		UserId:    "usr-1",
		Username:  "user1",
		Password:  "$2a$10$uw412H9isn8E6/oMjg/9nu1/0Zn5J5fr.7gVme2c.7ZRGM87zesuy", //* password123
		Name:      "john doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}

func Login(c *fiber.Ctx) error {
	//* parse body request
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "error while parse body request",
			"error":   err.Error(),
		})
	}

	//* get user by username
	user, err := getUserByUsername(input.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "user not found",
			"error":   err.Error(),
		})
	}

	//* validate password
	passwordMatched := hashhelper.CompareHashString(input.Password, user.Password)
	if passwordMatched == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "password is not matching",
		})
	}

	//* generate token
	token, err := jwthelper.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to generate jwt token",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success to login",
		"token":   token,
	})
}
