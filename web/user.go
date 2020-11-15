package web

import (
	"net/http"

	"github.com/jeyem/tinyurl/user"
	"github.com/labstack/echo"
)

type authForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/**
 * @api {post} /auth/register user register
 * @apiVersion 1.0.0
 * @apiName register
 * @apiGroup auth
 *
 * @apiParam {String} name fullname of user
 * @apiParam {String} email user email
 * @apiParam {String} password user password
 *
 * @apiSuccess {String} token Bearer token
 * @apiSuccess {String} message api success message
 * @apiError {String} error api error message
 *
 */

func register(c echo.Context) error {
	form := new(authForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	u, err := user.Create(txn, form.Email, form.Password, form.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	token, err := u.CreateToken(txn, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "registered successfully",
		"token":   token,
	})
}

/**
 * @api {post} /auth/login user login
 * @apiVersion 1.0.0
 * @apiName login
 * @apiGroup auth
 *
 * @apiParam {String} email user email
 * @apiParam {String} password user password
 *
 * @apiSuccess {String} token Bearer token
 * @apiSuccess {String} message api success message
 * @apiError {String} error api error message
 *
 */

func login(c echo.Context) error {
	form := new(authForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	u, err := user.Auth(txn, form.Email, form.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	token, err := u.CreateToken(txn, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "authenticated successfully",
		"token":   token,
	})
}
