package main

import (
	"database/sql"
	"log"
)

func Reclustering(db *sql.DB) error {
	//faces table: set all clusterIDs to ""
	_, err := db.Exec("UPDATE faces SET clusterID=?", "")
	if err != nil {
		log.Println(err)
		return err
	}
	//faceclusters table: delete all entries
	_, err = db.Exec("DELETE FROM faceclusters")
	if err != nil {
		log.Println(err)
		return err
	}
	//UpdateFacesAndClusters
	err = UpdateFacesAndClusters(db)
	if err != nil {
		log.Println(err)
	}
	//return
	return nil
}
