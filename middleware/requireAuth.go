package middleware

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"DadGpt/initializers"
	"DadGpt/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *fiber.Ctx) error {
	// Set some security headers:
	// c.Set("X-XSS-Protection", "1; mode=block")
	// c.Set("X-Content-Type-Options", "nosniff")
	// c.Set("X-Download-Options", "noopen")
	// c.Set("Strict-Transport-Security", "max-age=5184000")
	// c.Set("X-Frame-Options", "SAMEORIGIN")
	// c.Set("X-DNS-Prefetch-Control", "off")
	// fmt.Println("In middleware")

	// Get the cookie off req
	tokenString := c.Cookies("Authorization")

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration.
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Session has expired.",
			})
		}

		// Fine the user with token sub
		id := int(claims["sub"].(float64))

		var user models.User
		result := initializers.DB.First(&user, "id=?", id)
		if user.ID == 0 {
			c.Status(fiber.StatusForbidden)
			message := "Cannot find user with ID: " + strconv.Itoa(id)
			return c.JSON(fiber.Map{
				"message": message,
				"user":    user,
			})
		}
		if result.Error != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Error getting user from database.",
			})
		}

		// attach to req
		c.Locals("user", user)

		// Go to next middleware:
		return c.Next()
	} else {
		c.Status(fiber.StatusUnauthorized)
		return err
	}

}
