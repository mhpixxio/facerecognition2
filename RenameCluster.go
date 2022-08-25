package main

import (
	"database/sql"
	"log"
)

func RenameCluster(db *sql.DB, clusterID string, newPersonName string) error {
	_, err := db.Exec("UPDATE faceclusters SET personName=? WHERE clusterID=?", newPersonName, clusterID)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Exec("UPDATE faces SET personName=? WHERE clusterID=?", newPersonName, clusterID)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("clusterID: %v personName changed to: %v\n", clusterID, newPersonName)
	return nil

}
