package main

import (
	"database/sql"
	"log"
)

func MergeClusters(db *sql.DB, clusterIDs []string) error {
	//first entry becomes the new clusterID for all
	newClusterID := clusterIDs[0]
	//get all other affected faces
	var affectedFaceIDs []int
	for i := 1; i < len(clusterIDs); i++ {
		rows, err := db.Query("SELECT * FROM faces WHERE clusterID = ?", clusterIDs[i])
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
			affectedFaceIDs = append(affectedFaceIDs, newFace.faceID)
		}
	}
	//move all affected face to the new clusterID
	err := ManuallyMoveFacesToAnotherCluster(db, affectedFaceIDs, newClusterID)
	if err != nil {
		log.Println(err)
		return err
	}
	//return
	return nil
}
