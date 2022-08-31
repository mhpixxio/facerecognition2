package main

import (
	"database/sql"
	"log"
)

func DeleteAllPersonNames(db *sql.DB) error {

	_, err := db.Exec("UPDATE faceclusters SET personName=? WHERE personName IS NULL", "")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("UPDATE faceclusters SET personName=?", "")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("UPDATE faces SET personName=? WHERE personName IS NULL", "")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("UPDATE faces SET personName=?", "")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("deleted all personNames\n")
	return nil

}
