package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectToSQLDatabase() (*sql.DB, error) {

	//---------------------------------- set the sql connection for the databases ----------------------------------
	cfg := mysql.Config{
		User:   "root",
		Passwd: "123", //security obviously has to be changed at deployment
		Net:    "tcp",
		Addr:   "host.docker.internal:3306",
		DBName: "facerecognition",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Println(pingErr)
		return nil, pingErr
	}
	return db, nil
}
