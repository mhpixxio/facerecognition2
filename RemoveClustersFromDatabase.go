package main

import (
	"database/sql"
	"log"
)

//removes the given faces from the sql database "faces"
func RemoveClustersFromDatabase(db *sql.DB, clusterIDs []string) (err_return error) {
	var count int
	count = 0
	for i := 0; i < len(clusterIDs); i++ {
		count_new := 0
		err := db.QueryRow("SELECT COUNT(*) FROM faceclusters WHERE clusterID = ?", clusterIDs[i]).Scan(&count_new)
		if err != nil {
			log.Println(err)
			return err
		}
		count = count + count_new
		_, err = db.Exec("DELETE FROM faceclusters WHERE (clusterID = ?)", clusterIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
		//remove the clusterIDs of the corresponding entries in the "faces" database
		_, err = db.Exec("UPDATE faces SET clusterID=? WHERE clusterID=?", "", clusterIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//log.Printf("removed %v faceclusters from the database\n", count)
	return nil
}
