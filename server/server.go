package server

import (
	"context"

	pb "northpole/grpc"
	"northpole/match"
	"northpole/usecase"
	"northpole/user"

	"northpole/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var port = 8080

type northPoleServer struct {
	matchUsecase usecase.MatchUsecase
}

func (s *northPoleServer) JoinPublicMatch(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPublicMatchServer) error {
	userId, err := uuid.Parse(userInfo.Id)
	if err != nil {
		return err
	}
	u := user.New(userId)

	m, err := s.matchUsecase.JoinPublicMatch(u)
	if err != nil {
		return err
	}

	for {
		cm := <-m.Channel()
		if cm.ID() == uuid.Nil {
			break
		}
		stream.Send(cm.MatchInfo())
	}
	return nil
}

func (s *northPoleServer) CreatePrivateMatch(userInfo *pb.UserInfo, stream pb.NorthPole_CreatePrivateMatchServer) error {
	userId, err := uuid.Parse(userInfo.Id)
	if err != nil {
		return err
	}
	u := user.New(userId)

	m, err := s.matchUsecase.CreatePrivateMatch(u)
	if err != nil {
		return err
	}

	for {
		cm := <-m.Channel()
		if cm.ID() == uuid.Nil {
			break
		}
		stream.Send(cm.MatchInfo())
	}
	return nil
}

func (s *northPoleServer) JoinPrivateMatch(midAndUid *pb.MatchIDAndUserID, stream pb.NorthPole_JoinPrivateMatchServer) error {
	userId, err := uuid.Parse(midAndUid.UserId)
	if err != nil {
		return err
	}
	u := user.New(userId)
	matchId, err := uuid.Parse(midAndUid.MatchId)
	if err != nil {
		return err
	}
	m := match.New(matchId)

	m, err = s.matchUsecase.JoinPrivateMatch(u, m)
	if err != nil {
		return err
	}

	for {
		cm := <-m.Channel()
		if cm.ID() == uuid.Nil {
			break
		}
		stream.Send(cm.MatchInfo())
	}
	return nil
}

func (s *northPoleServer) LeavePublicMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	userId, err := uuid.Parse(midAndUid.UserId)
	if err != nil {
		return nil, err
	}
	u := user.New(userId)
	matchId, err := uuid.Parse(midAndUid.MatchId)
	if err != nil {
		return nil, err
	}
	m := match.New(matchId)

	m, err = s.matchUsecase.LeavePublicMatch(u, m)
	if err != nil {
		return nil, err
	}

	return m.MatchInfo(), nil
}

func (s *northPoleServer) LeavePrivateMatch(ctx context.Context, midAndUid *pb.MatchIDAndUserID) (*pb.MatchInfo, error) {
	userId, err := uuid.Parse(midAndUid.UserId)
	if err != nil {
		return nil, err
	}
	u := user.New(userId)
	matchId, err := uuid.Parse(midAndUid.MatchId)
	if err != nil {
		return nil, err
	}
	m := match.New(matchId)

	m, err = s.matchUsecase.LeavePrivateMatch(u, m)
	if err != nil {
		return nil, err
	}

	return m.MatchInfo(), nil
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	pubms := storage.NewMatchStorage()
	prims := storage.NewMatchStorage()
	matchUsecase := usecase.NewMatchUsecase(pubms, prims)
	pb.RegisterNorthPoleServer(grpcServer, &northPoleServer{matchUsecase: matchUsecase})
	return grpcServer
}
