package storage

import "northpole/room"

var _ RoomStorage = RoomStorageMock{}

type RoomStorageMock struct {
	RoomMock  room.Room
	ErrorMock error
}

func (ms RoomStorageMock) Find(m room.Room) (room.Room, error) {
	return ms.RoomMock, ms.ErrorMock
}

func (ms RoomStorageMock) FindFirst() (room.Room, error) {
	return ms.RoomMock, ms.ErrorMock
}

func (ms RoomStorageMock) Add(m room.Room) error {
	return ms.ErrorMock
}

func (ms RoomStorageMock) Remove(m room.Room) error {
	return ms.ErrorMock
}
