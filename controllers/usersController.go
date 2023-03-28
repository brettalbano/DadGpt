package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/brettalbano/DadGpt/initializers"
	"github.com/brettalbano/DadGpt/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserIndex(c *fiber.Ctx) error {
	return c.Render("users/index", fiber.Map{
		"hello": "world",
	})
}

func RegisterUser(c *fiber.Ctx) error {
	// Get the info of the user (username, password, open_ai_key)
	var body struct {
		Username             string `json:"username"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirm"`
		OpenAiKey            string `json:"open_ai_key"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if body.Password != body.PasswordConfirmation {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match.",
		})
	}

	// Hash the password.
	cost, _ := strconv.Atoi(os.Getenv("COST"))
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), cost)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Cannot encrypt password.",
		})
	}

	// Create the user.
	if body.OpenAiKey == "" {
		body.OpenAiKey = os.Getenv("OPENAIKEY")
	}
	user := models.User{
		UserName:  body.Username,
		Password:  string(hash),
		OpenAiKey: body.OpenAiKey,
	}
	result := initializers.DB.Create(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Respond.
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Registered New User",
		"user":    user,
	})
}

func LoginUser(c *fiber.Ctx) error {
	// Get username/password off req body.
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	// Get user from DB.
	var user models.User
	result := initializers.DB.First(&user, "user_name=?", body.Username)
	if user.ID == 0 {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Cannot find user with given username/password",
			"user":    user,
		})
	}
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error getting user from database.",
		})
	}

	// Verify password.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Cannot find user with given username/password",
		})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    user.ID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
		"ai_key": user.OpenAiKey,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error while creating token.",
		})
	}

	// Create cookie.
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 2)
	cookie.HTTPOnly = true
	cookie.SameSite = "Strict"

	c.Cookie(cookie)

	c.Status(fiber.StatusAccepted)
	return c.JSON(fiber.Map{
		"Message": "Successfully logged in.",
		"Token":   tokenString,
		"UserId":  user.ID,
	})
}

func LogoutUser(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success!",
	})

}

func ValidateUser(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(fiber.Map{
		"message": "Got user.",
		"user":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var userList []models.User

	result := initializers.DB.Find(&userList)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": userList,
	})
}
