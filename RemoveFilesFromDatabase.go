package main

import (
	"database/sql"
	"log"
)

func RemoveFilesFromDatabase(db *sql.DB, fileIDs []int) error {
	//update files table
	for i := 0; i < len(fileIDs); i++ {
		_, err := db.Exec("UPDATE files SET forRemoval=? WHERE fileID=?", true, fileIDs[i])
		_, err = db.Exec("UPDATE files SET removed=? WHERE fileID=?", false, fileIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//do the update
	err := UpdateFacesAndClusters(db)
	if err != nil {
		log.Println(err)
	}
	//return
	return nil
}
