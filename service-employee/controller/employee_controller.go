package controller

import (
	"fmt"
	"net/http"
	"service-employee/database"
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var user_uri string = "http://localhost:3001/user"

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func CreateEmployee(c *fiber.Ctx) error {
	var requestBody model.Employee

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	requestBody.Id = uuid.New().String()

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token: Access token missing")
	}

	req, err := http.NewRequest("GET", user_uri+"/auth", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token")
	}

	db := database.GetDB()
	query := `INSERT INTO employee (id, name) VALUES ($1, $2)`
	_, err = db.Exec(query, requestBody.Id, requestBody.Name)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "unable to create employee"})
	}

	return c.Status(http.StatusCreated).JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody,
	})
}
