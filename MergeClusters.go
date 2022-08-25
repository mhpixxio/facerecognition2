package main

import (
	"database/sql"
	"log"
)

func MergeClusters(db *sql.DB, clusterIDs []string) error {
	var count int
	count = 0
	for i := 0; i < len(clusterIDs); i++ {
		countNew := 0
		err := db.QueryRow("SELECT COUNT(*) FROM faceclusters WHERE clusterID = ?", clusterIDs[i]).Scan(&countNew)
		if err != nil {
			log.Println(err)
			return err
		}
		count = count + countNew
	}
	for i := 1; i < len(clusterIDs); i++ {
		//merge i in 0
		_, err := db.Exec("UPDATE faces SET clusterID=? WHERE clusterID=?", clusterIDs[0], clusterIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	log.Printf("merged %v faceclusters into 1\n", count)
	err := RemoveClustersFromDatabase(db, clusterIDs[1:])
	if err != nil {
		log.Println(err)
		return err
	}
	err = UpdateMeanVector(db, clusterIDs[0])
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
