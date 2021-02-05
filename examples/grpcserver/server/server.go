package server

import (
	"context"

	"grpcserver/utils"

	"github.com/google/uuid"
	match "github.com/k-jun/northpole"
	"github.com/k-jun/northpole/room"
	"github.com/k-jun/northpole/storage"
	"github.com/k-jun/northpole/user"
	"google.golang.org/grpc"

	pb "grpcserver/grpc"
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

func roomToRoomInfo(r room.Room) *pb.RoomInfo {
	status := pb.RoomStatus_Close
	if r.IsOpen() {
		status = pb.RoomStatus_Open
	}
	return &pb.RoomInfo{
		Id:                   r.ID(),
		Status:               status,
		CurrentNumberOfUsers: int64(r.CurrentNumberOfUsers()),
		MaxNumberOfUsers:     int64(r.MaxNumberOfUsers()),
	}

}

func (s *northPoleServer) JoinPublicRoom(mi *pb.MatchInfo, stream pb.NorthPole_JoinPublicRoomServer) error {
	uid, err := uuid.Parse(mi.UserId)
	if err != nil {
		return err
	}
	u := user.New(uid.String())

	channel, err := s.publicMatch.JoinRandomRoom(u)
	if err != nil {
		if err == storage.RoomStorageRoomNotFound {
			nid := utils.NewUUID()
			r := room.New(nid.String(), deafultMNOU, func(_ string) error { return nil })
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
		if r == nil {
			break
		}
		stream.Send(roomToRoomInfo(r))
	}
	return nil
}

func (s *northPoleServer) CreatePrivateRoom(rci *pb.RoomCreateInfo, stream pb.NorthPole_CreatePrivateRoomServer) error {
	uid, err := uuid.Parse(rci.UserId)
	if err != nil {
		return err
	}
	rid, err := uuid.Parse(rci.RoomId)
	if err != nil {
		return err
	}
	u := user.New(uid.String())
	r := room.New(rid.String(), int(rci.MaxNumberOfUsers), func(_ string) error { return nil })

	channel, err := s.privateMatch.CreateRoom(u, r)
	if err != nil {
		return err
	}

	for {
		r := <-channel
		if r == nil {
			break
		}
		stream.Send(roomToRoomInfo(r))
	}
	return nil
}

func (s *northPoleServer) JoinPrivateRoom(mi *pb.MatchInfo, stream pb.NorthPole_JoinPrivateRoomServer) error {
	uid, err := uuid.Parse(mi.UserId)
	if err != nil {
		return err
	}
	rid, err := uuid.Parse(mi.RoomId)
	if err != nil {
		return err
	}
	u := user.New(uid.String())
	r := room.New(rid.String(), 0, func(_ string) error { return nil })

	channel, err := s.privateMatch.JoinRoom(u, r)
	if err != nil {
		return err
	}

	for {
		r := <-channel
		if r == nil {
			break
		}
		stream.Send(roomToRoomInfo(r))
	}
	return nil
}

func (s *northPoleServer) LeaveRoom(ctx context.Context, mi *pb.MatchInfo) (*pb.Empty, error) {
	uid, err := uuid.Parse(mi.UserId)
	if err != nil {
		return &pb.Empty{}, err
	}
	rid, err := uuid.Parse(mi.RoomId)
	if err != nil {
		return &pb.Empty{}, err
	}
	u := user.New(uid.String())
	r := room.New(rid.String(), 0, func(_ string) error { return nil })

	err = s.publicMatch.LeaveRoom(u, r)
	if err != nil {
		if err == storage.RoomStorageRoomNotFound {
			return &pb.Empty{}, s.privateMatch.LeaveRoom(u, r)
		} else {
			return &pb.Empty{}, err
		}
	}

	return &pb.Empty{}, nil
}
