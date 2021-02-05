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

func (m RoomMock) IsOpen() bool {
	return m.StatusMock == Open
}
func (m RoomMock) CurrentNumberOfUsers() int {
	return m.CurrentNumberOfUsersMock
}
func (m RoomMock) MaxNumberOfUsers() int {
	return m.MaxNumberOfUsersMock
}

func (m RoomMock) JoinUser(u user.User) (chan Room, error) {
	return m.ChannelMock, m.ErrorMock
}

func (m RoomMock) LeaveUser(u user.User) error {
	return m.ErrorMock
}

func (m RoomMock) ID() string {
	return m.IDMock
}
