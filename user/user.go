package user

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v2"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

// User identify client object
type User struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

// Auth authenticate user with email and password
func Auth(txn *badger.Txn, email, password string) (*User, error) {
	autherr := errors.New("email or password not matched")
	u, err := Load(txn, email)
	if err != nil {
		return nil, autherr
	}
	if ok := passwd.Check(password, u.Password); !ok {
		return nil, autherr
	}
	return u, nil
}

// CreateToken create token with request context
func (u *User) CreateToken(txn *badger.Txn, c echo.Context) (string, error) {

	claims := new(JwtClaims)
	claims.Email = u.Email
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return t, err
	}
	u.LastLoginAt = time.Now()
	u.Save(txn)
	return t, nil
}

func Create(txn *badger.Txn, email, password, name string) (*User, error) {
	if _, err := Load(txn, email); err == nil {
		return nil, errors.New("user already exists")
	}
	u := &User{
		Name:        name,
		Email:       email,
		Password:    passwd.Make(password),
		CreatedAt:   time.Now(),
		LastLoginAt: time.Now(),
	}

	if err := u.Save(txn); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Save(txn *badger.Txn) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return txn.Set([]byte(u.Email), data)
}

func Load(txn *badger.Txn, email string) (*User, error) {
	u := new(User)
	item, err := txn.Get([]byte(email))
	if err != nil {
		return nil, err
	}
	var data []byte
	result, err := item.ValueCopy(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, u); err != nil {
		return nil, err
	}
	return u, nil
}
