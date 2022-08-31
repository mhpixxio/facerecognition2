package main

import (
	"database/sql"
	"log"
	"math"
)

//thresholds when comparing the Squared Euclidean Distance between two face vectors
const threshold1 float64 = 0.3   //goes into cluster for more comparison
const threshold21 float64 = 0.29 //comparison of individual vectors. depends on the number of faces already in cluster. For 1 face in cluster the threshold is threshold21. For threshold2range number of files, the threshold is threshold22. Values between are linear.
const threshold22 float64 = 0.24
const threshold2range float64 = 10
const maxThresholdInCluster float64 = 0.6 //max distance between two vectors in cluster

func SearchCluster(db *sql.DB, faceID int) error {
	//set variables
	var faceReturn faceStruct
	newVector := [128]float32{}
	faceReturn.vector = newVector[:]
	newClusterID := ""
	//get face from database
	row := db.QueryRow("SELECT * FROM faces WHERE faceID = ?", faceID)
	err := row.Scan(&faceReturn.faceID, &faceReturn.fileID, &faceReturn.x1, &faceReturn.y1, &faceReturn.x2, &faceReturn.y2, &faceReturn.clusterID, &faceReturn.personName, &faceReturn.vector[0], &faceReturn.vector[1], &faceReturn.vector[2], &faceReturn.vector[3], &faceReturn.vector[4], &faceReturn.vector[5], &faceReturn.vector[6], &faceReturn.vector[7], &faceReturn.vector[8], &faceReturn.vector[9], &faceReturn.vector[10], &faceReturn.vector[11], &faceReturn.vector[12], &faceReturn.vector[13], &faceReturn.vector[14], &faceReturn.vector[15], &faceReturn.vector[16], &faceReturn.vector[17], &faceReturn.vector[18], &faceReturn.vector[19], &faceReturn.vector[20], &faceReturn.vector[21], &faceReturn.vector[22], &faceReturn.vector[23], &faceReturn.vector[24], &faceReturn.vector[25], &faceReturn.vector[26], &faceReturn.vector[27], &faceReturn.vector[28], &faceReturn.vector[29], &faceReturn.vector[30], &faceReturn.vector[31], &faceReturn.vector[32], &faceReturn.vector[33], &faceReturn.vector[34], &faceReturn.vector[35], &faceReturn.vector[36], &faceReturn.vector[37], &faceReturn.vector[38], &faceReturn.vector[39], &faceReturn.vector[40], &faceReturn.vector[41], &faceReturn.vector[42], &faceReturn.vector[43], &faceReturn.vector[44], &faceReturn.vector[45], &faceReturn.vector[46], &faceReturn.vector[47], &faceReturn.vector[48], &faceReturn.vector[49], &faceReturn.vector[50], &faceReturn.vector[51], &faceReturn.vector[52], &faceReturn.vector[53], &faceReturn.vector[54], &faceReturn.vector[55], &faceReturn.vector[56], &faceReturn.vector[57], &faceReturn.vector[58], &faceReturn.vector[59], &faceReturn.vector[60], &faceReturn.vector[61], &faceReturn.vector[62], &faceReturn.vector[63], &faceReturn.vector[64], &faceReturn.vector[65], &faceReturn.vector[66], &faceReturn.vector[67], &faceReturn.vector[68], &faceReturn.vector[69], &faceReturn.vector[70], &faceReturn.vector[71], &faceReturn.vector[72], &faceReturn.vector[73], &faceReturn.vector[74], &faceReturn.vector[75], &faceReturn.vector[76], &faceReturn.vector[77], &faceReturn.vector[78], &faceReturn.vector[79], &faceReturn.vector[80], &faceReturn.vector[81], &faceReturn.vector[82], &faceReturn.vector[83], &faceReturn.vector[84], &faceReturn.vector[85], &faceReturn.vector[86], &faceReturn.vector[87], &faceReturn.vector[88], &faceReturn.vector[89], &faceReturn.vector[90], &faceReturn.vector[91], &faceReturn.vector[92], &faceReturn.vector[93], &faceReturn.vector[94], &faceReturn.vector[95], &faceReturn.vector[96], &faceReturn.vector[97], &faceReturn.vector[98], &faceReturn.vector[99], &faceReturn.vector[100], &faceReturn.vector[101], &faceReturn.vector[102], &faceReturn.vector[103], &faceReturn.vector[104], &faceReturn.vector[105], &faceReturn.vector[106], &faceReturn.vector[107], &faceReturn.vector[108], &faceReturn.vector[109], &faceReturn.vector[110], &faceReturn.vector[111], &faceReturn.vector[112], &faceReturn.vector[113], &faceReturn.vector[114], &faceReturn.vector[115], &faceReturn.vector[116], &faceReturn.vector[117], &faceReturn.vector[118], &faceReturn.vector[119], &faceReturn.vector[120], &faceReturn.vector[121], &faceReturn.vector[122], &faceReturn.vector[123], &faceReturn.vector[124], &faceReturn.vector[125], &faceReturn.vector[126], &faceReturn.vector[127])
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("faceById %d: no such face\n", faceID)
			return err
		}
		log.Printf("faceById %d: %v\n", faceID, err)
		return err
	}
	v := faceReturn.vector //v = the vector of the face currently used
	//check if cluster with that name already exists
	var newCluster clusterStruct
	newMeanVectorCluster := [128]float32{}
	newCluster.meanVector = newMeanVectorCluster[:]
	row = db.QueryRow("SELECT * FROM faceclusters WHERE personName = ?", faceReturn.personName)
	err = row.Scan(&newCluster.clusterID, &newCluster.personName, &newCluster.numberFaces, &newCluster.meanVector[0], &newCluster.meanVector[1], &newCluster.meanVector[2], &newCluster.meanVector[3], &newCluster.meanVector[4], &newCluster.meanVector[5], &newCluster.meanVector[6], &newCluster.meanVector[7], &newCluster.meanVector[8], &newCluster.meanVector[9], &newCluster.meanVector[10], &newCluster.meanVector[11], &newCluster.meanVector[12], &newCluster.meanVector[13], &newCluster.meanVector[14], &newCluster.meanVector[15], &newCluster.meanVector[16], &newCluster.meanVector[17], &newCluster.meanVector[18], &newCluster.meanVector[19], &newCluster.meanVector[20], &newCluster.meanVector[21], &newCluster.meanVector[22], &newCluster.meanVector[23], &newCluster.meanVector[24], &newCluster.meanVector[25], &newCluster.meanVector[26], &newCluster.meanVector[27], &newCluster.meanVector[28], &newCluster.meanVector[29], &newCluster.meanVector[30], &newCluster.meanVector[31], &newCluster.meanVector[32], &newCluster.meanVector[33], &newCluster.meanVector[34], &newCluster.meanVector[35], &newCluster.meanVector[36], &newCluster.meanVector[37], &newCluster.meanVector[38], &newCluster.meanVector[39], &newCluster.meanVector[40], &newCluster.meanVector[41], &newCluster.meanVector[42], &newCluster.meanVector[43], &newCluster.meanVector[44], &newCluster.meanVector[45], &newCluster.meanVector[46], &newCluster.meanVector[47], &newCluster.meanVector[48], &newCluster.meanVector[49], &newCluster.meanVector[50], &newCluster.meanVector[51], &newCluster.meanVector[52], &newCluster.meanVector[53], &newCluster.meanVector[54], &newCluster.meanVector[55], &newCluster.meanVector[56], &newCluster.meanVector[57], &newCluster.meanVector[58], &newCluster.meanVector[59], &newCluster.meanVector[60], &newCluster.meanVector[61], &newCluster.meanVector[62], &newCluster.meanVector[63], &newCluster.meanVector[64], &newCluster.meanVector[65], &newCluster.meanVector[66], &newCluster.meanVector[67], &newCluster.meanVector[68], &newCluster.meanVector[69], &newCluster.meanVector[70], &newCluster.meanVector[71], &newCluster.meanVector[72], &newCluster.meanVector[73], &newCluster.meanVector[74], &newCluster.meanVector[75], &newCluster.meanVector[76], &newCluster.meanVector[77], &newCluster.meanVector[78], &newCluster.meanVector[79], &newCluster.meanVector[80], &newCluster.meanVector[81], &newCluster.meanVector[82], &newCluster.meanVector[83], &newCluster.meanVector[84], &newCluster.meanVector[85], &newCluster.meanVector[86], &newCluster.meanVector[87], &newCluster.meanVector[88], &newCluster.meanVector[89], &newCluster.meanVector[90], &newCluster.meanVector[91], &newCluster.meanVector[92], &newCluster.meanVector[93], &newCluster.meanVector[94], &newCluster.meanVector[95], &newCluster.meanVector[96], &newCluster.meanVector[97], &newCluster.meanVector[98], &newCluster.meanVector[99], &newCluster.meanVector[100], &newCluster.meanVector[101], &newCluster.meanVector[102], &newCluster.meanVector[103], &newCluster.meanVector[104], &newCluster.meanVector[105], &newCluster.meanVector[106], &newCluster.meanVector[107], &newCluster.meanVector[108], &newCluster.meanVector[109], &newCluster.meanVector[110], &newCluster.meanVector[111], &newCluster.meanVector[112], &newCluster.meanVector[113], &newCluster.meanVector[114], &newCluster.meanVector[115], &newCluster.meanVector[116], &newCluster.meanVector[117], &newCluster.meanVector[118], &newCluster.meanVector[119], &newCluster.meanVector[120], &newCluster.meanVector[121], &newCluster.meanVector[122], &newCluster.meanVector[123], &newCluster.meanVector[124], &newCluster.meanVector[125], &newCluster.meanVector[126], &newCluster.meanVector[127])
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error: %v\n", err)
		return err
	}
	if err != sql.ErrNoRows && faceReturn.personName != "" {
		//move the face to that cluster
		err = MoveFacesToAnotherCluster(db, []int{faceID}, newCluster.clusterID)
		if err != nil {
			log.Println(err)
			return err
		}
		//return
		return nil
	}
	//get possible clusters
	possibleClusters, endpoints, err := GetPossibleClusters(db, v, threshold1, "1")
	if err != nil {
		log.Println(err)
		return err
	}
	//narrow down the clusters
	var distances []float64
	for i := 0; i < len(possibleClusters); i++ {
		newMinDistance, _, err := CheckAllVectorsInCluster(db, possibleClusters[i], v, threshold21, threshold22, threshold2range, maxThresholdInCluster)
		if err != nil {
			log.Println(err)
			return err
		}
		distances = append(distances, newMinDistance)
	}
	//look for the one with the closest match
	var minDistance float64
	minDistance = 99
	minPosition := 0
	if len(distances) > 0 {
		minDistance = distances[0]
		for i := 0; i < len(distances); i++ {
			if distances[i] < minDistance {
				minDistance = distances[i]
				minPosition = i
			}
		}
	}
	//check if one was found
	if minDistance < 99 {
		newClusterID = possibleClusters[minPosition]
	} else {
		//if none was found, check wich endpoint from the first selection is the closest and create new cluster there
		var endpointDistances []float64
		for i := 0; i < len(endpoints); i++ {
			var tempCluster clusterStruct
			newVectorTempCluster := [128]float32{}
			tempCluster.meanVector = newVectorTempCluster[:]
			row := db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", endpoints[i])
			err := row.Scan(&tempCluster.clusterID, &tempCluster.personName, &tempCluster.numberFaces, &tempCluster.meanVector[0], &tempCluster.meanVector[1], &tempCluster.meanVector[2], &tempCluster.meanVector[3], &tempCluster.meanVector[4], &tempCluster.meanVector[5], &tempCluster.meanVector[6], &tempCluster.meanVector[7], &tempCluster.meanVector[8], &tempCluster.meanVector[9], &tempCluster.meanVector[10], &tempCluster.meanVector[11], &tempCluster.meanVector[12], &tempCluster.meanVector[13], &tempCluster.meanVector[14], &tempCluster.meanVector[15], &tempCluster.meanVector[16], &tempCluster.meanVector[17], &tempCluster.meanVector[18], &tempCluster.meanVector[19], &tempCluster.meanVector[20], &tempCluster.meanVector[21], &tempCluster.meanVector[22], &tempCluster.meanVector[23], &tempCluster.meanVector[24], &tempCluster.meanVector[25], &tempCluster.meanVector[26], &tempCluster.meanVector[27], &tempCluster.meanVector[28], &tempCluster.meanVector[29], &tempCluster.meanVector[30], &tempCluster.meanVector[31], &tempCluster.meanVector[32], &tempCluster.meanVector[33], &tempCluster.meanVector[34], &tempCluster.meanVector[35], &tempCluster.meanVector[36], &tempCluster.meanVector[37], &tempCluster.meanVector[38], &tempCluster.meanVector[39], &tempCluster.meanVector[40], &tempCluster.meanVector[41], &tempCluster.meanVector[42], &tempCluster.meanVector[43], &tempCluster.meanVector[44], &tempCluster.meanVector[45], &tempCluster.meanVector[46], &tempCluster.meanVector[47], &tempCluster.meanVector[48], &tempCluster.meanVector[49], &tempCluster.meanVector[50], &tempCluster.meanVector[51], &tempCluster.meanVector[52], &tempCluster.meanVector[53], &tempCluster.meanVector[54], &tempCluster.meanVector[55], &tempCluster.meanVector[56], &tempCluster.meanVector[57], &tempCluster.meanVector[58], &tempCluster.meanVector[59], &tempCluster.meanVector[60], &tempCluster.meanVector[61], &tempCluster.meanVector[62], &tempCluster.meanVector[63], &tempCluster.meanVector[64], &tempCluster.meanVector[65], &tempCluster.meanVector[66], &tempCluster.meanVector[67], &tempCluster.meanVector[68], &tempCluster.meanVector[69], &tempCluster.meanVector[70], &tempCluster.meanVector[71], &tempCluster.meanVector[72], &tempCluster.meanVector[73], &tempCluster.meanVector[74], &tempCluster.meanVector[75], &tempCluster.meanVector[76], &tempCluster.meanVector[77], &tempCluster.meanVector[78], &tempCluster.meanVector[79], &tempCluster.meanVector[80], &tempCluster.meanVector[81], &tempCluster.meanVector[82], &tempCluster.meanVector[83], &tempCluster.meanVector[84], &tempCluster.meanVector[85], &tempCluster.meanVector[86], &tempCluster.meanVector[87], &tempCluster.meanVector[88], &tempCluster.meanVector[89], &tempCluster.meanVector[90], &tempCluster.meanVector[91], &tempCluster.meanVector[92], &tempCluster.meanVector[93], &tempCluster.meanVector[94], &tempCluster.meanVector[95], &tempCluster.meanVector[96], &tempCluster.meanVector[97], &tempCluster.meanVector[98], &tempCluster.meanVector[99], &tempCluster.meanVector[100], &tempCluster.meanVector[101], &tempCluster.meanVector[102], &tempCluster.meanVector[103], &tempCluster.meanVector[104], &tempCluster.meanVector[105], &tempCluster.meanVector[106], &tempCluster.meanVector[107], &tempCluster.meanVector[108], &tempCluster.meanVector[109], &tempCluster.meanVector[110], &tempCluster.meanVector[111], &tempCluster.meanVector[112], &tempCluster.meanVector[113], &tempCluster.meanVector[114], &tempCluster.meanVector[115], &tempCluster.meanVector[116], &tempCluster.meanVector[117], &tempCluster.meanVector[118], &tempCluster.meanVector[119], &tempCluster.meanVector[120], &tempCluster.meanVector[121], &tempCluster.meanVector[122], &tempCluster.meanVector[123], &tempCluster.meanVector[124], &tempCluster.meanVector[125], &tempCluster.meanVector[126], &tempCluster.meanVector[127])
			if err != nil && err != sql.ErrNoRows {
				log.Printf("error: %v\n", err)
				return err
			}
			if err == sql.ErrNoRows { //check if cluster exists
				endpointDistances = append(endpointDistances, 99)
			} else {
				endpointDistances = append(endpointDistances, SquaredEuclideanDistance(v, tempCluster.meanVector))
			}
		}
		//check minimum endpoint distance
		minEndpointDistance := endpointDistances[0]
		minEndpointPosition := 0
		for i := 0; i < len(endpointDistances); i++ {
			if endpointDistances[i] < minEndpointDistance {
				minEndpointDistance = endpointDistances[i]
				minEndpointPosition = i
			}
		}
		//create new cluster at the closest endpoint
		//check if cluster 2 exists next to it
		clusterID2 := endpoints[minEndpointPosition][:len(endpoints[minEndpointPosition])-1] + "2"
		var cluster2 clusterStruct
		newVectorCluster2 := [128]float32{}
		cluster2.meanVector = newVectorCluster2[:]
		row = db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", clusterID2)
		err = row.Scan(&cluster2.clusterID, &cluster2.personName, &cluster2.numberFaces, &cluster2.meanVector[0], &cluster2.meanVector[1], &cluster2.meanVector[2], &cluster2.meanVector[3], &cluster2.meanVector[4], &cluster2.meanVector[5], &cluster2.meanVector[6], &cluster2.meanVector[7], &cluster2.meanVector[8], &cluster2.meanVector[9], &cluster2.meanVector[10], &cluster2.meanVector[11], &cluster2.meanVector[12], &cluster2.meanVector[13], &cluster2.meanVector[14], &cluster2.meanVector[15], &cluster2.meanVector[16], &cluster2.meanVector[17], &cluster2.meanVector[18], &cluster2.meanVector[19], &cluster2.meanVector[20], &cluster2.meanVector[21], &cluster2.meanVector[22], &cluster2.meanVector[23], &cluster2.meanVector[24], &cluster2.meanVector[25], &cluster2.meanVector[26], &cluster2.meanVector[27], &cluster2.meanVector[28], &cluster2.meanVector[29], &cluster2.meanVector[30], &cluster2.meanVector[31], &cluster2.meanVector[32], &cluster2.meanVector[33], &cluster2.meanVector[34], &cluster2.meanVector[35], &cluster2.meanVector[36], &cluster2.meanVector[37], &cluster2.meanVector[38], &cluster2.meanVector[39], &cluster2.meanVector[40], &cluster2.meanVector[41], &cluster2.meanVector[42], &cluster2.meanVector[43], &cluster2.meanVector[44], &cluster2.meanVector[45], &cluster2.meanVector[46], &cluster2.meanVector[47], &cluster2.meanVector[48], &cluster2.meanVector[49], &cluster2.meanVector[50], &cluster2.meanVector[51], &cluster2.meanVector[52], &cluster2.meanVector[53], &cluster2.meanVector[54], &cluster2.meanVector[55], &cluster2.meanVector[56], &cluster2.meanVector[57], &cluster2.meanVector[58], &cluster2.meanVector[59], &cluster2.meanVector[60], &cluster2.meanVector[61], &cluster2.meanVector[62], &cluster2.meanVector[63], &cluster2.meanVector[64], &cluster2.meanVector[65], &cluster2.meanVector[66], &cluster2.meanVector[67], &cluster2.meanVector[68], &cluster2.meanVector[69], &cluster2.meanVector[70], &cluster2.meanVector[71], &cluster2.meanVector[72], &cluster2.meanVector[73], &cluster2.meanVector[74], &cluster2.meanVector[75], &cluster2.meanVector[76], &cluster2.meanVector[77], &cluster2.meanVector[78], &cluster2.meanVector[79], &cluster2.meanVector[80], &cluster2.meanVector[81], &cluster2.meanVector[82], &cluster2.meanVector[83], &cluster2.meanVector[84], &cluster2.meanVector[85], &cluster2.meanVector[86], &cluster2.meanVector[87], &cluster2.meanVector[88], &cluster2.meanVector[89], &cluster2.meanVector[90], &cluster2.meanVector[91], &cluster2.meanVector[92], &cluster2.meanVector[93], &cluster2.meanVector[94], &cluster2.meanVector[95], &cluster2.meanVector[96], &cluster2.meanVector[97], &cluster2.meanVector[98], &cluster2.meanVector[99], &cluster2.meanVector[100], &cluster2.meanVector[101], &cluster2.meanVector[102], &cluster2.meanVector[103], &cluster2.meanVector[104], &cluster2.meanVector[105], &cluster2.meanVector[106], &cluster2.meanVector[107], &cluster2.meanVector[108], &cluster2.meanVector[109], &cluster2.meanVector[110], &cluster2.meanVector[111], &cluster2.meanVector[112], &cluster2.meanVector[113], &cluster2.meanVector[114], &cluster2.meanVector[115], &cluster2.meanVector[116], &cluster2.meanVector[117], &cluster2.meanVector[118], &cluster2.meanVector[119], &cluster2.meanVector[120], &cluster2.meanVector[121], &cluster2.meanVector[122], &cluster2.meanVector[123], &cluster2.meanVector[124], &cluster2.meanVector[125], &cluster2.meanVector[126], &cluster2.meanVector[127])
		if err != nil && err != sql.ErrNoRows {
			log.Printf("error: %v\n", err)
			return err
		}
		if err == sql.ErrNoRows && endpoints[minEndpointPosition] != "1" { //check if cluster exsists
			newClusterID = clusterID2
		} else {
			newClusterID = endpoints[minEndpointPosition] + "1"
		}
	}
	//move to new cluster
	err = MoveFacesToAnotherCluster(db, []int{faceID}, newClusterID)
	if err != nil {
		log.Println(err)
		return err
	}
	//if this face was the first entry to the cluster, and the face has a personName -> rename cluster to that personName. At "UpdateFacesAndClusters.go" it always starts with the face entries which have personNames (important for this step!). The number of faces in the cluster gets updated in "UpdateMeanVector.go"
	if faceReturn.personName != "" {
		var newNewCluster clusterStruct
		newVectorCluster := [128]float32{}
		newNewCluster.meanVector = newVectorCluster[:]
		row = db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", newClusterID)
		err = row.Scan(&newNewCluster.clusterID, &newNewCluster.personName, &newNewCluster.numberFaces, &newNewCluster.meanVector[0], &newNewCluster.meanVector[1], &newNewCluster.meanVector[2], &newNewCluster.meanVector[3], &newNewCluster.meanVector[4], &newNewCluster.meanVector[5], &newNewCluster.meanVector[6], &newNewCluster.meanVector[7], &newNewCluster.meanVector[8], &newNewCluster.meanVector[9], &newNewCluster.meanVector[10], &newNewCluster.meanVector[11], &newNewCluster.meanVector[12], &newNewCluster.meanVector[13], &newNewCluster.meanVector[14], &newNewCluster.meanVector[15], &newNewCluster.meanVector[16], &newNewCluster.meanVector[17], &newNewCluster.meanVector[18], &newNewCluster.meanVector[19], &newNewCluster.meanVector[20], &newNewCluster.meanVector[21], &newNewCluster.meanVector[22], &newNewCluster.meanVector[23], &newNewCluster.meanVector[24], &newNewCluster.meanVector[25], &newNewCluster.meanVector[26], &newNewCluster.meanVector[27], &newNewCluster.meanVector[28], &newNewCluster.meanVector[29], &newNewCluster.meanVector[30], &newNewCluster.meanVector[31], &newNewCluster.meanVector[32], &newNewCluster.meanVector[33], &newNewCluster.meanVector[34], &newNewCluster.meanVector[35], &newNewCluster.meanVector[36], &newNewCluster.meanVector[37], &newNewCluster.meanVector[38], &newNewCluster.meanVector[39], &newNewCluster.meanVector[40], &newNewCluster.meanVector[41], &newNewCluster.meanVector[42], &newNewCluster.meanVector[43], &newNewCluster.meanVector[44], &newNewCluster.meanVector[45], &newNewCluster.meanVector[46], &newNewCluster.meanVector[47], &newNewCluster.meanVector[48], &newNewCluster.meanVector[49], &newNewCluster.meanVector[50], &newNewCluster.meanVector[51], &newNewCluster.meanVector[52], &newNewCluster.meanVector[53], &newNewCluster.meanVector[54], &newNewCluster.meanVector[55], &newNewCluster.meanVector[56], &newNewCluster.meanVector[57], &newNewCluster.meanVector[58], &newNewCluster.meanVector[59], &newNewCluster.meanVector[60], &newNewCluster.meanVector[61], &newNewCluster.meanVector[62], &newNewCluster.meanVector[63], &newNewCluster.meanVector[64], &newNewCluster.meanVector[65], &newNewCluster.meanVector[66], &newNewCluster.meanVector[67], &newNewCluster.meanVector[68], &newNewCluster.meanVector[69], &newNewCluster.meanVector[70], &newNewCluster.meanVector[71], &newNewCluster.meanVector[72], &newNewCluster.meanVector[73], &newNewCluster.meanVector[74], &newNewCluster.meanVector[75], &newNewCluster.meanVector[76], &newNewCluster.meanVector[77], &newNewCluster.meanVector[78], &newNewCluster.meanVector[79], &newNewCluster.meanVector[80], &newNewCluster.meanVector[81], &newNewCluster.meanVector[82], &newNewCluster.meanVector[83], &newNewCluster.meanVector[84], &newNewCluster.meanVector[85], &newNewCluster.meanVector[86], &newNewCluster.meanVector[87], &newNewCluster.meanVector[88], &newNewCluster.meanVector[89], &newNewCluster.meanVector[90], &newNewCluster.meanVector[91], &newNewCluster.meanVector[92], &newNewCluster.meanVector[93], &newNewCluster.meanVector[94], &newNewCluster.meanVector[95], &newNewCluster.meanVector[96], &newNewCluster.meanVector[97], &newNewCluster.meanVector[98], &newNewCluster.meanVector[99], &newNewCluster.meanVector[100], &newNewCluster.meanVector[101], &newNewCluster.meanVector[102], &newNewCluster.meanVector[103], &newNewCluster.meanVector[104], &newNewCluster.meanVector[105], &newNewCluster.meanVector[106], &newNewCluster.meanVector[107], &newNewCluster.meanVector[108], &newNewCluster.meanVector[109], &newNewCluster.meanVector[110], &newNewCluster.meanVector[111], &newNewCluster.meanVector[112], &newNewCluster.meanVector[113], &newNewCluster.meanVector[114], &newNewCluster.meanVector[115], &newNewCluster.meanVector[116], &newNewCluster.meanVector[117], &newNewCluster.meanVector[118], &newNewCluster.meanVector[119], &newNewCluster.meanVector[120], &newNewCluster.meanVector[121], &newNewCluster.meanVector[122], &newNewCluster.meanVector[123], &newNewCluster.meanVector[124], &newNewCluster.meanVector[125], &newNewCluster.meanVector[126], &newNewCluster.meanVector[127])
		if err != nil && err != sql.ErrNoRows {
			log.Printf("error: %v\n", err)
			return err
		}
		if newNewCluster.numberFaces == 1 {
			err = RenameCluster(db, newClusterID, faceReturn.personName)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	//return
	return nil
}

func GetPossibleClusters(db *sql.DB, v []float32, threshold1 float64, currentPath string) (possibleClusters []string, endpoints []string, errReturn error) {
	var clusterID1, clusterID2 string
	var distance1, distance2 float64
	//get distance1
	clusterID1 = currentPath + "1"
	var cluster1 clusterStruct
	newVectorCluster1 := [128]float32{}
	cluster1.meanVector = newVectorCluster1[:]
	row := db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", clusterID1)
	err := row.Scan(&cluster1.clusterID, &cluster1.personName, &cluster1.numberFaces, &cluster1.meanVector[0], &cluster1.meanVector[1], &cluster1.meanVector[2], &cluster1.meanVector[3], &cluster1.meanVector[4], &cluster1.meanVector[5], &cluster1.meanVector[6], &cluster1.meanVector[7], &cluster1.meanVector[8], &cluster1.meanVector[9], &cluster1.meanVector[10], &cluster1.meanVector[11], &cluster1.meanVector[12], &cluster1.meanVector[13], &cluster1.meanVector[14], &cluster1.meanVector[15], &cluster1.meanVector[16], &cluster1.meanVector[17], &cluster1.meanVector[18], &cluster1.meanVector[19], &cluster1.meanVector[20], &cluster1.meanVector[21], &cluster1.meanVector[22], &cluster1.meanVector[23], &cluster1.meanVector[24], &cluster1.meanVector[25], &cluster1.meanVector[26], &cluster1.meanVector[27], &cluster1.meanVector[28], &cluster1.meanVector[29], &cluster1.meanVector[30], &cluster1.meanVector[31], &cluster1.meanVector[32], &cluster1.meanVector[33], &cluster1.meanVector[34], &cluster1.meanVector[35], &cluster1.meanVector[36], &cluster1.meanVector[37], &cluster1.meanVector[38], &cluster1.meanVector[39], &cluster1.meanVector[40], &cluster1.meanVector[41], &cluster1.meanVector[42], &cluster1.meanVector[43], &cluster1.meanVector[44], &cluster1.meanVector[45], &cluster1.meanVector[46], &cluster1.meanVector[47], &cluster1.meanVector[48], &cluster1.meanVector[49], &cluster1.meanVector[50], &cluster1.meanVector[51], &cluster1.meanVector[52], &cluster1.meanVector[53], &cluster1.meanVector[54], &cluster1.meanVector[55], &cluster1.meanVector[56], &cluster1.meanVector[57], &cluster1.meanVector[58], &cluster1.meanVector[59], &cluster1.meanVector[60], &cluster1.meanVector[61], &cluster1.meanVector[62], &cluster1.meanVector[63], &cluster1.meanVector[64], &cluster1.meanVector[65], &cluster1.meanVector[66], &cluster1.meanVector[67], &cluster1.meanVector[68], &cluster1.meanVector[69], &cluster1.meanVector[70], &cluster1.meanVector[71], &cluster1.meanVector[72], &cluster1.meanVector[73], &cluster1.meanVector[74], &cluster1.meanVector[75], &cluster1.meanVector[76], &cluster1.meanVector[77], &cluster1.meanVector[78], &cluster1.meanVector[79], &cluster1.meanVector[80], &cluster1.meanVector[81], &cluster1.meanVector[82], &cluster1.meanVector[83], &cluster1.meanVector[84], &cluster1.meanVector[85], &cluster1.meanVector[86], &cluster1.meanVector[87], &cluster1.meanVector[88], &cluster1.meanVector[89], &cluster1.meanVector[90], &cluster1.meanVector[91], &cluster1.meanVector[92], &cluster1.meanVector[93], &cluster1.meanVector[94], &cluster1.meanVector[95], &cluster1.meanVector[96], &cluster1.meanVector[97], &cluster1.meanVector[98], &cluster1.meanVector[99], &cluster1.meanVector[100], &cluster1.meanVector[101], &cluster1.meanVector[102], &cluster1.meanVector[103], &cluster1.meanVector[104], &cluster1.meanVector[105], &cluster1.meanVector[106], &cluster1.meanVector[107], &cluster1.meanVector[108], &cluster1.meanVector[109], &cluster1.meanVector[110], &cluster1.meanVector[111], &cluster1.meanVector[112], &cluster1.meanVector[113], &cluster1.meanVector[114], &cluster1.meanVector[115], &cluster1.meanVector[116], &cluster1.meanVector[117], &cluster1.meanVector[118], &cluster1.meanVector[119], &cluster1.meanVector[120], &cluster1.meanVector[121], &cluster1.meanVector[122], &cluster1.meanVector[123], &cluster1.meanVector[124], &cluster1.meanVector[125], &cluster1.meanVector[126], &cluster1.meanVector[127])
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error: %v\n", err)
		return nil, nil, err
	}
	if err == sql.ErrNoRows { //check if cluster exsists
		distance1 = 99
	} else {
		distance1 = SquaredEuclideanDistance(v, cluster1.meanVector)
	}
	//get distance2
	clusterID2 = currentPath + "2"
	var cluster2 clusterStruct
	newVectorCluster2 := [128]float32{}
	cluster2.meanVector = newVectorCluster2[:]
	row = db.QueryRow("SELECT * FROM faceclusters WHERE clusterID = ?", clusterID2)
	err = row.Scan(&cluster2.clusterID, &cluster2.personName, &cluster2.numberFaces, &cluster2.meanVector[0], &cluster2.meanVector[1], &cluster2.meanVector[2], &cluster2.meanVector[3], &cluster2.meanVector[4], &cluster2.meanVector[5], &cluster2.meanVector[6], &cluster2.meanVector[7], &cluster2.meanVector[8], &cluster2.meanVector[9], &cluster2.meanVector[10], &cluster2.meanVector[11], &cluster2.meanVector[12], &cluster2.meanVector[13], &cluster2.meanVector[14], &cluster2.meanVector[15], &cluster2.meanVector[16], &cluster2.meanVector[17], &cluster2.meanVector[18], &cluster2.meanVector[19], &cluster2.meanVector[20], &cluster2.meanVector[21], &cluster2.meanVector[22], &cluster2.meanVector[23], &cluster2.meanVector[24], &cluster2.meanVector[25], &cluster2.meanVector[26], &cluster2.meanVector[27], &cluster2.meanVector[28], &cluster2.meanVector[29], &cluster2.meanVector[30], &cluster2.meanVector[31], &cluster2.meanVector[32], &cluster2.meanVector[33], &cluster2.meanVector[34], &cluster2.meanVector[35], &cluster2.meanVector[36], &cluster2.meanVector[37], &cluster2.meanVector[38], &cluster2.meanVector[39], &cluster2.meanVector[40], &cluster2.meanVector[41], &cluster2.meanVector[42], &cluster2.meanVector[43], &cluster2.meanVector[44], &cluster2.meanVector[45], &cluster2.meanVector[46], &cluster2.meanVector[47], &cluster2.meanVector[48], &cluster2.meanVector[49], &cluster2.meanVector[50], &cluster2.meanVector[51], &cluster2.meanVector[52], &cluster2.meanVector[53], &cluster2.meanVector[54], &cluster2.meanVector[55], &cluster2.meanVector[56], &cluster2.meanVector[57], &cluster2.meanVector[58], &cluster2.meanVector[59], &cluster2.meanVector[60], &cluster2.meanVector[61], &cluster2.meanVector[62], &cluster2.meanVector[63], &cluster2.meanVector[64], &cluster2.meanVector[65], &cluster2.meanVector[66], &cluster2.meanVector[67], &cluster2.meanVector[68], &cluster2.meanVector[69], &cluster2.meanVector[70], &cluster2.meanVector[71], &cluster2.meanVector[72], &cluster2.meanVector[73], &cluster2.meanVector[74], &cluster2.meanVector[75], &cluster2.meanVector[76], &cluster2.meanVector[77], &cluster2.meanVector[78], &cluster2.meanVector[79], &cluster2.meanVector[80], &cluster2.meanVector[81], &cluster2.meanVector[82], &cluster2.meanVector[83], &cluster2.meanVector[84], &cluster2.meanVector[85], &cluster2.meanVector[86], &cluster2.meanVector[87], &cluster2.meanVector[88], &cluster2.meanVector[89], &cluster2.meanVector[90], &cluster2.meanVector[91], &cluster2.meanVector[92], &cluster2.meanVector[93], &cluster2.meanVector[94], &cluster2.meanVector[95], &cluster2.meanVector[96], &cluster2.meanVector[97], &cluster2.meanVector[98], &cluster2.meanVector[99], &cluster2.meanVector[100], &cluster2.meanVector[101], &cluster2.meanVector[102], &cluster2.meanVector[103], &cluster2.meanVector[104], &cluster2.meanVector[105], &cluster2.meanVector[106], &cluster2.meanVector[107], &cluster2.meanVector[108], &cluster2.meanVector[109], &cluster2.meanVector[110], &cluster2.meanVector[111], &cluster2.meanVector[112], &cluster2.meanVector[113], &cluster2.meanVector[114], &cluster2.meanVector[115], &cluster2.meanVector[116], &cluster2.meanVector[117], &cluster2.meanVector[118], &cluster2.meanVector[119], &cluster2.meanVector[120], &cluster2.meanVector[121], &cluster2.meanVector[122], &cluster2.meanVector[123], &cluster2.meanVector[124], &cluster2.meanVector[125], &cluster2.meanVector[126], &cluster2.meanVector[127])
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error: %v\n", err)
		return nil, nil, err
	}
	if err == sql.ErrNoRows { //check if cluster exsists
		distance2 = 99
	} else {
		distance2 = SquaredEuclideanDistance(v, cluster2.meanVector)
	}

	//check if distances are small enough to append to possibleClusters
	if distance1 < threshold1 {
		possibleClusters = append(possibleClusters, currentPath+"1")
	}
	if distance2 < threshold1 {
		possibleClusters = append(possibleClusters, currentPath+"2")
	}

	//check if endpoints
	if distance1 == 99 {
		endpoints = append(endpoints, currentPath)
		return possibleClusters, endpoints, nil
	} else {
		if distance2 == 99 {
			endpoints = append(endpoints, currentPath+"1")
			return possibleClusters, endpoints, nil
		}
	}

	//check if distance1 and distance2 are too close to another. If so, check both possible ways
	if distance1 < 99 && math.Abs(distance2-distance1) < threshold1 {
		newPossibleClusters, newEndpoints, err := GetPossibleClusters(db, v, threshold1, currentPath+"1")
		if err != nil {
			log.Printf("error: %v\n", err)
			return nil, nil, err
		}
		possibleClusters = append(possibleClusters, newPossibleClusters...)
		endpoints = append(endpoints, newEndpoints...)
		newPossibleClusters, newEndpoints, err = GetPossibleClusters(db, v, threshold1, currentPath+"2")
		if err != nil {
			log.Printf("error: %v\n", err)
			return nil, nil, err
		}
		possibleClusters = append(possibleClusters, newPossibleClusters...)
		endpoints = append(endpoints, newEndpoints...)
	} else {
		//otherwiese just check the closer way
		if distance1 < distance2 {
			newPossibleClusters, newEndpoints, err := GetPossibleClusters(db, v, threshold1, currentPath+"1")
			if err != nil {
				log.Printf("error: %v\n", err)
				return nil, nil, err
			}
			possibleClusters = append(possibleClusters, newPossibleClusters...)
			endpoints = append(endpoints, newEndpoints...)
		} else {
			newPossibleClusters, newEndpoints, err := GetPossibleClusters(db, v, threshold1, currentPath+"2")
			if err != nil {
				log.Printf("error: %v\n", err)
				return nil, nil, err
			}
			possibleClusters = append(possibleClusters, newPossibleClusters...)
			endpoints = append(endpoints, newEndpoints...)
		}
	}

	return possibleClusters, endpoints, nil
}

func SquaredEuclideanDistance(vector1 []float32, vector2 []float32) (sum float64) {
	for i := range vector1 {
		sum = sum + math.Pow(float64(vector2[i]-vector1[i]), 2)
	}
	return sum
}

func CheckAllVectorsInCluster(db *sql.DB, clusterID string, vector []float32, threshold21 float64, threshold22 float64, threshold2range float64, maxThresholdInCluster float64) (minDistanceReturn float64, maxDistanceReturn float64, errReturn error) {
	//count faces already in cluster
	count := 0
	err := db.QueryRow("SELECT COUNT(*) FROM faces WHERE clusterID = ?", clusterID).Scan(&count)
	if err != nil {
		log.Println(err)
		return 99, 99, err
	}
	//calculate threshold2 depending on the number of faces already in cluster
	var threshold2 float64
	if float64(count) <= 1 {
		threshold2 = threshold21
	} else {
		if float64(count) >= threshold2range {
			threshold2 = threshold22
		} else {
			threshold2 = (float64(count)/threshold2range)*(threshold22-threshold21) + threshold21
		}
	}
	if count > 0 {
		//get all faces in clusters and check the smallest distance to any face
		rows, err := db.Query("SELECT * FROM faces WHERE clusterID = ?", clusterID)
		if err != nil {
			log.Println(err)
			return 99, 99, err
		}
		defer rows.Close()
		var distances []float64
		for rows.Next() {
			var newFace faceStruct
			var newVector [128]float32
			newFace.vector = newVector[:]
			err = rows.Scan(&newFace.faceID, &newFace.fileID, &newFace.x1, &newFace.y1, &newFace.x2, &newFace.y2, &newFace.clusterID, &newFace.personName, &newFace.vector[0], &newFace.vector[1], &newFace.vector[2], &newFace.vector[3], &newFace.vector[4], &newFace.vector[5], &newFace.vector[6], &newFace.vector[7], &newFace.vector[8], &newFace.vector[9], &newFace.vector[10], &newFace.vector[11], &newFace.vector[12], &newFace.vector[13], &newFace.vector[14], &newFace.vector[15], &newFace.vector[16], &newFace.vector[17], &newFace.vector[18], &newFace.vector[19], &newFace.vector[20], &newFace.vector[21], &newFace.vector[22], &newFace.vector[23], &newFace.vector[24], &newFace.vector[25], &newFace.vector[26], &newFace.vector[27], &newFace.vector[28], &newFace.vector[29], &newFace.vector[30], &newFace.vector[31], &newFace.vector[32], &newFace.vector[33], &newFace.vector[34], &newFace.vector[35], &newFace.vector[36], &newFace.vector[37], &newFace.vector[38], &newFace.vector[39], &newFace.vector[40], &newFace.vector[41], &newFace.vector[42], &newFace.vector[43], &newFace.vector[44], &newFace.vector[45], &newFace.vector[46], &newFace.vector[47], &newFace.vector[48], &newFace.vector[49], &newFace.vector[50], &newFace.vector[51], &newFace.vector[52], &newFace.vector[53], &newFace.vector[54], &newFace.vector[55], &newFace.vector[56], &newFace.vector[57], &newFace.vector[58], &newFace.vector[59], &newFace.vector[60], &newFace.vector[61], &newFace.vector[62], &newFace.vector[63], &newFace.vector[64], &newFace.vector[65], &newFace.vector[66], &newFace.vector[67], &newFace.vector[68], &newFace.vector[69], &newFace.vector[70], &newFace.vector[71], &newFace.vector[72], &newFace.vector[73], &newFace.vector[74], &newFace.vector[75], &newFace.vector[76], &newFace.vector[77], &newFace.vector[78], &newFace.vector[79], &newFace.vector[80], &newFace.vector[81], &newFace.vector[82], &newFace.vector[83], &newFace.vector[84], &newFace.vector[85], &newFace.vector[86], &newFace.vector[87], &newFace.vector[88], &newFace.vector[89], &newFace.vector[90], &newFace.vector[91], &newFace.vector[92], &newFace.vector[93], &newFace.vector[94], &newFace.vector[95], &newFace.vector[96], &newFace.vector[97], &newFace.vector[98], &newFace.vector[99], &newFace.vector[100], &newFace.vector[101], &newFace.vector[102], &newFace.vector[103], &newFace.vector[104], &newFace.vector[105], &newFace.vector[106], &newFace.vector[107], &newFace.vector[108], &newFace.vector[109], &newFace.vector[110], &newFace.vector[111], &newFace.vector[112], &newFace.vector[113], &newFace.vector[114], &newFace.vector[115], &newFace.vector[116], &newFace.vector[117], &newFace.vector[118], &newFace.vector[119], &newFace.vector[120], &newFace.vector[121], &newFace.vector[122], &newFace.vector[123], &newFace.vector[124], &newFace.vector[125], &newFace.vector[126], &newFace.vector[127])
			if err != nil {
				log.Println(err)
				return 99, 99, err
			}
			distances = append(distances, SquaredEuclideanDistance(vector, newFace.vector))
		}
		//get smallest distance to any face in cluster
		minDistance := distances[0]
		for i := 0; i < len(distances); i++ {
			if distances[i] < minDistance {
				minDistance = distances[i]
			}
		}
		//get largest distance to any face in cluster
		maxDistance := distances[0]
		for i := 0; i < len(distances); i++ {
			if distances[i] > maxDistance {
				maxDistance = distances[i]
			}
		}
		//check thresholds and return values
		if minDistance <= threshold2 && maxDistance <= maxThresholdInCluster {
			return minDistance, maxDistance, nil
		}
	}
	//else return 99 99
	return 99, 99, nil
}
