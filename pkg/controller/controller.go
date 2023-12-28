package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

var salt = "whynot"
var store *session.Store

func init() {
	store = session.New()
}

func LogIn(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	att := sess.Get("attempts")
	if att == nil {
		sess.Set("attempts", 0)
	}

	attempts := sess.Get("attempts").(int)

	if attempts >= 3 {
		return err
	}

	userName := c.Query("username")
	passWord := c.Query("password")

	passWord = passWord + salt

	passHash, err := bcrypt.GenerateFromPassword([]byte(passWord), 14)

	if err != nil {
		return err
	}

	sess.Set("attempts", attempts+1)

	if err := sess.Save(); err != nil {
		return err
	}

	return c.Render("page", fiber.Map{
		"Data": userName + ":" + passWord + ":" + string(passHash),
	})
}

func LogInPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}
