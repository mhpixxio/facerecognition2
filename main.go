package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

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
type serverMergeClusters struct {
	pbface2.MergeClustersServiceServer
}
type serverMoveFacesToAnotherCluster struct {
	pbface2.MoveFacesToAnotherClusterServiceServer
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
	pbface2.RegisterMergeClustersServiceServer(s, &serverMergeClusters{})
	pbface2.RegisterMoveFacesToAnotherClusterServiceServer(s, &serverMoveFacesToAnotherCluster{})
	pbface2.RegisterRemoveFacesFromDatabaseServiceServer(s, &serverRemoveFacesFromDatabase{})
	pbface2.RegisterRemoveFilesFromDatabaseServiceServer(s, &serverRemoveFilesFromDatabase{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	//---------------------------------- set the sql connection for the databases ----------------------------------
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
	}

	//---------------------------------- loop for updating the sql tables ----------------------------------
	for {
		time.Sleep(time.Minute)
		err = UpdateFacesAndClusters(db)
		if err != nil {
			log.Println(err)
		}
	}
}

//---------------------------------- the grpc functions that call their corresponding functions ----------------------------------

func (s *serverReclustering) ReclusteringFunc(ctx context.Context, request *pbface2.EmptyMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = Reclustering(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverUpdateFacesAndClusters) UpdateFacesAndClustersFunc(ctx context.Context, request *pbface2.EmptyMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = UpdateFacesAndClusters(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverRenameCluster) RenameClusterFunc(ctx context.Context, request *pbface2.RenameClusterMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = RenameCluster(db, request.ClusterID, request.NewPersonName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverMergeClusters) MergeClusterFunc(ctx context.Context, request *pbface2.MergeClustersMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var clusterIDs []string
	for i := 0; i < len(request.ClusterIDs); i++ {
		clusterIDs[i] = request.ClusterIDs[i].ClusterID
	}
	err = MergeClusters(db, clusterIDs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverMoveFacesToAnotherCluster) MoveFacesToAnotherClusterFunc(ctx context.Context, request *pbface2.MoveFacesToAnotherClusterMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var faceIDs []int
	for i := 0; i < len(request.FaceIDs); i++ {
		faceIDs[i] = int(request.FaceIDs[i].FaceID)
	}
	err = MoveFacesToAnotherCluster(db, faceIDs, request.ClusterID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverRemoveFacesFromDatabase) RemoveFacesFromDatabaseFunc(ctx context.Context, request *pbface2.RemoveFacesFromDatabaseMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var faceIDs []int
	for i := 0; i < len(request.FaceIDs); i++ {
		faceIDs[i] = int(request.FaceIDs[i].FaceID)
	}
	err = RemoveFacesFromDatabase(db, faceIDs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *serverRemoveFilesFromDatabase) RemoveFilesFromDatabaseFunc(ctx context.Context, request *pbface2.RemoveFilesFromDatabaseMessage) (*pbface2.ErrMessage, error) {
	db, err := ConnectToSQLDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var fileIDs []int
	for i := 0; i < len(request.FileIDs); i++ {
		fileIDs[i] = int(request.FileIDs[i].FileID)
	}
	err = RemoveFilesFromDatabase(db, fileIDs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}
