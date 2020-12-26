package main

import (
	"errors"
	"northpole/room"
	"northpole/storage"
	"northpole/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJoinPublicRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                 user.User
		beforeRooms            room.Room
		beforeRoomStorageError error
		beforeRoomError        error
		outError               error
	}{
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        nil,
			outError:               nil,
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: errors.New(""),
			beforeRoomError:        nil,
			outError:               errors.New(""),
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        errors.New(""),
			outError:               errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := room.RoomMock{IDMock: uuid2, ErrorMock: c.beforeRoomError}
		ms := storage.RoomStorageMock{RoomMock: m, ErrorMock: c.beforeRoomStorageError}
		u := RoomUsecaseImpl{publicRoomStorage: ms}

		Room, err := u.JoinPublicRoom(c.inUser)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, Room.ID())
	}
}

func TestJoinPrivateRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                 user.User
		inRoom                 room.Room
		beforeRoomStorageError error
		beforeRoomError        error
		outError               error
	}{
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        nil,
			outError:               nil,
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: errors.New(""),
			beforeRoomError:        nil,
			outError:               errors.New(""),
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        errors.New(""),
			outError:               errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := room.RoomMock{IDMock: uuid2, ErrorMock: c.beforeRoomError}
		ms := storage.RoomStorageMock{RoomMock: m, ErrorMock: c.beforeRoomStorageError}
		u := RoomUsecaseImpl{privateRoomStorage: ms}

		Room, err := u.JoinPrivateRoom(c.inUser, c.inRoom)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, Room.ID())
	}
}

func TestCreatePrivateRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                 user.User
		beforeRoomStorageError error
		outError               error
	}{
		{
			inUser:                 user.New(uuid1),
			beforeRoomStorageError: nil,
			outError:               nil,
		},
		{
			inUser:                 user.New(uuid1),
			beforeRoomStorageError: errors.New(""),
			outError:               errors.New(""),
		},
	}

	for _, c := range cases {
		ms := storage.RoomStorageMock{ErrorMock: c.beforeRoomStorageError}
		u := RoomUsecaseImpl{privateRoomStorage: ms}

		Room, err := u.CreatePrivateRoom(c.inUser)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.NotEqual(t, uuid.Nil, Room.ID())
	}
}

func TestLeavePublicRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                 user.User
		inRoom                 room.Room
		beforeRoomStorageError error
		beforeRoomError        error
		outError               error
	}{
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        nil,
			outError:               nil,
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: errors.New(""),
			beforeRoomError:        nil,
			outError:               errors.New(""),
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        errors.New(""),
			outError:               errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := room.RoomMock{IDMock: uuid2, ErrorMock: c.beforeRoomError}
		ms := storage.RoomStorageMock{RoomMock: m, ErrorMock: c.beforeRoomStorageError}
		u := RoomUsecaseImpl{publicRoomStorage: ms}

		Room, err := u.LeavePublicRoom(c.inUser, c.inRoom)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, Room.ID())
	}
}

func TestLeavePrivateRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                 user.User
		inRoom                 room.Room
		beforeRoomStorageError error
		beforeRoomError        error
		outError               error
	}{
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        nil,
			outError:               nil,
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: errors.New(""),
			beforeRoomError:        nil,
			outError:               errors.New(""),
		},
		{
			inUser:                 user.New(uuid1),
			inRoom:                 room.RoomMock{},
			beforeRoomStorageError: nil,
			beforeRoomError:        errors.New(""),
			outError:               errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := room.RoomMock{IDMock: uuid2, ErrorMock: c.beforeRoomError}
		ms := storage.RoomStorageMock{RoomMock: m, ErrorMock: c.beforeRoomStorageError}
		u := RoomUsecaseImpl{privateRoomStorage: ms}

		Room, err := u.LeavePrivateRoom(c.inUser, c.inRoom)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, Room.ID())
	}
}
