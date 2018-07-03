package config

import (
	"gopkg.in/gorethink/gorethink.v4"
	"log"
)


func RethinkSession() *gorethink.Session{
	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address: "localhost:32769",
		Database: "shopping",
	})
	if err !=nil {
		log.Fatal(err.Error())
	}

	return session
}