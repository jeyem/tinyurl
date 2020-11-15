package main

import (
	"os"
	"strconv"

	"github.com/dgraph-io/badger/v2"
	"github.com/jeyem/tinyurl/web"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("test-db"))
	if err != nil {
		logrus.Fatal(err)
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port < 80 {
		port = 8000
	}
	web.Start(web.Options{
		Port: port,
		DB:   db,
	})
}
