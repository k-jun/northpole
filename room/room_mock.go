package room

import (
	"northpole/user"

	"github.com/google/uuid"
)

var _ Room = RoomMock{}

type RoomMock struct {
	ErrorMock   error
	StatusMock  RoomStatus
	IDMock      uuid.UUID
	ChannelMock chan Room
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
