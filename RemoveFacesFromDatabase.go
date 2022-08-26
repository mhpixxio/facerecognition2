package main

import (
	"database/sql"
	"log"
)

func RemoveFacesFromDatabase(db *sql.DB, faceIDs []int) error {
	var affectedClusterIDs []string
	for i := 0; i < len(faceIDs); i++ {
		var faceReturn faceStruct
		row := db.QueryRow("SELECT * FROM faces WHERE faceID = ?", faceIDs[i])
		err := row.Scan(&faceReturn.faceID, &faceReturn.fileID, &faceReturn.x1, &faceReturn.y1, &faceReturn.x2, &faceReturn.y2, &faceReturn.clusterID, &faceReturn.personName, &faceReturn.vector[0], &faceReturn.vector[1], &faceReturn.vector[2], &faceReturn.vector[3], &faceReturn.vector[4], &faceReturn.vector[5], &faceReturn.vector[6], &faceReturn.vector[7], &faceReturn.vector[8], &faceReturn.vector[9], &faceReturn.vector[10], &faceReturn.vector[11], &faceReturn.vector[12], &faceReturn.vector[13], &faceReturn.vector[14], &faceReturn.vector[15], &faceReturn.vector[16], &faceReturn.vector[17], &faceReturn.vector[18], &faceReturn.vector[19], &faceReturn.vector[20], &faceReturn.vector[21], &faceReturn.vector[22], &faceReturn.vector[23], &faceReturn.vector[24], &faceReturn.vector[25], &faceReturn.vector[26], &faceReturn.vector[27], &faceReturn.vector[28], &faceReturn.vector[29], &faceReturn.vector[30], &faceReturn.vector[31], &faceReturn.vector[32], &faceReturn.vector[33], &faceReturn.vector[34], &faceReturn.vector[35], &faceReturn.vector[36], &faceReturn.vector[37], &faceReturn.vector[38], &faceReturn.vector[39], &faceReturn.vector[40], &faceReturn.vector[41], &faceReturn.vector[42], &faceReturn.vector[43], &faceReturn.vector[44], &faceReturn.vector[45], &faceReturn.vector[46], &faceReturn.vector[47], &faceReturn.vector[48], &faceReturn.vector[49], &faceReturn.vector[50], &faceReturn.vector[51], &faceReturn.vector[52], &faceReturn.vector[53], &faceReturn.vector[54], &faceReturn.vector[55], &faceReturn.vector[56], &faceReturn.vector[57], &faceReturn.vector[58], &faceReturn.vector[59], &faceReturn.vector[60], &faceReturn.vector[61], &faceReturn.vector[62], &faceReturn.vector[63], &faceReturn.vector[64], &faceReturn.vector[65], &faceReturn.vector[66], &faceReturn.vector[67], &faceReturn.vector[68], &faceReturn.vector[69], &faceReturn.vector[70], &faceReturn.vector[71], &faceReturn.vector[72], &faceReturn.vector[73], &faceReturn.vector[74], &faceReturn.vector[75], &faceReturn.vector[76], &faceReturn.vector[77], &faceReturn.vector[78], &faceReturn.vector[79], &faceReturn.vector[80], &faceReturn.vector[81], &faceReturn.vector[82], &faceReturn.vector[83], &faceReturn.vector[84], &faceReturn.vector[85], &faceReturn.vector[86], &faceReturn.vector[87], &faceReturn.vector[88], &faceReturn.vector[89], &faceReturn.vector[90], &faceReturn.vector[91], &faceReturn.vector[92], &faceReturn.vector[93], &faceReturn.vector[94], &faceReturn.vector[95], &faceReturn.vector[96], &faceReturn.vector[97], &faceReturn.vector[98], &faceReturn.vector[99], &faceReturn.vector[100], &faceReturn.vector[101], &faceReturn.vector[102], &faceReturn.vector[103], &faceReturn.vector[104], &faceReturn.vector[105], &faceReturn.vector[106], &faceReturn.vector[107], &faceReturn.vector[108], &faceReturn.vector[109], &faceReturn.vector[110], &faceReturn.vector[111], &faceReturn.vector[112], &faceReturn.vector[113], &faceReturn.vector[114], &faceReturn.vector[115], &faceReturn.vector[116], &faceReturn.vector[117], &faceReturn.vector[118], &faceReturn.vector[119], &faceReturn.vector[120], &faceReturn.vector[121], &faceReturn.vector[122], &faceReturn.vector[123], &faceReturn.vector[124], &faceReturn.vector[125], &faceReturn.vector[126], &faceReturn.vector[127])
		affectedClusterIDs = append(affectedClusterIDs, faceReturn.clusterID)
		//delete face
		_, err = db.Exec("DELETE FROM faces WHERE fileID = ?", faceIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	// //update meanVector of clusters
	// for i := 0; i < len(affectedClusterIDs); i++ {
	// 	err := UpdateMeanVector(db, affectedClusterIDs[i])
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// }
	return nil
}
