package server

import (
	match "northpole"
	"northpole/room"
	"northpole/storage"
	"northpole/user"
	"northpole/utils"

	pb "northpole/example/grpcserver/grpc"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	deafultMNOU = 4
)

type northPoleServer struct {
	privateMatch match.Match
	publicMatch  match.Match
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	prim := match.New()
	pubm := match.New()

	pb.RegisterNorthPoleServer(grpcServer, &northPoleServer{privateMatch: prim, publicMatch: pubm})
	return grpcServer
}

func roomToRoomInfo(r room.Room) pb.RoomInfo {

}

func (s *northPoleServer) JoinPublicRoom(userInfo *pb.UserInfo, stream pb.NorthPole_JoinPublicRoomServer) error {
	userId, err := uuid.Parse(userInfo.Id)
	if err != nil {
		return err
	}
	u := user.New(userId)

	channel, err := s.publicMatch.JoinRandomRoom(u)
	if err != nil {
		if err == storage.RoomStorageRoomNotFound {
			nid := utils.NewUUID()
			r := room.New(nid, deafultMNOU)
			channel, err = s.publicMatch.CreateRoom(u, r)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for {
		r := <-channel
		if r.ID() == uuid.Nil {
			break
		}
		stream.Send(nil)
	}
	return nil
}

// func (s *northPoleServer) CreatePrivateRoom(userInfo *pb.UserInfo, stream pb.NorthPole_CreatePrivateRoomServer) error {
// 	userId, err := uuid.Parse(userInfo.Id)
// 	if err != nil {
// 		return err
// 	}
// 	u := user.New(userId)
//
// 	m, err := s.RoomUsecase.CreatePrivateRoom(u)
// 	if err != nil {
// 		return err
// 	}
//
// 	for {
// 		cm := <-m.Channel()
// 		if cm.ID() == uuid.Nil {
// 			break
// 		}
// 		stream.Send(cm.RoomInfo())
// 	}
// 	return nil
// }
//
// func (s *northPoleServer) JoinPrivateRoom(midAndUid *pb.RoomIDAndUserID, stream pb.NorthPole_JoinPrivateRoomServer) error {
// 	userId, err := uuid.Parse(midAndUid.UserId)
// 	if err != nil {
// 		return err
// 	}
// 	u := user.New(userId)
// 	RoomId, err := uuid.Parse(midAndUid.RoomId)
// 	if err != nil {
// 		return err
// 	}
// 	m := Room.New(RoomId)
//
// 	m, err = s.RoomUsecase.JoinPrivateRoom(u, m)
// 	if err != nil {
// 		return err
// 	}
//
// 	for {
// 		cm := <-m.Channel()
// 		if cm.ID() == uuid.Nil {
// 			break
// 		}
// 		stream.Send(cm.RoomInfo())
// 	}
// 	return nil
// }
//
// func (s *northPoleServer) LeavePublicRoom(ctx context.Context, midAndUid *pb.RoomIDAndUserID) (*pb.RoomInfo, error) {
// 	userId, err := uuid.Parse(midAndUid.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := user.New(userId)
// 	RoomId, err := uuid.Parse(midAndUid.RoomId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	m := Room.New(RoomId)
//
// 	m, err = s.RoomUsecase.LeavePublicRoom(u, m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return m.RoomInfo(), nil
// }
//
// func (s *northPoleServer) LeavePrivateRoom(ctx context.Context, midAndUid *pb.RoomIDAndUserID) (*pb.RoomInfo, error) {
// 	userId, err := uuid.Parse(midAndUid.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := user.New(userId)
// 	RoomId, err := uuid.Parse(midAndUid.RoomId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	m := Room.New(RoomId)
//
// 	m, err = s.RoomUsecase.LeavePrivateRoom(u, m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return m.RoomInfo(), nil
// }
//
