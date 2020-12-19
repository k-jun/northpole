package server

import (
	"context"

	pb "northpole/grpc"

	"northpole/storage"

	"google.golang.org/grpc"
)

var port = 8080

type northPoleServer struct {
	privateMatchStroage storage.MatchStorage
	publicMatchStroage  storage.MatchStorage
}

func (s *northPoleServer) JoinPublicMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPublicMatchServer) error {
	// validation
	// userId, err := uuid.Parse(userInfo.Id)
	// if err != nil {
	// 	return err
	// }
	// // view
	// targetUser := user.New(userId)
	//
	// // usecase
	// targetMatch := s.privateMatchStroage.FindFirst()
	// if targetMatch == nil {
	// 	newUUID := uuid.New()
	// 	targetMatch = match.New(newUUID)
	// 	s.publicMatchStroage.Add(targetMatch)
	// }
	//
	// err = targetMatch.JoinUser(targetUser)
	// if err != nil {
	// 	return err
	// }

	// channel
	// for {
	// 	m := <-targetMatch.Channel()
	// 	// view
	// 	testMatchInfo := &pb.MatchInfo{
	// 		Id:                   m.ID().String(),
	// 		Status:               m.Status(),
	// 		CurrentNumberOfUsers: 3,
	// 		MaxNumberOfUsers:     4,
	// 	}
	// 	stream.Send(testMatchInfo)
	//
	// }
	return nil
}

func (s *northPoleServer) CreatePrivateMatch(userInfo *pb.UserInfo, stream pb.NorthPole_CreatePrivateMatchServer) error {
	for i := 0; i < 3; i++ {
		testMatchInfo := &pb.MatchInfo{
			Id:                   "c0310d40-a450-3af9-9e4e-3f4e0f4c26df",
			Status:               pb.MatchStatus_Availabel,
			CurrentNumberOfUsers: 3,
			MaxNumberOfUsers:     4,
		}
		stream.Send(testMatchInfo)
	}
	return nil
}

func (s *northPoleServer) JoinPrivateMatch(midAndUid *pb.MatchIDAndUserID, stream pb.NorthPole_JoinPrivateMatchServer) error {
	for i := 0; i < 3; i++ {
		testMatchInfo := &pb.MatchInfo{
			Id:                   "c0310d40-a450-3af9-9e4e-3f4e0f4c26df",
			Status:               pb.MatchStatus_Availabel,
			CurrentNumberOfUsers: 3,
			MaxNumberOfUsers:     4,
		}
		stream.Send(testMatchInfo)
	}
	return nil
}

func (s *northPoleServer) LeavePublicMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	testMatchInfo := &pb.MatchInfo{
		Id:                   "03240404-7d61-388f-8584-99a7e0438363",
		Status:               pb.MatchStatus_Availabel,
		CurrentNumberOfUsers: 3,
		MaxNumberOfUsers:     4,
	}
	return testMatchInfo, nil
}

func (s *northPoleServer) LeavePrivateMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	testMatchInfo := &pb.MatchInfo{
		Id:                   "03240404-7d61-388f-8584-99a7e0438363",
		Status:               pb.MatchStatus_Availabel,
		CurrentNumberOfUsers: 3,
		MaxNumberOfUsers:     4,
	}
	return testMatchInfo, nil
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	pubms := storage.NewMatchStorage()
	prims := storage.NewMatchStorage()
	pb.RegisterNorthPoleServer(grpcServer, &northPoleServer{publicMatchStroage: pubms, privateMatchStroage: prims})
	return grpcServer
}
