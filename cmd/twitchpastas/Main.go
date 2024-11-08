package main

import (
	"database/sql"
	"log"
)

type DBConfig struct {
	login    string
	password string
}

var users = map[string]DBConfig{}

func main() {
	users["Ravendza"] = DBConfig{login: "Shtopor", password: "sosal"}

	db, err := sql.Open("sqlite3", "./fiber.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
