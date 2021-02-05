package northpole

import (
	"errors"
	"testing"

	"github.com/k-jun/northpole/room"
	"github.com/k-jun/northpole/storage"
	"github.com/k-jun/northpole/user"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	uuid1 := "c7fba570-12a2-343e-98de-028e91c410da"
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
		m := matchImpl{roomStorage: c.beforeRoomStorage}

		_, err := m.CreateRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}

func TestJoinRoom(t *testing.T) {
	uuid1 := "439f9aa5-b7aa-33c7-b423-7d6f9843d158"
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

		m := matchImpl{roomStorage: c.beforeRoomStorage}
		_, err := m.JoinRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}

func TestJoinRandomRoom(t *testing.T) {
	uuid1 := "b37107d5-58e3-30f7-a66a-3f3acb185cb0"
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

		m := matchImpl{roomStorage: c.beforeRoomStorage}
		_, err := m.JoinRandomRoom(c.inUser)
		assert.Equal(t, c.outError, err)
	}
}

func TestLeaveRoom(t *testing.T) {
	uuid1 := "27d6f505-7fb7-36ff-8a43-587864cf9b42"
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

		m := matchImpl{roomStorage: c.beforeRoomStorage}
		err := m.LeaveRoom(c.inUser, c.inRoom)
		assert.Equal(t, c.outError, err)
	}
}
