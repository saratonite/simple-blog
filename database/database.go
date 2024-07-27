package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect(dbUrl string) *sql.DB {

	db, err := sql.Open("sqlite3", dbUrl)

	if err != nil {
		fmt.Println("Databse connection error.", err)
	}

	DB = db

	return DB

}

func InitTables() {
	_, err := DB.Exec(TABLE_QUERY)

	if err != nil {
		fmt.Println("Databse table initialization error ", err)
	}

}

const TABLE_QUERY string = `
CREATE TABLE IF NOT EXISTS posts
(
id INTEGER PRIMARY KEY AUTOINCREMENT,
slug TEXT NOT NULL UNIQUE,
title TEXT NOT NULL,
body TEXT NOT NULL
)
`
