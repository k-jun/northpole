package room

import (
	"github.com/k-jun/northpole/user"

	"github.com/google/uuid"
)

var _ Room = RoomMock{}

type RoomMock struct {
	ErrorMock                error
	StatusMock               RoomStatus
	IDMock                   uuid.UUID
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

func (m RoomMock) ID() uuid.UUID {
	return m.IDMock
}
