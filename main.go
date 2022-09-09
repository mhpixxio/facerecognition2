package main

//----------- instructions -----------
//build the SQL tables. In the MySQL 8.0 Command Line:
//$ use facerecognition
//$ source <absolute path to sql file create-tables.sql found in the directory SQLFiles>
//build the Docker Image (Docker Dekstop must be running)
//$ docker build -t facerecognition2 .
//run the docker image as container
//$ docker run -d -p 8080:8080 -v C:/Users/MichaelHuber/Desktop/EnvironmentForFaceRecognition/files:/app/files/ --memory 1000m --restart unless-stopped facerecognition2
//use the client foun at github.com/mhpixxio/clientfacerecognition2 to represent the Front End

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbface2 "github.com/mhpixxio/pbface2"
)

type fileStruct struct {
	fileID     int
	pathToFile string
	processed  bool
	forRemoval bool
	removed    bool
}

type faceStruct struct {
	faceID     int
	fileID     int
	x1         int
	y1         int
	x2         int
	y2         int
	clusterID  string
	personName string
	vector     []float32
}

type clusterStruct struct {
	clusterID   string
	personName  string
	numberFaces int
	meanVector  []float32
}

type serverReclustering struct {
	pbface2.ReclusteringServiceServer
}
type serverUpdateFacesAndClusters struct {
	pbface2.UpdateFacesAndClustersServiceServer
}
type serverRenameCluster struct {
	pbface2.RenameClusterServiceServer
}
type serverDeleteAllPersonNames struct {
	pbface2.DeleteAllPersonNamesServiceServer
}
type serverMergeClusters struct {
	pbface2.MergeClustersServiceServer
}
type serverManuallyMoveFacesToAnotherCluster struct {
	pbface2.ManuallyMoveFacesToAnotherClusterServiceServer
}
type serverRemoveFacesFromDatabase struct {
	pbface2.RemoveFacesFromDatabaseServiceServer
}
type serverRemoveFilesFromDatabase struct {
	pbface2.RemoveFilesFromDatabaseServiceServer
}

func main() {
	//---------------------------------- set the grpc connection ----------------------------------
	//flags
	portAddressFlag := flag.String("portAddress", ":8080", "the portAddress")
	flag.Parse()
	portAddress := *portAddressFlag
	log.Printf("portAddress: %v", portAddress)
	//start server
	log.Printf("starting server at port" + portAddress + "\n")
	listener, err := net.Listen("tcp", portAddress)
	if err != nil {
		log.Printf("unable to listen: %v", err)
	}
	//define calloptions
	maxSize := 1 //in megabytes
	calloptionRecv := grpc.MaxRecvMsgSize(maxSize * 8000000)
	calloptionSend := grpc.MaxSendMsgSize(maxSize * 8000000)
	//start services
	s := grpc.NewServer(calloptionRecv, calloptionSend)
	pbface2.RegisterReclusteringServiceServer(s, &serverReclustering{})
	pbface2.RegisterUpdateFacesAndClustersServiceServer(s, &serverUpdateFacesAndClusters{})
	pbface2.RegisterRenameClusterServiceServer(s, &serverRenameCluster{})
	pbface2.RegisterDeleteAllPersonNamesServiceServer(s, &serverDeleteAllPersonNames{})
	pbface2.RegisterMergeClustersServiceServer(s, &serverMergeClusters{})
	pbface2.RegisterManuallyMoveFacesToAnotherClusterServiceServer(s, &serverManuallyMoveFacesToAnotherCluster{})
	pbface2.RegisterRemoveFacesFromDatabaseServiceServer(s, &serverRemoveFacesFromDatabase{})
	pbface2.RegisterRemoveFilesFromDatabaseServiceServer(s, &serverRemoveFilesFromDatabase{})
	reflection.Register(s)

	//---------------------------------- set the sql connection for the databases ----------------------------------
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
	}

	//---------------------------------- updating the sql tables ----------------------------------
	numberOfUnprocessedFiles := 1
	for numberOfUnprocessedFiles > 0 {
		err = UpdateFacesAndClusters(db)
		if err != nil {
			log.Println(err)
		}
		err = db.QueryRow("SELECT COUNT(*) FROM files WHERE processed = ?", false).Scan(&numberOfUnprocessedFiles)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%v number of files left for processing", numberOfUnprocessedFiles)
	}

	//---------------------------------- start listening to clients ----------------------------------
	log.Println("ready for service")
	if err := s.Serve(listener); err != nil {
		log.Printf("failed to serve: %v", err)
	}

}

//---------------------------------- the grpc functions that call their corresponding functions ----------------------------------

func (s *serverReclustering) ReclusteringFunc(ctx context.Context, request *pbface2.EmptyMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	err = Reclustering(db)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverUpdateFacesAndClusters) UpdateFacesAndClustersFunc(ctx context.Context, request *pbface2.EmptyMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	err = UpdateFacesAndClusters(db)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverRenameCluster) RenameClusterFunc(ctx context.Context, request *pbface2.RenameClusterMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	err = RenameCluster(db, request.ClusterID, request.NewPersonName)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverDeleteAllPersonNames) DeleteAllPersonNamesFunc(ctx context.Context, request *pbface2.EmptyMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	err = DeleteAllPersonNames(db)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverMergeClusters) MergeClustersFunc(ctx context.Context, request *pbface2.MergeClustersMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	var clusterIDs []string
	for i := 0; i < len(request.ClusterIDs); i++ {
		clusterIDs[i] = request.ClusterIDs[i].ClusterID
	}
	err = MergeClusters(db, clusterIDs)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverManuallyMoveFacesToAnotherCluster) ManuallyMoveFacesToAnotherClusterFunc(ctx context.Context, request *pbface2.ManuallyMoveFacesToAnotherClusterMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	var faceIDs []int
	for i := 0; i < len(request.FaceIDs); i++ {
		faceIDs[i] = int(request.FaceIDs[i].FaceID)
	}
	err = ManuallyMoveFacesToAnotherCluster(db, faceIDs, request.ClusterID)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverRemoveFacesFromDatabase) RemoveFacesFromDatabaseFunc(ctx context.Context, request *pbface2.RemoveFacesFromDatabaseMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	var faceIDs []int
	for i := 0; i < len(request.FaceIDs); i++ {
		faceIDs[i] = int(request.FaceIDs[i].FaceID)
	}
	err = RemoveFacesFromDatabase(db, faceIDs)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}

func (s *serverRemoveFilesFromDatabase) RemoveFilesFromDatabaseFunc(ctx context.Context, request *pbface2.RemoveFilesFromDatabaseMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	var fileIDs []int
	for i := 0; i < len(request.FileIDs); i++ {
		fileIDs[i] = int(request.FileIDs[i].FileID)
	}
	err = RemoveFilesFromDatabase(db, fileIDs)
	if err != nil {
		log.Println(err)
		return &pbface2.ErrMessage{ErrString: ""}, err
	}
	return &pbface2.ErrMessage{ErrString: ""}, nil
}
