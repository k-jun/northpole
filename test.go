package northpole

import (
	"errors"
	"testing"

	"github.com/k-jun/northpole/room"
	"github.com/k-jun/northpole/storage"
	"github.com/k-jun/northpole/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser            user.User
		inRoom            room.Room
		beforeRoomStorage storage.RoomStorage
		outError          error
	}{
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{},
			outError:          nil,
		},
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{ErrorMock: errors.New("")},
			outError:          errors.New(""),
		},
	}

	for _, c := range cases {
		m := MatchImpl{roomStorage: c.beforeRoomStorage}

		_, err := m.CreateRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}

func TestJoinRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser            user.User
		inRoom            room.Room
		beforeRoomStorage storage.RoomStorage
		outError          error
	}{
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}},
			outError:          nil,
		},
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}, ErrorMock: errors.New("")},
			outError:          errors.New(""),
		},
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{ErrorMock: errors.New("")}},
			outError:          errors.New(""),
		},
	}

	for _, c := range cases {

		m := MatchImpl{roomStorage: c.beforeRoomStorage}
		_, err := m.JoinRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}

func TestJoinRandomRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser            user.User
		beforeRoomStorage storage.RoomStorage
		outError          error
	}{
		{
			inUser:            user.New(uuid1),
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}},
			outError:          nil,
		},
		{
			inUser:            user.New(uuid1),
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}, ErrorMock: errors.New("")},
			outError:          errors.New(""),
		},
		{
			inUser:            user.New(uuid1),
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{ErrorMock: errors.New("")}},
			outError:          errors.New(""),
		},
	}

	for _, c := range cases {

		m := MatchImpl{roomStorage: c.beforeRoomStorage}
		_, err := m.JoinRandomRoom(c.inUser)
		assert.Equal(t, c.outError, err)
	}
}

func TestLeaveRoom(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser            user.User
		inRoom            room.Room
		beforeRoomStorage storage.RoomStorage
		outError          error
	}{
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}},
			outError:          nil,
		},
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{}, ErrorMock: errors.New("")},
			outError:          errors.New(""),
		},
		{
			inUser:            user.New(uuid1),
			inRoom:            room.RoomMock{},
			beforeRoomStorage: storage.RoomStorageMock{RoomMock: room.RoomMock{ErrorMock: errors.New("")}},
			outError:          errors.New(""),
		},
	}

	for _, c := range cases {

		m := MatchImpl{roomStorage: c.beforeRoomStorage}
		err := m.LeaveRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}
