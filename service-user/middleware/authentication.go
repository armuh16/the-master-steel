package middleware

import (
	"fmt"
	_ "service-user/config"
	"service-user/database"
	"service-user/helpers"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	access_token := c.Get("access_token")

	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(access_token)
	if err != nil {
		return c.Status(401).SendString("Invalid token: Failed to verify token")
	}

	email := checkToken["email"].(string)

	var user model.User
	db := database.GetDB()

	err = db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err, "Error fetching user from database")
		return c.Status(401).SendString("Invalid token: User not found")
	}

	// Set user data in context for future use
	c.Locals("user", user)

	// Continue processing if user is found
	return c.Next()
}
