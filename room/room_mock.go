package room

import (
	"github.com/k-jun/northpole/user"
)

var _ Room = RoomMock{}

type RoomMock struct {
	ErrorMock                error
	StatusMock               RoomStatus
	IDMock                   string
	ChannelMock              chan Room
	CurrentNumberOfUsersMock int
	MaxNumberOfUsersMock     int
}

func (r RoomMock) IsOpen() bool {
	return r.StatusMock == Open
}
func (r RoomMock) CurrentNumberOfUsers() int {
	return r.CurrentNumberOfUsersMock
}
func (r RoomMock) MaxNumberOfUsers() int {
	return r.MaxNumberOfUsersMock
}

func (r RoomMock) JoinUser(u user.User) (chan Room, error) {
	return r.ChannelMock, r.ErrorMock
}

func (r RoomMock) LeaveUser(u user.User) error {
	return r.ErrorMock
}

func (r RoomMock) CloseRoom() error {
	return r.ErrorMock

}

func (r RoomMock) ID() string {
	return r.IDMock
}
