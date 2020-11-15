package main

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/jeyem/tinyurl/web"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("test-db"))
	if err != nil {
		logrus.Fatal(err)
	}
	web.Start(web.Options{
		Port: 8000,
		DB:   db,
	})
}
