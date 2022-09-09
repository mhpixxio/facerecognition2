package main

import (
	"database/sql"
	"log"
)

func UpdateMeanVector(db *sql.DB, clusterID string) error {
	//get all faces of the cluster
	var facesInCluster []faceStruct
	var rows *sql.Rows
	rows, err := db.Query("SELECT * FROM faces WHERE clusterID = ?", clusterID)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	NumberOfRows := 0
	for rows.Next() {
		var newFace faceStruct
		var newVector [128]float32
		facesInCluster = append(facesInCluster, newFace)
		facesInCluster[NumberOfRows].vector = newVector[:]
		err = rows.Scan(&facesInCluster[NumberOfRows].faceID, &facesInCluster[NumberOfRows].fileID, &facesInCluster[NumberOfRows].x1, &facesInCluster[NumberOfRows].y1, &facesInCluster[NumberOfRows].x2, &facesInCluster[NumberOfRows].y2, &facesInCluster[NumberOfRows].clusterID, &facesInCluster[NumberOfRows].personName, &facesInCluster[NumberOfRows].vector[0], &facesInCluster[NumberOfRows].vector[1], &facesInCluster[NumberOfRows].vector[2], &facesInCluster[NumberOfRows].vector[3], &facesInCluster[NumberOfRows].vector[4], &facesInCluster[NumberOfRows].vector[5], &facesInCluster[NumberOfRows].vector[6], &facesInCluster[NumberOfRows].vector[7], &facesInCluster[NumberOfRows].vector[8], &facesInCluster[NumberOfRows].vector[9], &facesInCluster[NumberOfRows].vector[10], &facesInCluster[NumberOfRows].vector[11], &facesInCluster[NumberOfRows].vector[12], &facesInCluster[NumberOfRows].vector[13], &facesInCluster[NumberOfRows].vector[14], &facesInCluster[NumberOfRows].vector[15], &facesInCluster[NumberOfRows].vector[16], &facesInCluster[NumberOfRows].vector[17], &facesInCluster[NumberOfRows].vector[18], &facesInCluster[NumberOfRows].vector[19], &facesInCluster[NumberOfRows].vector[20], &facesInCluster[NumberOfRows].vector[21], &facesInCluster[NumberOfRows].vector[22], &facesInCluster[NumberOfRows].vector[23], &facesInCluster[NumberOfRows].vector[24], &facesInCluster[NumberOfRows].vector[25], &facesInCluster[NumberOfRows].vector[26], &facesInCluster[NumberOfRows].vector[27], &facesInCluster[NumberOfRows].vector[28], &facesInCluster[NumberOfRows].vector[29], &facesInCluster[NumberOfRows].vector[30], &facesInCluster[NumberOfRows].vector[31], &facesInCluster[NumberOfRows].vector[32], &facesInCluster[NumberOfRows].vector[33], &facesInCluster[NumberOfRows].vector[34], &facesInCluster[NumberOfRows].vector[35], &facesInCluster[NumberOfRows].vector[36], &facesInCluster[NumberOfRows].vector[37], &facesInCluster[NumberOfRows].vector[38], &facesInCluster[NumberOfRows].vector[39], &facesInCluster[NumberOfRows].vector[40], &facesInCluster[NumberOfRows].vector[41], &facesInCluster[NumberOfRows].vector[42], &facesInCluster[NumberOfRows].vector[43], &facesInCluster[NumberOfRows].vector[44], &facesInCluster[NumberOfRows].vector[45], &facesInCluster[NumberOfRows].vector[46], &facesInCluster[NumberOfRows].vector[47], &facesInCluster[NumberOfRows].vector[48], &facesInCluster[NumberOfRows].vector[49], &facesInCluster[NumberOfRows].vector[50], &facesInCluster[NumberOfRows].vector[51], &facesInCluster[NumberOfRows].vector[52], &facesInCluster[NumberOfRows].vector[53], &facesInCluster[NumberOfRows].vector[54], &facesInCluster[NumberOfRows].vector[55], &facesInCluster[NumberOfRows].vector[56], &facesInCluster[NumberOfRows].vector[57], &facesInCluster[NumberOfRows].vector[58], &facesInCluster[NumberOfRows].vector[59], &facesInCluster[NumberOfRows].vector[60], &facesInCluster[NumberOfRows].vector[61], &facesInCluster[NumberOfRows].vector[62], &facesInCluster[NumberOfRows].vector[63], &facesInCluster[NumberOfRows].vector[64], &facesInCluster[NumberOfRows].vector[65], &facesInCluster[NumberOfRows].vector[66], &facesInCluster[NumberOfRows].vector[67], &facesInCluster[NumberOfRows].vector[68], &facesInCluster[NumberOfRows].vector[69], &facesInCluster[NumberOfRows].vector[70], &facesInCluster[NumberOfRows].vector[71], &facesInCluster[NumberOfRows].vector[72], &facesInCluster[NumberOfRows].vector[73], &facesInCluster[NumberOfRows].vector[74], &facesInCluster[NumberOfRows].vector[75], &facesInCluster[NumberOfRows].vector[76], &facesInCluster[NumberOfRows].vector[77], &facesInCluster[NumberOfRows].vector[78], &facesInCluster[NumberOfRows].vector[79], &facesInCluster[NumberOfRows].vector[80], &facesInCluster[NumberOfRows].vector[81], &facesInCluster[NumberOfRows].vector[82], &facesInCluster[NumberOfRows].vector[83], &facesInCluster[NumberOfRows].vector[84], &facesInCluster[NumberOfRows].vector[85], &facesInCluster[NumberOfRows].vector[86], &facesInCluster[NumberOfRows].vector[87], &facesInCluster[NumberOfRows].vector[88], &facesInCluster[NumberOfRows].vector[89], &facesInCluster[NumberOfRows].vector[90], &facesInCluster[NumberOfRows].vector[91], &facesInCluster[NumberOfRows].vector[92], &facesInCluster[NumberOfRows].vector[93], &facesInCluster[NumberOfRows].vector[94], &facesInCluster[NumberOfRows].vector[95], &facesInCluster[NumberOfRows].vector[96], &facesInCluster[NumberOfRows].vector[97], &facesInCluster[NumberOfRows].vector[98], &facesInCluster[NumberOfRows].vector[99], &facesInCluster[NumberOfRows].vector[100], &facesInCluster[NumberOfRows].vector[101], &facesInCluster[NumberOfRows].vector[102], &facesInCluster[NumberOfRows].vector[103], &facesInCluster[NumberOfRows].vector[104], &facesInCluster[NumberOfRows].vector[105], &facesInCluster[NumberOfRows].vector[106], &facesInCluster[NumberOfRows].vector[107], &facesInCluster[NumberOfRows].vector[108], &facesInCluster[NumberOfRows].vector[109], &facesInCluster[NumberOfRows].vector[110], &facesInCluster[NumberOfRows].vector[111], &facesInCluster[NumberOfRows].vector[112], &facesInCluster[NumberOfRows].vector[113], &facesInCluster[NumberOfRows].vector[114], &facesInCluster[NumberOfRows].vector[115], &facesInCluster[NumberOfRows].vector[116], &facesInCluster[NumberOfRows].vector[117], &facesInCluster[NumberOfRows].vector[118], &facesInCluster[NumberOfRows].vector[119], &facesInCluster[NumberOfRows].vector[120], &facesInCluster[NumberOfRows].vector[121], &facesInCluster[NumberOfRows].vector[122], &facesInCluster[NumberOfRows].vector[123], &facesInCluster[NumberOfRows].vector[124], &facesInCluster[NumberOfRows].vector[125], &facesInCluster[NumberOfRows].vector[126], &facesInCluster[NumberOfRows].vector[127])
		if err != nil {
			log.Println(err)
			return err
		}
		NumberOfRows++
	}
	//update numberFaces in cluster
	_, err = db.Exec("UPDATE faceclusters SET numberFaces=? WHERE clusterID=?", NumberOfRows, clusterID)
	if err != nil {
		log.Println(err)
		return err
	}
	//only when the number of remaining faces is > 0 the meanVector gets updated. Otherwise it will stay the same. Clusters shouldn't be deleted because of the tree structure when searchiung for clusters
	if NumberOfRows > 0 {
		//get the current personName of the cluster to keep it the same
		var newCluster clusterStruct
		newClusterVector := [128]float32{}
		newCluster.meanVector = newClusterVector[:]
		row := db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", clusterID)
		err = row.Scan(&newCluster.clusterID, &newCluster.personName, &newCluster.numberFaces, &newCluster.meanVector[0], &newCluster.meanVector[1], &newCluster.meanVector[2], &newCluster.meanVector[3], &newCluster.meanVector[4], &newCluster.meanVector[5], &newCluster.meanVector[6], &newCluster.meanVector[7], &newCluster.meanVector[8], &newCluster.meanVector[9], &newCluster.meanVector[10], &newCluster.meanVector[11], &newCluster.meanVector[12], &newCluster.meanVector[13], &newCluster.meanVector[14], &newCluster.meanVector[15], &newCluster.meanVector[16], &newCluster.meanVector[17], &newCluster.meanVector[18], &newCluster.meanVector[19], &newCluster.meanVector[20], &newCluster.meanVector[21], &newCluster.meanVector[22], &newCluster.meanVector[23], &newCluster.meanVector[24], &newCluster.meanVector[25], &newCluster.meanVector[26], &newCluster.meanVector[27], &newCluster.meanVector[28], &newCluster.meanVector[29], &newCluster.meanVector[30], &newCluster.meanVector[31], &newCluster.meanVector[32], &newCluster.meanVector[33], &newCluster.meanVector[34], &newCluster.meanVector[35], &newCluster.meanVector[36], &newCluster.meanVector[37], &newCluster.meanVector[38], &newCluster.meanVector[39], &newCluster.meanVector[40], &newCluster.meanVector[41], &newCluster.meanVector[42], &newCluster.meanVector[43], &newCluster.meanVector[44], &newCluster.meanVector[45], &newCluster.meanVector[46], &newCluster.meanVector[47], &newCluster.meanVector[48], &newCluster.meanVector[49], &newCluster.meanVector[50], &newCluster.meanVector[51], &newCluster.meanVector[52], &newCluster.meanVector[53], &newCluster.meanVector[54], &newCluster.meanVector[55], &newCluster.meanVector[56], &newCluster.meanVector[57], &newCluster.meanVector[58], &newCluster.meanVector[59], &newCluster.meanVector[60], &newCluster.meanVector[61], &newCluster.meanVector[62], &newCluster.meanVector[63], &newCluster.meanVector[64], &newCluster.meanVector[65], &newCluster.meanVector[66], &newCluster.meanVector[67], &newCluster.meanVector[68], &newCluster.meanVector[69], &newCluster.meanVector[70], &newCluster.meanVector[71], &newCluster.meanVector[72], &newCluster.meanVector[73], &newCluster.meanVector[74], &newCluster.meanVector[75], &newCluster.meanVector[76], &newCluster.meanVector[77], &newCluster.meanVector[78], &newCluster.meanVector[79], &newCluster.meanVector[80], &newCluster.meanVector[81], &newCluster.meanVector[82], &newCluster.meanVector[83], &newCluster.meanVector[84], &newCluster.meanVector[85], &newCluster.meanVector[86], &newCluster.meanVector[87], &newCluster.meanVector[88], &newCluster.meanVector[89], &newCluster.meanVector[90], &newCluster.meanVector[91], &newCluster.meanVector[92], &newCluster.meanVector[93], &newCluster.meanVector[94], &newCluster.meanVector[95], &newCluster.meanVector[96], &newCluster.meanVector[97], &newCluster.meanVector[98], &newCluster.meanVector[99], &newCluster.meanVector[100], &newCluster.meanVector[101], &newCluster.meanVector[102], &newCluster.meanVector[103], &newCluster.meanVector[104], &newCluster.meanVector[105], &newCluster.meanVector[106], &newCluster.meanVector[107], &newCluster.meanVector[108], &newCluster.meanVector[109], &newCluster.meanVector[110], &newCluster.meanVector[111], &newCluster.meanVector[112], &newCluster.meanVector[113], &newCluster.meanVector[114], &newCluster.meanVector[115], &newCluster.meanVector[116], &newCluster.meanVector[117], &newCluster.meanVector[118], &newCluster.meanVector[119], &newCluster.meanVector[120], &newCluster.meanVector[121], &newCluster.meanVector[122], &newCluster.meanVector[123], &newCluster.meanVector[124], &newCluster.meanVector[125], &newCluster.meanVector[126], &newCluster.meanVector[127])
		if err != nil && err != sql.ErrNoRows {
			log.Printf("error: %v\n", err)
		}
		if err == sql.ErrNoRows {
			newCluster.personName = ""
		}
		//calculate the new mean vector
		var meanVector [128]float32
		for m := 0; m < 128; m++ {
			for p := 0; p < NumberOfRows; p++ {
				meanVector[m] = meanVector[m] + facesInCluster[p].vector[m]
			}
			meanVector[m] = meanVector[m] / float32(NumberOfRows)
		}
		//count the new number of faces in cluster
		count := 0
		err = db.QueryRow("SELECT COUNT(*) FROM faces WHERE clusterID = ?", clusterID).Scan(&count)
		if err != nil {
			log.Println(err)
			return err
		}
		//write the new mean vector and numberFaces in the "faceclusters" database. If the facecluster didn't exits before, it gets created here
		_, err = db.Exec("REPLACE INTO faceclusters (clusterID, personName, numberFaces, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v120, v121, v122, v123, v124, v125, v126, v127) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", clusterID, newCluster.personName, count, meanVector[0], meanVector[1], meanVector[2], meanVector[3], meanVector[4], meanVector[5], meanVector[6], meanVector[7], meanVector[8], meanVector[9], meanVector[10], meanVector[11], meanVector[12], meanVector[13], meanVector[14], meanVector[15], meanVector[16], meanVector[17], meanVector[18], meanVector[19], meanVector[20], meanVector[21], meanVector[22], meanVector[23], meanVector[24], meanVector[25], meanVector[26], meanVector[27], meanVector[28], meanVector[29], meanVector[30], meanVector[31], meanVector[32], meanVector[33], meanVector[34], meanVector[35], meanVector[36], meanVector[37], meanVector[38], meanVector[39], meanVector[40], meanVector[41], meanVector[42], meanVector[43], meanVector[44], meanVector[45], meanVector[46], meanVector[47], meanVector[48], meanVector[49], meanVector[50], meanVector[51], meanVector[52], meanVector[53], meanVector[54], meanVector[55], meanVector[56], meanVector[57], meanVector[58], meanVector[59], meanVector[60], meanVector[61], meanVector[62], meanVector[63], meanVector[64], meanVector[65], meanVector[66], meanVector[67], meanVector[68], meanVector[69], meanVector[70], meanVector[71], meanVector[72], meanVector[73], meanVector[74], meanVector[75], meanVector[76], meanVector[77], meanVector[78], meanVector[79], meanVector[80], meanVector[81], meanVector[82], meanVector[83], meanVector[84], meanVector[85], meanVector[86], meanVector[87], meanVector[88], meanVector[89], meanVector[90], meanVector[91], meanVector[92], meanVector[93], meanVector[94], meanVector[95], meanVector[96], meanVector[97], meanVector[98], meanVector[99], meanVector[100], meanVector[101], meanVector[102], meanVector[103], meanVector[104], meanVector[105], meanVector[106], meanVector[107], meanVector[108], meanVector[109], meanVector[110], meanVector[111], meanVector[112], meanVector[113], meanVector[114], meanVector[115], meanVector[116], meanVector[117], meanVector[118], meanVector[119], meanVector[120], meanVector[121], meanVector[122], meanVector[123], meanVector[124], meanVector[125], meanVector[126], meanVector[127])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
