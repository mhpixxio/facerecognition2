package main

import (
	"database/sql"
	"log"
)

func MoveFacesToAnotherCluster(db *sql.DB, faceIDs []int, clusterID string) error {
	for i := 0; i < len(faceIDs); i++ {
		_, err := db.Exec("UPDATE faces SET clusterID=? WHERE faceID=?", clusterID, faceIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
