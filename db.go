package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
