package web

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger/v2"
	"github.com/jeyem/tinyurl/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var txn *badger.Txn

// Options webserver options
type Options struct {
	Port int
	DB   *badger.DB
}

// Start webserver blocking routin
func Start(o Options) {
	e := echo.New()
	txn = o.DB.NewTransaction(true)
	go checkSignal(e)
	e.Use(middleware.Logger())
	routes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", o.Port)))

}

func routes(e *echo.Echo) {
	e.POST("/auth/login", login)
	e.POST("/auth/register", register)

	e.POST("/encode", encode, userRequired)
	e.GET("/:hash", redirect)
	e.GET("/:hash/info", info)
}

func userRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := user.LoadByRequest(txn, c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "loading user from token " + err.Error()})
		}
		c.Set("user", u)
		return next(c)
	}
}

func checkSignal(e *echo.Echo) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	txn.Commit()
	e.Close()
}
