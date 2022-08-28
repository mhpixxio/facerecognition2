package main

import (
	"database/sql"
	"log"
)

func UpdateFacesAndClusters(db *sql.DB) error {
	//--------------------------------------- files table: go through all rows and check for files for removal -> remove corresponding faces ---------------------------------------
	rowsFiles, err := db.Query("SELECT * FROM files WHERE forRemoval = ? AND removed = ?", true, false)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rowsFiles.Close()
	for rowsFiles.Next() {
		//get fileID
		var fileReturn fileStruct
		err = rowsFiles.Scan(&fileReturn.fileID, &fileReturn.pathToFile, &fileReturn.processed, &fileReturn.forRemoval, &fileReturn.removed)
		if err != nil {
			log.Println(err)
			return err
		}
		//get affected faceIDs
		var faceReturn faceStruct
		newVector := [128]float32{}
		faceReturn.vector = newVector[:]
		rowsFaces, err := db.Query("SELECT * FROM faces WHERE fileID = ?", fileReturn.fileID)
		if err != nil {
			log.Println(err)
			return err
		}
		defer rowsFaces.Close()
		for rowsFaces.Next() {
			var faceReturn faceStruct
			newVector := [128]float32{}
			faceReturn.vector = newVector[:]
			err = rowsFaces.Scan(&faceReturn.faceID, &faceReturn.fileID, &faceReturn.x1, &faceReturn.y1, &faceReturn.x2, &faceReturn.y2, &faceReturn.clusterID, &faceReturn.personName, &faceReturn.vector[0], &faceReturn.vector[1], &faceReturn.vector[2], &faceReturn.vector[3], &faceReturn.vector[4], &faceReturn.vector[5], &faceReturn.vector[6], &faceReturn.vector[7], &faceReturn.vector[8], &faceReturn.vector[9], &faceReturn.vector[10], &faceReturn.vector[11], &faceReturn.vector[12], &faceReturn.vector[13], &faceReturn.vector[14], &faceReturn.vector[15], &faceReturn.vector[16], &faceReturn.vector[17], &faceReturn.vector[18], &faceReturn.vector[19], &faceReturn.vector[20], &faceReturn.vector[21], &faceReturn.vector[22], &faceReturn.vector[23], &faceReturn.vector[24], &faceReturn.vector[25], &faceReturn.vector[26], &faceReturn.vector[27], &faceReturn.vector[28], &faceReturn.vector[29], &faceReturn.vector[30], &faceReturn.vector[31], &faceReturn.vector[32], &faceReturn.vector[33], &faceReturn.vector[34], &faceReturn.vector[35], &faceReturn.vector[36], &faceReturn.vector[37], &faceReturn.vector[38], &faceReturn.vector[39], &faceReturn.vector[40], &faceReturn.vector[41], &faceReturn.vector[42], &faceReturn.vector[43], &faceReturn.vector[44], &faceReturn.vector[45], &faceReturn.vector[46], &faceReturn.vector[47], &faceReturn.vector[48], &faceReturn.vector[49], &faceReturn.vector[50], &faceReturn.vector[51], &faceReturn.vector[52], &faceReturn.vector[53], &faceReturn.vector[54], &faceReturn.vector[55], &faceReturn.vector[56], &faceReturn.vector[57], &faceReturn.vector[58], &faceReturn.vector[59], &faceReturn.vector[60], &faceReturn.vector[61], &faceReturn.vector[62], &faceReturn.vector[63], &faceReturn.vector[64], &faceReturn.vector[65], &faceReturn.vector[66], &faceReturn.vector[67], &faceReturn.vector[68], &faceReturn.vector[69], &faceReturn.vector[70], &faceReturn.vector[71], &faceReturn.vector[72], &faceReturn.vector[73], &faceReturn.vector[74], &faceReturn.vector[75], &faceReturn.vector[76], &faceReturn.vector[77], &faceReturn.vector[78], &faceReturn.vector[79], &faceReturn.vector[80], &faceReturn.vector[81], &faceReturn.vector[82], &faceReturn.vector[83], &faceReturn.vector[84], &faceReturn.vector[85], &faceReturn.vector[86], &faceReturn.vector[87], &faceReturn.vector[88], &faceReturn.vector[89], &faceReturn.vector[90], &faceReturn.vector[91], &faceReturn.vector[92], &faceReturn.vector[93], &faceReturn.vector[94], &faceReturn.vector[95], &faceReturn.vector[96], &faceReturn.vector[97], &faceReturn.vector[98], &faceReturn.vector[99], &faceReturn.vector[100], &faceReturn.vector[101], &faceReturn.vector[102], &faceReturn.vector[103], &faceReturn.vector[104], &faceReturn.vector[105], &faceReturn.vector[106], &faceReturn.vector[107], &faceReturn.vector[108], &faceReturn.vector[109], &faceReturn.vector[110], &faceReturn.vector[111], &faceReturn.vector[112], &faceReturn.vector[113], &faceReturn.vector[114], &faceReturn.vector[115], &faceReturn.vector[116], &faceReturn.vector[117], &faceReturn.vector[118], &faceReturn.vector[119], &faceReturn.vector[120], &faceReturn.vector[121], &faceReturn.vector[122], &faceReturn.vector[123], &faceReturn.vector[124], &faceReturn.vector[125], &faceReturn.vector[126], &faceReturn.vector[127])
			if err != nil {
				log.Println(err)
				return err
			}
			err = RemoveFacesFromDatabase(db, []int{faceReturn.faceID})
			if err != nil {
				log.Println(err)
				return err
			}
		}

	}
	//update files table
	_, err = db.Exec("UPDATE files SET removed=? WHERE forRemoval=?", true, true)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.Exec("UPDATE files SET processed=? WHERE forRemoval=?", true, true)
	if err != nil {
		log.Println(err)
		return err
	}
	//--------------------------------------- files table: go through all rows and check for unprocessed files -> do face recognition ---------------------------------------
	rows, err := db.Query("SELECT * FROM files WHERE processed = ?", false)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var fileReturn fileStruct
		err = rows.Scan(&fileReturn.fileID, &fileReturn.pathToFile, &fileReturn.processed, &fileReturn.forRemoval, &fileReturn.removed)
		if err != nil {
			log.Println(err)
			return err
		}
		//delete all existing faces of that fileID
		_, err = db.Exec("DELETE FROM faces WHERE (fileID = ?)", fileReturn.fileID)
		if err != nil {
			log.Println(err)
			return err
		}
		err = FaceRecognition(db, fileReturn.fileID)
		_, err2 := db.Exec("REPLACE INTO files (fileID, pathToFile, processed, forRemoval, removed) VALUES(?, ?, ?, ?, ?)", fileReturn.fileID, fileReturn.pathToFile, true, false, false)
		if err2 != nil {
			log.Println(err2)
			return err2
		}
		if err != nil {
			log.Println(err)
			return err
		}

	}
	//--------------------------------------- faces table: go through all rows and check for clusterID is "" or NULL-> assign cluster ---------------------------------------
	_, err = db.Exec("UPDATE faces SET clusterID=? WHERE clusterID IS NULL", "")
	if err != nil {
		log.Println(err)
		return err
	}
	rows, err = db.Query("SELECT * FROM faces WHERE clusterID = ? ORDER BY RAND()", "")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var newFace faceStruct
		var newVector [128]float32
		newFace.vector = newVector[:]
		err = rows.Scan(&newFace.faceID, &newFace.fileID, &newFace.x1, &newFace.y1, &newFace.x2, &newFace.y2, &newFace.clusterID, &newFace.personName, &newFace.vector[0], &newFace.vector[1], &newFace.vector[2], &newFace.vector[3], &newFace.vector[4], &newFace.vector[5], &newFace.vector[6], &newFace.vector[7], &newFace.vector[8], &newFace.vector[9], &newFace.vector[10], &newFace.vector[11], &newFace.vector[12], &newFace.vector[13], &newFace.vector[14], &newFace.vector[15], &newFace.vector[16], &newFace.vector[17], &newFace.vector[18], &newFace.vector[19], &newFace.vector[20], &newFace.vector[21], &newFace.vector[22], &newFace.vector[23], &newFace.vector[24], &newFace.vector[25], &newFace.vector[26], &newFace.vector[27], &newFace.vector[28], &newFace.vector[29], &newFace.vector[30], &newFace.vector[31], &newFace.vector[32], &newFace.vector[33], &newFace.vector[34], &newFace.vector[35], &newFace.vector[36], &newFace.vector[37], &newFace.vector[38], &newFace.vector[39], &newFace.vector[40], &newFace.vector[41], &newFace.vector[42], &newFace.vector[43], &newFace.vector[44], &newFace.vector[45], &newFace.vector[46], &newFace.vector[47], &newFace.vector[48], &newFace.vector[49], &newFace.vector[50], &newFace.vector[51], &newFace.vector[52], &newFace.vector[53], &newFace.vector[54], &newFace.vector[55], &newFace.vector[56], &newFace.vector[57], &newFace.vector[58], &newFace.vector[59], &newFace.vector[60], &newFace.vector[61], &newFace.vector[62], &newFace.vector[63], &newFace.vector[64], &newFace.vector[65], &newFace.vector[66], &newFace.vector[67], &newFace.vector[68], &newFace.vector[69], &newFace.vector[70], &newFace.vector[71], &newFace.vector[72], &newFace.vector[73], &newFace.vector[74], &newFace.vector[75], &newFace.vector[76], &newFace.vector[77], &newFace.vector[78], &newFace.vector[79], &newFace.vector[80], &newFace.vector[81], &newFace.vector[82], &newFace.vector[83], &newFace.vector[84], &newFace.vector[85], &newFace.vector[86], &newFace.vector[87], &newFace.vector[88], &newFace.vector[89], &newFace.vector[90], &newFace.vector[91], &newFace.vector[92], &newFace.vector[93], &newFace.vector[94], &newFace.vector[95], &newFace.vector[96], &newFace.vector[97], &newFace.vector[98], &newFace.vector[99], &newFace.vector[100], &newFace.vector[101], &newFace.vector[102], &newFace.vector[103], &newFace.vector[104], &newFace.vector[105], &newFace.vector[106], &newFace.vector[107], &newFace.vector[108], &newFace.vector[109], &newFace.vector[110], &newFace.vector[111], &newFace.vector[112], &newFace.vector[113], &newFace.vector[114], &newFace.vector[115], &newFace.vector[116], &newFace.vector[117], &newFace.vector[118], &newFace.vector[119], &newFace.vector[120], &newFace.vector[121], &newFace.vector[122], &newFace.vector[123], &newFace.vector[124], &newFace.vector[125], &newFace.vector[126], &newFace.vector[127])
		if err != nil {
			log.Println(err)
			return err
		}
		//assign cluster
		err = SearchCluster(db, newFace.faceID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//return
	return nil
}
