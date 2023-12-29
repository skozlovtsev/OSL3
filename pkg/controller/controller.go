package controller

import (
	"osl3/pkg/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tekkamanendless/go-recaptcha"
	"golang.org/x/crypto/bcrypt"
)

var (
	salt              = "whynot"
	recaptchaVerifier *recaptcha.Recaptcha
)

func init() {
	recaptchaVerifier = recaptcha.New("server-token")
}

func LogIn(c *fiber.Ctx) error {
	recaptchaResponse := c.Query("g-recaptcha-response")

	success, err := recaptchaVerifier.Verify(recaptchaResponse)
	if err != nil {
		return c.Render("login", fiber.Map{})
	}
	if !success {
		return c.Render("login", fiber.Map{})
	}

	username := c.Query("username")
	password := c.Query("password")

	password = password + salt

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return c.Render("login", fiber.Map{})
	}

	var user model.User
	user, err = model.GetUserByUsernamePassword(username, string(passHash))
	if err != nil {
		return c.Render("login", fiber.Map{})
	}

	return c.Render("page", fiber.Map{
		"Data": user.ID,
	})
}

func LogInPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func CheckUser(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Render("page", fiber.Map{})
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		return c.Render("page", fiber.Map{
			"Result": err.Error(),
		})
	}

	user, err := model.GetUserByID(i)
	if err != nil {
		return c.Render("page", fiber.Map{
			"Result": err.Error(),
		})
	}

	return c.Render("page", fiber.Map{
		"Result": user.Username + " : " + user.Password,
	})
}
