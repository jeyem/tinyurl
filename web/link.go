package web

import (
	"net/http"

	"time"

	"github.com/jeyem/tinyurl/link"
	"github.com/jeyem/tinyurl/user"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type encodeForm struct {
	Link   string    `json:"link"`
	Expire time.Time `json:"expire"`
}

/**
 * @api {post} /encode encode url
 * @apiVersion 1.0.0
 * @apiName encode
 * @apiGroup link
 *
 * @apiParam {String} link orginal link to be shorten
 * @apiParam {String} expire expire time compatible with all standard layouts
 *
 * @apiSuccess {String} link  export shorten link of url
 * @apiError {String} error api error message
 *
 */

func encode(c echo.Context) error {
	u := c.Get("user").(*user.User)
	form := new(encodeForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	l, err := link.Create(txn, form.Link, u.Email, form.Expire)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	req := c.Request()
	schema := "http://"
	if req.TLS != nil {
		schema = "https://"
	}
	return c.JSON(http.StatusAccepted, echo.Map{"link": l.Shorten(schema + req.Host)})
}

func redirect(c echo.Context) error {
	l, err := link.Load(txn, c.Param("hash"), true)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	l.Visited++
	l.Save(txn)
	logrus.Info(l.Orginal)
	return c.Redirect(http.StatusSeeOther, l.Orginal)
}

/**
 * @api {post} /:hash/info link info
 * @apiVersion 1.0.0
 * @apiName info
 * @apiGroup link
 *
 * @apiParam {URLParam} hash hash of encoded url
 *
 * @apiError {String} error api error message
 *
 */

func info(c echo.Context) error {
	l, err := link.Load(txn, c.Param("hash"), false)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, l)
}
