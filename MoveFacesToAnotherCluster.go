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
		log.Printf("faceID %v moved to clusterID %v\n", faceIDs[i], clusterID)
	}
	//update mean vector of cluster
	err := UpdateMeanVector(db, clusterID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
