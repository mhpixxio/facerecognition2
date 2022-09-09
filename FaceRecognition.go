package main

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	face "github.com/Kagami/go-face"
)

func FaceRecognition(db *sql.DB, fileID int) (numberOfFaces int, errReturn error) {
	// start := time.Now()
	//get pathtofile
	var fileReturn fileStruct
	row := db.QueryRow("SELECT * FROM files WHERE fileID = ?", fileID)
	err := row.Scan(&fileReturn.fileID, &fileReturn.pathToFile, &fileReturn.processed, &fileReturn.forRemoval, &fileReturn.removed)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("fileId %d: no such file\n", fileID)
			return 0, err
		}
		return 0, err
	}
	pathToFile := fileReturn.pathToFile
	//check file size
	fi, err := os.Stat(pathToFile)
	if err != nil {
		log.Printf("%v", err)
		return 0, err
	}
	if fi.Size() > 100*1000000 {
		err = errors.New("file too big")
		log.Printf("%v, %v bytes", err, fi.Size())
		return 0, err
	}
	//check the format, and convert if necessary
	format := filepath.Ext(pathToFile)
	var bytesFile []byte
	if format == ".jpg" || format == ".jepg" {
		bytesFile, err = ioutil.ReadFile(pathToFile)
		if err != nil {
			return 0, err
		}
	} else {
		os.Mkdir("./tempfiles/", os.ModePerm)
		args := []string{"-quality", "90", pathToFile + "[0]", "-background", "white", "-alpha", "remove", "./tempfiles/tempfile.jpg"}
		cmd := exec.Command("convert", args...)
		_, err = cmd.Output()
		if err != nil {
			log.Println(fi.Size())
			return 0, err
		}
		bytesFile, err = ioutil.ReadFile("./tempfiles/tempfile.jpg")
		if err != nil {
			return 0, err
		}
		err = os.Remove("./tempfiles/tempfile.jpg")
		if err != nil {
			return 0, err
		}
	}
	//set variables
	var x1, y1, x2, y2 int
	var faceID int
	var v []float32
	var faces []faceStruct
	var newFace faceStruct
	//set the model for the recognizer
	rec, err := face.NewRecognizer("models")
	if err != nil {
		return 0, err
	}
	//run face recognization
	recognizedfaces, err := rec.Recognize(bytesFile)
	bytesFile = nil //free memory
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(recognizedfaces); i++ {
		faceID = fileID*1000 + i
		x1 = recognizedfaces[i].Rectangle.Min.X
		y1 = recognizedfaces[i].Rectangle.Min.Y
		x2 = recognizedfaces[i].Rectangle.Max.X
		y2 = recognizedfaces[i].Rectangle.Max.Y
		v = recognizedfaces[i].Descriptor[:]
		newFace.faceID = faceID
		newFace.fileID = fileID
		newFace.x1 = x1
		newFace.y1 = y1
		newFace.x2 = x2
		newFace.y2 = y2
		newFace.clusterID = ""
		newFace.personName = ""
		newFace.vector = v
		faces = append(faces, newFace)
	}
	//add faces to faces table
	for i := 0; i < len(faces); i++ {
		_, err := db.Exec("REPLACE INTO faces (faceID, fileID, x1, y1, x2, y2, clusterID, personName, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v120, v121, v122, v123, v124, v125, v126, v127) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", faces[i].faceID, faces[i].fileID, faces[i].x1, faces[i].y1, faces[i].x2, faces[i].y2, faces[i].clusterID, faces[i].personName, faces[i].vector[0], faces[i].vector[1], faces[i].vector[2], faces[i].vector[3], faces[i].vector[4], faces[i].vector[5], faces[i].vector[6], faces[i].vector[7], faces[i].vector[8], faces[i].vector[9], faces[i].vector[10], faces[i].vector[11], faces[i].vector[12], faces[i].vector[13], faces[i].vector[14], faces[i].vector[15], faces[i].vector[16], faces[i].vector[17], faces[i].vector[18], faces[i].vector[19], faces[i].vector[20], faces[i].vector[21], faces[i].vector[22], faces[i].vector[23], faces[i].vector[24], faces[i].vector[25], faces[i].vector[26], faces[i].vector[27], faces[i].vector[28], faces[i].vector[29], faces[i].vector[30], faces[i].vector[31], faces[i].vector[32], faces[i].vector[33], faces[i].vector[34], faces[i].vector[35], faces[i].vector[36], faces[i].vector[37], faces[i].vector[38], faces[i].vector[39], faces[i].vector[40], faces[i].vector[41], faces[i].vector[42], faces[i].vector[43], faces[i].vector[44], faces[i].vector[45], faces[i].vector[46], faces[i].vector[47], faces[i].vector[48], faces[i].vector[49], faces[i].vector[50], faces[i].vector[51], faces[i].vector[52], faces[i].vector[53], faces[i].vector[54], faces[i].vector[55], faces[i].vector[56], faces[i].vector[57], faces[i].vector[58], faces[i].vector[59], faces[i].vector[60], faces[i].vector[61], faces[i].vector[62], faces[i].vector[63], faces[i].vector[64], faces[i].vector[65], faces[i].vector[66], faces[i].vector[67], faces[i].vector[68], faces[i].vector[69], faces[i].vector[70], faces[i].vector[71], faces[i].vector[72], faces[i].vector[73], faces[i].vector[74], faces[i].vector[75], faces[i].vector[76], faces[i].vector[77], faces[i].vector[78], faces[i].vector[79], faces[i].vector[80], faces[i].vector[81], faces[i].vector[82], faces[i].vector[83], faces[i].vector[84], faces[i].vector[85], faces[i].vector[86], faces[i].vector[87], faces[i].vector[88], faces[i].vector[89], faces[i].vector[90], faces[i].vector[91], faces[i].vector[92], faces[i].vector[93], faces[i].vector[94], faces[i].vector[95], faces[i].vector[96], faces[i].vector[97], faces[i].vector[98], faces[i].vector[99], faces[i].vector[100], faces[i].vector[101], faces[i].vector[102], faces[i].vector[103], faces[i].vector[104], faces[i].vector[105], faces[i].vector[106], faces[i].vector[107], faces[i].vector[108], faces[i].vector[109], faces[i].vector[110], faces[i].vector[111], faces[i].vector[112], faces[i].vector[113], faces[i].vector[114], faces[i].vector[115], faces[i].vector[116], faces[i].vector[117], faces[i].vector[118], faces[i].vector[119], faces[i].vector[120], faces[i].vector[121], faces[i].vector[122], faces[i].vector[123], faces[i].vector[124], faces[i].vector[125], faces[i].vector[126], faces[i].vector[127])
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	//return
	// elapsed := float64(time.Since(start))/1000000000
	// fmt.Printf("%v   %v\n", elapsed, len(faces))
	//log.Printf("fileID %v; found %v faces\n", fileID, len(faces))
	return len(faces), nil
}
