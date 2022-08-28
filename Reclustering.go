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
	//Update the names, i.e. try to keep the names after the Reclustering. For that purpose, the personNames get also stored in the faces table.
	//get all clusterIDs
	rows, err := db.Query("SELECT * FROM faceclusters ORDER BY numberFaces DESC")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var allClusterIDs []string
	for rows.Next() {
		var newCluster clusterStruct
		newClusterVector := [128]float32{}
		newCluster.meanVector = newClusterVector[:]
		err = rows.Scan(&newCluster.clusterID, &newCluster.personName, &newCluster.numberFaces, &newCluster.meanVector[0], &newCluster.meanVector[1], &newCluster.meanVector[2], &newCluster.meanVector[3], &newCluster.meanVector[4], &newCluster.meanVector[5], &newCluster.meanVector[6], &newCluster.meanVector[7], &newCluster.meanVector[8], &newCluster.meanVector[9], &newCluster.meanVector[10], &newCluster.meanVector[11], &newCluster.meanVector[12], &newCluster.meanVector[13], &newCluster.meanVector[14], &newCluster.meanVector[15], &newCluster.meanVector[16], &newCluster.meanVector[17], &newCluster.meanVector[18], &newCluster.meanVector[19], &newCluster.meanVector[20], &newCluster.meanVector[21], &newCluster.meanVector[22], &newCluster.meanVector[23], &newCluster.meanVector[24], &newCluster.meanVector[25], &newCluster.meanVector[26], &newCluster.meanVector[27], &newCluster.meanVector[28], &newCluster.meanVector[29], &newCluster.meanVector[30], &newCluster.meanVector[31], &newCluster.meanVector[32], &newCluster.meanVector[33], &newCluster.meanVector[34], &newCluster.meanVector[35], &newCluster.meanVector[36], &newCluster.meanVector[37], &newCluster.meanVector[38], &newCluster.meanVector[39], &newCluster.meanVector[40], &newCluster.meanVector[41], &newCluster.meanVector[42], &newCluster.meanVector[43], &newCluster.meanVector[44], &newCluster.meanVector[45], &newCluster.meanVector[46], &newCluster.meanVector[47], &newCluster.meanVector[48], &newCluster.meanVector[49], &newCluster.meanVector[50], &newCluster.meanVector[51], &newCluster.meanVector[52], &newCluster.meanVector[53], &newCluster.meanVector[54], &newCluster.meanVector[55], &newCluster.meanVector[56], &newCluster.meanVector[57], &newCluster.meanVector[58], &newCluster.meanVector[59], &newCluster.meanVector[60], &newCluster.meanVector[61], &newCluster.meanVector[62], &newCluster.meanVector[63], &newCluster.meanVector[64], &newCluster.meanVector[65], &newCluster.meanVector[66], &newCluster.meanVector[67], &newCluster.meanVector[68], &newCluster.meanVector[69], &newCluster.meanVector[70], &newCluster.meanVector[71], &newCluster.meanVector[72], &newCluster.meanVector[73], &newCluster.meanVector[74], &newCluster.meanVector[75], &newCluster.meanVector[76], &newCluster.meanVector[77], &newCluster.meanVector[78], &newCluster.meanVector[79], &newCluster.meanVector[80], &newCluster.meanVector[81], &newCluster.meanVector[82], &newCluster.meanVector[83], &newCluster.meanVector[84], &newCluster.meanVector[85], &newCluster.meanVector[86], &newCluster.meanVector[87], &newCluster.meanVector[88], &newCluster.meanVector[89], &newCluster.meanVector[90], &newCluster.meanVector[91], &newCluster.meanVector[92], &newCluster.meanVector[93], &newCluster.meanVector[94], &newCluster.meanVector[95], &newCluster.meanVector[96], &newCluster.meanVector[97], &newCluster.meanVector[98], &newCluster.meanVector[99], &newCluster.meanVector[100], &newCluster.meanVector[101], &newCluster.meanVector[102], &newCluster.meanVector[103], &newCluster.meanVector[104], &newCluster.meanVector[105], &newCluster.meanVector[106], &newCluster.meanVector[107], &newCluster.meanVector[108], &newCluster.meanVector[109], &newCluster.meanVector[110], &newCluster.meanVector[111], &newCluster.meanVector[112], &newCluster.meanVector[113], &newCluster.meanVector[114], &newCluster.meanVector[115], &newCluster.meanVector[116], &newCluster.meanVector[117], &newCluster.meanVector[118], &newCluster.meanVector[119], &newCluster.meanVector[120], &newCluster.meanVector[121], &newCluster.meanVector[122], &newCluster.meanVector[123], &newCluster.meanVector[124], &newCluster.meanVector[125], &newCluster.meanVector[126], &newCluster.meanVector[127])
		if err != nil && err != sql.ErrNoRows {
			log.Printf("error: %v\n", err)
		}
		allClusterIDs = append(allClusterIDs, newCluster.clusterID)
	}
	//go through allClusterIDs and update the personName
	for i := 0; i < len(allClusterIDs); i++ {
		//get all personNames in that cluster
		rows, err := db.Query("SELECT * FROM faces WHERE clusterID = ?", allClusterIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
		defer rows.Close()
		var personNames []string
		for rows.Next() {
			var newFace faceStruct
			var new_vector [128]float32
			newFace.vector = new_vector[:]
			err = rows.Scan(&newFace.faceID, &newFace.fileID, &newFace.x1, &newFace.y1, &newFace.x2, &newFace.y2, &newFace.clusterID, &newFace.personName, &newFace.vector[0], &newFace.vector[1], &newFace.vector[2], &newFace.vector[3], &newFace.vector[4], &newFace.vector[5], &newFace.vector[6], &newFace.vector[7], &newFace.vector[8], &newFace.vector[9], &newFace.vector[10], &newFace.vector[11], &newFace.vector[12], &newFace.vector[13], &newFace.vector[14], &newFace.vector[15], &newFace.vector[16], &newFace.vector[17], &newFace.vector[18], &newFace.vector[19], &newFace.vector[20], &newFace.vector[21], &newFace.vector[22], &newFace.vector[23], &newFace.vector[24], &newFace.vector[25], &newFace.vector[26], &newFace.vector[27], &newFace.vector[28], &newFace.vector[29], &newFace.vector[30], &newFace.vector[31], &newFace.vector[32], &newFace.vector[33], &newFace.vector[34], &newFace.vector[35], &newFace.vector[36], &newFace.vector[37], &newFace.vector[38], &newFace.vector[39], &newFace.vector[40], &newFace.vector[41], &newFace.vector[42], &newFace.vector[43], &newFace.vector[44], &newFace.vector[45], &newFace.vector[46], &newFace.vector[47], &newFace.vector[48], &newFace.vector[49], &newFace.vector[50], &newFace.vector[51], &newFace.vector[52], &newFace.vector[53], &newFace.vector[54], &newFace.vector[55], &newFace.vector[56], &newFace.vector[57], &newFace.vector[58], &newFace.vector[59], &newFace.vector[60], &newFace.vector[61], &newFace.vector[62], &newFace.vector[63], &newFace.vector[64], &newFace.vector[65], &newFace.vector[66], &newFace.vector[67], &newFace.vector[68], &newFace.vector[69], &newFace.vector[70], &newFace.vector[71], &newFace.vector[72], &newFace.vector[73], &newFace.vector[74], &newFace.vector[75], &newFace.vector[76], &newFace.vector[77], &newFace.vector[78], &newFace.vector[79], &newFace.vector[80], &newFace.vector[81], &newFace.vector[82], &newFace.vector[83], &newFace.vector[84], &newFace.vector[85], &newFace.vector[86], &newFace.vector[87], &newFace.vector[88], &newFace.vector[89], &newFace.vector[90], &newFace.vector[91], &newFace.vector[92], &newFace.vector[93], &newFace.vector[94], &newFace.vector[95], &newFace.vector[96], &newFace.vector[97], &newFace.vector[98], &newFace.vector[99], &newFace.vector[100], &newFace.vector[101], &newFace.vector[102], &newFace.vector[103], &newFace.vector[104], &newFace.vector[105], &newFace.vector[106], &newFace.vector[107], &newFace.vector[108], &newFace.vector[109], &newFace.vector[110], &newFace.vector[111], &newFace.vector[112], &newFace.vector[113], &newFace.vector[114], &newFace.vector[115], &newFace.vector[116], &newFace.vector[117], &newFace.vector[118], &newFace.vector[119], &newFace.vector[120], &newFace.vector[121], &newFace.vector[122], &newFace.vector[123], &newFace.vector[124], &newFace.vector[125], &newFace.vector[126], &newFace.vector[127])
			if err != nil {
				log.Println(err)
				return err
			}
			personNames = append(personNames, newFace.personName)
		}
		//get the most common name (must be in at least 80% of the entries)
		m := map[string]int{}
		var maxCount int
		var mostCommonName string
		for _, currentName := range personNames {
			m[currentName]++
			if m[currentName] > maxCount {
				maxCount = m[currentName]
				mostCommonName = currentName
			}
		}
		var newPersonName string
		if float32(maxCount) >= 0.8*float32(len(personNames)) {
			newPersonName = mostCommonName
		} else {
			newPersonName = ""
		}
		//overwrite personName entries in faces and faceclusters
		err = RenameCluster(db, allClusterIDs[i], newPersonName)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//return
	return nil
}
