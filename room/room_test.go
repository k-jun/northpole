package room

import (
	"errors"
	"testing"

	"github.com/k-jun/northpole/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testUuid := uuid.New()
	testMNOU := 4
	testCallback := func(id uuid.UUID) error { return nil }
	room := New(testUuid, testMNOU, testCallback)
	assert.Equal(t, testUuid, room.ID())
}

func TestJoinUser(t *testing.T) {
	testUser := &user.UserMock{IdMock: uuid.New()}
	testRoomUser := &roomUser{u: testUser, c: nil}
	cases := []struct {
		beforeUsers            []*roomUser
		beforeStatus           RoomStatus
		beforeMaxNumberOfUsers int
		beforeCallback         func(uuid.UUID) error
		inUser                 user.User
		afterUsers             []*roomUser
		afterStatus            RoomStatus
		outChannel             chan Room
		outError               error
	}{
		{
			beforeUsers:            []*roomUser{},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 4,
			inUser:                 testUser,
			afterUsers:             []*roomUser{testRoomUser},
			afterStatus:            Open,
			outChannel:             nil,
			outError:               nil,
		},
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser), mockRoomUser(testUser)},
			beforeStatus:           Close,
			beforeMaxNumberOfUsers: 3,
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Close,
			outChannel:             nil,
			outError:               RoomCloseErr,
		},
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			beforeCallback:         func(id uuid.UUID) error { return nil },
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Close,
			outChannel:             nil,
			outError:               nil,
		},
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			beforeCallback:         func(id uuid.UUID) error { return errors.New("") },
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Open,
			outChannel:             nil,
			outError:               RoomCallbackErr,
		},
	}

	for _, c := range cases {
		room := roomImpl{
			status:           c.beforeStatus,
			users:            c.beforeUsers,
			maxNumberOfUsers: c.beforeMaxNumberOfUsers,
			callback:         c.beforeCallback,
		}
		_, err := room.JoinUser(c.inUser)
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterStatus, room.status)
		for i, ru := range room.users {
			assert.Equal(t, c.afterUsers[i].u, ru.u)
		}
	}
}

func TestLeaveUser(t *testing.T) {
	testUser := &user.UserMock{IdMock: uuid.New()}
	noExistUser := &user.UserMock{}
	cases := []struct {
		beforeUsers            []*roomUser
		beforeStatus           RoomStatus
		beforeMaxNumberOfUsers int
		inUser                 user.User
		afterUsers             []*roomUser
		afterStatus            RoomStatus
		outChannel             chan Room
		outError               error
	}{
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 4,
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser)},
			afterStatus:            Open,
			outChannel:             nil,
			outError:               nil,
		},
		{
			beforeUsers:            []*roomUser{},
			beforeStatus:           Close,
			beforeMaxNumberOfUsers: 1,
			inUser:                 nil,
			afterUsers:             []*roomUser{},
			afterStatus:            Close,
			outChannel:             nil,
			outError:               RoomCloseErr,
		},
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			inUser:                 testUser,
			afterUsers:             []*roomUser{},
			afterStatus:            Close,
			outChannel:             nil,
			outError:               nil,
		},
		{
			beforeUsers:            []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 3,
			inUser:                 noExistUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Open,
			outChannel:             nil,
			outError:               RoomUserNotFoundErr,
		},
	}

	for _, c := range cases {
		room := roomImpl{
			status:           c.beforeStatus,
			users:            c.beforeUsers,
			maxNumberOfUsers: c.beforeMaxNumberOfUsers,
		}
		err := room.LeaveUser(c.inUser)
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterStatus, room.status)
		for i, ru := range room.users {
			assert.Equal(t, c.afterUsers[i].u, ru.u)
		}
	}
}

func TestJoinUserBroadcast(t *testing.T) {
	testUser1 := &user.UserMock{IdMock: uuid.New()}
	testUser2 := &user.UserMock{IdMock: uuid.New()}
	channel1 := make(chan Room)
	testRoomUser := &roomUser{u: testUser1, c: channel1}
	cases := []struct {
		beforeUsers []*roomUser
		inUser      user.User
	}{
		{
			beforeUsers: []*roomUser{testRoomUser},
			inUser:      testUser2,
		},
	}

	for _, c := range cases {
		room := roomImpl{
			status:           Open,
			users:            c.beforeUsers,
			maxNumberOfUsers: 10,
		}
		channel2, _ := room.JoinUser(c.inUser)
		room1 := <-channel1
		room2 := <-channel2
		assert.Equal(t, room1, room2)
	}
}

func TestLeaveUserBroadcast(t *testing.T) {
	testUser1 := &user.UserMock{IdMock: uuid.New()}
	testUser2 := &user.UserMock{IdMock: uuid.New()}
	channel1 := make(chan Room)
	channel2 := make(chan Room)
	testRoomUser1 := &roomUser{u: testUser1, c: channel1}
	testRoomUser2 := &roomUser{u: testUser2, c: channel2}
	cases := []struct {
		beforeUsers []*roomUser
		inUser      user.User
	}{
		{
			beforeUsers: []*roomUser{testRoomUser1, testRoomUser2},
			inUser:      testUser2,
		},
	}

	for _, c := range cases {
		room := &roomImpl{
			status:           Open,
			users:            c.beforeUsers,
			maxNumberOfUsers: 10,
		}
		_ = room.LeaveUser(c.inUser)
		room1 := <-channel1
		room2 := <-channel2

		// can't compare as mutex state
		assert.Equal(t, room1.ID(), room.ID())
		assert.Equal(t, nil, room2)
	}
}

func mockRoomUser(u user.User) *roomUser {
	return &roomUser{u: u, c: make(chan Room)}
}
