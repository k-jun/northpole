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

type matchImpl struct {
	roomStorage storage.RoomStorage
}

func New() Match {
	rs := storage.NewRoomStorage()
	return &matchImpl{roomStorage: rs}
}

func (np *matchImpl) CreateRoom(u user.User, r room.Room) (chan room.Room, error) {
	if err := np.roomStorage.Add(r); err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *matchImpl) JoinRoom(u user.User, r room.Room) (chan room.Room, error) {
	r, err := np.roomStorage.Find(r)
	if err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *matchImpl) JoinRandomRoom(u user.User) (chan room.Room, error) {
	r, err := np.roomStorage.FindFirst()
	if err != nil {
		return nil, err
	}
	return r.JoinUser(u)
}

func (np *matchImpl) LeaveRoom(u user.User, r room.Room) error {
	r, err := np.roomStorage.Find(r)
	if err != nil {
		return err
	}
	err = r.LeaveUser(u)
	if err != nil {
		return err
	}

	if !r.IsOpen() {
		return np.roomStorage.Remove(r)
	}
	return nil
}
