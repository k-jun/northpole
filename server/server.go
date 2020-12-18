package server

import (
	"context"

	pb "northpole/grpc"

	"google.golang.org/grpc"
)

var port = 8080

type northPoleServer struct{}

func (s *northPoleServer) JoinPublicMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPublicMatchServer) error {
	return nil
}

func (s *northPoleServer) JoinPrivateMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPrivateMatchServer) error {
	return nil
}

func (s *northPoleServer) LeaveMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	return nil, nil
}

func NewServer() *grpc.Server {

	grpcServer := grpc.NewServer()
	pb.RegisterNorthPoleServer(grpcServer, &northPoleServer{})
	return grpcServer
}
