package user

import (
	"errors"
	"strings"

	"github.com/dgraph-io/badger/v2"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const jwtSecret = "This is secret should be in config"

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func LoadByRequest(txn *badger.Txn, c echo.Context) (*User, error) {
	token := GetToken(c)
	return LoadByToken(txn, token)
}

func LoadByToken(txn *badger.Txn, tokenStr string) (*User, error) {
	claims := &JwtClaims{}
	token, _ := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if token == nil || !token.Valid {
		return nil, errors.New("token not valid")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("converting token failed")
	}

	return Load(txn, claims.Email)
}

func GetToken(c echo.Context) string {
	req := c.Request()
	cleared := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	return strings.Replace(cleared, "Bearer", "", -1)
}
