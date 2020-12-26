package northpole

import (
	"github.com/k-jun/northpole/room"
	"github.com/k-jun/northpole/storage"
	"github.com/k-jun/northpole/user"
)

type Match interface {
	CreateRoom(user.User, room.Room) (chan room.Room, error)
	JoinRoom(user.User, room.Room) (chan room.Room, error)
	JoinRandomRoom(user.User) (chan room.Room, error)
	LeaveRoom(user.User, room.Room) error
}

type MatchImpl struct {
	roomStorage storage.RoomStorage
}

func New() Match {
	rs := storage.NewRoomStorage()
	return &MatchImpl{roomStorage: rs}
}

func (np *MatchImpl) CreateRoom(u user.User, r room.Room) (chan room.Room, error) {
	if err := np.roomStorage.Add(r); err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *MatchImpl) JoinRoom(u user.User, r room.Room) (chan room.Room, error) {
	r, err := np.roomStorage.Find(r)
	if err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *MatchImpl) JoinRandomRoom(u user.User) (chan room.Room, error) {
	r, err := np.roomStorage.FindFirst()
	if err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *MatchImpl) LeaveRoom(u user.User, r room.Room) error {
	r, err := np.roomStorage.Find(r)
	if err != nil {
		return err
	}
	return r.LeaveUser(u)
}
