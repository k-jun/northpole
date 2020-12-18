package server

import (
	"context"

	pb "northpole/grpc"

	"google.golang.org/grpc"
)

var port = 8080

type northPoleServer struct{}

func (s *northPoleServer) JoinPublicMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPublicMatchServer) error {
	for i := 0; i < 3; i++ {
		testMatchInfo := &pb.MatchInfo{
			Id:                    "c51a0ef2-550b-32e3-90e1-930f52691b0e",
			Status:                pb.MatchStatus_Waiting,
			CurrentNumberOfPlayer: 3,
			MaxNumberOfPlayer:     4,
		}
		stream.Send(testMatchInfo)
	}
	return nil
}

func (s *northPoleServer) JoinPrivateMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPrivateMatchServer) error {
	for i := 0; i < 3; i++ {
		testMatchInfo := &pb.MatchInfo{
			Id:                    "c0310d40-a450-3af9-9e4e-3f4e0f4c26df",
			Status:                pb.MatchStatus_Waiting,
			CurrentNumberOfPlayer: 3,
			MaxNumberOfPlayer:     4,
		}
		stream.Send(testMatchInfo)
	}
	return nil
}

func (s *northPoleServer) LeaveMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	testMatchInfo := &pb.MatchInfo{
		Id:                    "03240404-7d61-388f-8584-99a7e0438363",
		Status:                pb.MatchStatus_Waiting,
		CurrentNumberOfPlayer: 3,
		MaxNumberOfPlayer:     4,
	}
	return testMatchInfo, nil
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterNorthPoleServer(grpcServer, &northPoleServer{})
	return grpcServer
}
