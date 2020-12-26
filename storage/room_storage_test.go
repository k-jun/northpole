package storage

import (
	"northpole/room"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	openRoom := room.RoomMock{IDMock: uuid1}
	closeRoom := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[uuid.UUID]room.Room
		inRoom      room.Room
		outRoom     room.Room
		outError    error
	}{
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: openRoom,
			},
			inRoom:   openRoom,
			outRoom:  openRoom,
			outError: nil,
		},
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: openRoom,
			},
			inRoom:   closeRoom,
			outRoom:  nil,
			outError: RoomStorageRoomNotFound,
		},
	}

	for _, c := range cases {
		ms := RoomStorageImpl{rooms: c.beforeRooms}
		m, err := ms.Find(c.inRoom)
		assert.Equal(t, c.outRoom, m)
		assert.Equal(t, c.outError, err)
	}
}

func TestFindFirst(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	closeRoom := room.RoomMock{StatusMock: room.Close}
	openRoom := room.RoomMock{StatusMock: room.Open}
	cases := []struct {
		beforeRooms map[uuid.UUID]room.Room
		outRoom     room.Room
		outError    error
	}{
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: closeRoom,
				uuid2: openRoom,
			},
			outRoom:  openRoom,
			outError: nil,
		},
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: closeRoom,
				uuid2: closeRoom,
			},
			outRoom:  nil,
			outError: RoomStorageRoomNotFound,
		},
	}

	for _, c := range cases {
		ms := RoomStorageImpl{rooms: c.beforeRooms}
		Room, err := ms.FindFirst()
		assert.Equal(t, c.outRoom, Room)
		assert.Equal(t, c.outError, err)
	}
}

func TestAdd(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	room1 := room.RoomMock{IDMock: uuid1}
	room2 := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[uuid.UUID]room.Room
		inRoom      room.Room
		afterRooms  map[uuid.UUID]room.Room
		outError    error
	}{
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: room1,
			},
			inRoom: room2,
			afterRooms: map[uuid.UUID]room.Room{
				uuid1: room1,
				uuid2: room2,
			},
			outError: nil,
		},
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid2: room2,
			},
			inRoom:     room2,
			afterRooms: map[uuid.UUID]room.Room{},
			outError:   RoomStorageRoomAlreadyExistErr,
		},
	}

	for _, c := range cases {
		ms := RoomStorageImpl{rooms: c.beforeRooms}
		err := ms.Add(c.inRoom)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterRooms, ms.rooms)
	}
}

func TestRemove(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	room1 := room.RoomMock{IDMock: uuid1}
	room2 := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[uuid.UUID]room.Room
		inRoom      room.Room
		afterRooms  map[uuid.UUID]room.Room
		outError    error
	}{
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: room1,
			},
			inRoom:     room1,
			afterRooms: map[uuid.UUID]room.Room{},
			outError:   nil,
		},
		{
			beforeRooms: map[uuid.UUID]room.Room{
				uuid1: room1,
			},
			inRoom:     room2,
			afterRooms: map[uuid.UUID]room.Room{},
			outError:   RoomStorageRoomNotFound,
		},
	}

	for _, c := range cases {
		ms := RoomStorageImpl{rooms: c.beforeRooms}
		err := ms.Remove(c.inRoom)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterRooms, ms.rooms)
	}
}
