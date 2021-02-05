package storage

import (
	"testing"

	"github.com/k-jun/northpole/room"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	uuid1 := "b2ecdd9a-850e-3ad4-896d-400ad7f1bebf"
	uuid2 := "23d25efb-28cd-3593-8fa6-8ed061b5e381"
	openRoom := room.RoomMock{IDMock: uuid1}
	closeRoom := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[string]room.Room
		inRoom      room.Room
		outRoom     room.Room
		outError    error
	}{
		{
			beforeRooms: map[string]room.Room{
				uuid1: openRoom,
			},
			inRoom:   openRoom,
			outRoom:  openRoom,
			outError: nil,
		},
		{
			beforeRooms: map[string]room.Room{
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
	uuid1 := "f533413e-479d-302c-86af-bad6f0adb4c1"
	uuid2 := "6a5f95a8-e41c-37ad-872d-a2fdd1729b3b"
	closeRoom := room.RoomMock{StatusMock: room.Close}
	openRoom := room.RoomMock{StatusMock: room.Open}
	cases := []struct {
		beforeRooms map[string]room.Room
		outRoom     room.Room
		outError    error
	}{
		{
			beforeRooms: map[string]room.Room{
				uuid1: closeRoom,
				uuid2: openRoom,
			},
			outRoom:  openRoom,
			outError: nil,
		},
		{
			beforeRooms: map[string]room.Room{
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
	uuid1 := "657982e4-62db-35fa-ba4c-b6f341316fd9"
	uuid2 := "9d21aee7-a5d8-383b-988d-d22c8430b40c"
	room1 := room.RoomMock{IDMock: uuid1}
	room2 := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[string]room.Room
		inRoom      room.Room
		afterRooms  map[string]room.Room
		outError    error
	}{
		{
			beforeRooms: map[string]room.Room{
				uuid1: room1,
			},
			inRoom: room2,
			afterRooms: map[string]room.Room{
				uuid1: room1,
				uuid2: room2,
			},
			outError: nil,
		},
		{
			beforeRooms: map[string]room.Room{
				uuid2: room2,
			},
			inRoom:     room2,
			afterRooms: map[string]room.Room{},
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
	uuid1 := "b096e181-c755-35e6-933b-26f5c952a095"
	uuid2 := "0fc9e996-b4e2-3009-87c9-3705c666a4b7"
	room1 := room.RoomMock{IDMock: uuid1}
	room2 := room.RoomMock{IDMock: uuid2}
	cases := []struct {
		beforeRooms map[string]room.Room
		inRoom      room.Room
		afterRooms  map[string]room.Room
		outError    error
	}{
		{
			beforeRooms: map[string]room.Room{
				uuid1: room1,
			},
			inRoom:     room1,
			afterRooms: map[string]room.Room{},
			outError:   nil,
		},
		{
			beforeRooms: map[string]room.Room{
				uuid1: room1,
			},
			inRoom:     room2,
			afterRooms: map[string]room.Room{},
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
