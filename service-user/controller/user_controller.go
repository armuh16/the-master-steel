package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"service-user/database"
	"service-user/helpers"
	"service-user/model"
)

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User
	db := database.GetDB()

	requestBody.Id = uuid.New().String()

	c.BodyParser(&requestBody)

	stmt, err := db.Prepare("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err // Handle error
	}
	defer stmt.Close()

	_, err = stmt.Exec(requestBody.Id, requestBody.Email, helpers.HashPassword([]byte(requestBody.Password)))
	if err != nil {
		return err // Handle error
	}

	return c.JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody.Email,
	})
}

func Login(c *fiber.Ctx) error {
	var requestBody model.User
	db := database.GetDB()

	c.BodyParser(&requestBody)

	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", requestBody.Email).Scan(&requestBody.Id, &requestBody.Email, &requestBody.Password)
	if err != nil {
		return c.JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(requestBody.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct {
		Code        int
		Status      string
		AccessToken string
		Data        interface{}
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        requestBody,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
