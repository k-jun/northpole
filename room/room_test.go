package room

import (
	"errors"
	"testing"

	"github.com/k-jun/northpole/user"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testUuid := "3e6d0139-5fc7-39f8-aecd-15ae26bec824"
	testMNOU := 4
	testCallback := func(id string) error { return nil }
	room := New(testUuid, testMNOU, testCallback)
	assert.Equal(t, testUuid, room.ID())
}

func TestJoinUser(t *testing.T) {
	testUser := &user.UserMock{IdMock: "b5f89105-d69d-3ce3-8bc0-0c7816990e7d"}
	testRoomUser := &roomUser{u: testUser, c: nil}
	cases := []struct {
		name                   string
		beforeUsers            []*roomUser
		beforeStatus           RoomStatus
		beforeMaxNumberOfUsers int
		beforeCallback         func(string) error
		inUser                 user.User
		afterUsers             []*roomUser
		afterStatus            RoomStatus
		outChannel             chan Room
		outError               error
	}{
		{
			name:                   "success",
			beforeUsers:            []*roomUser{},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 4,
			inUser:                 testUser,
			afterUsers:             []*roomUser{testRoomUser},
			afterStatus:            Open,
		},
		{
			name:                   "failure: room close",
			beforeUsers:            []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser), mockRoomUser(testUser)},
			beforeStatus:           Close,
			beforeMaxNumberOfUsers: 3,
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Close,
			outError:               RoomCloseErr,
		},
		{
			name:                   "success: room is full",
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			beforeCallback:         func(id string) error { return nil },
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Close,
		},
		{
			name:                   "failure: callback is failed",
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			beforeCallback:         func(id string) error { return errors.New("") },
			inUser:                 testUser,
			afterUsers:             []*roomUser{mockRoomUser(testUser), mockRoomUser(testUser)},
			afterStatus:            Open,
			outChannel:             nil,
			outError:               RoomCallbackErr,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
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
		})
	}
}

func TestLeaveUser(t *testing.T) {
	testUser := &user.UserMock{IdMock: "08762cea-6e16-3424-a2bc-664790fefa2a"}
	noExistUser := &user.UserMock{}
	cases := []struct {
		name                   string
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
			name:                   "success: 2 -> 1",
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
			name:                   "failure: room close",
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
			name:                   "success: 1 -> 0",
			beforeUsers:            []*roomUser{mockRoomUser(testUser)},
			beforeStatus:           Open,
			beforeMaxNumberOfUsers: 2,
			inUser:                 testUser,
			afterUsers:             []*roomUser{},
			afterStatus:            Close,
		},
		{
			name:                   "failure: room user not found",
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
		t.Run(c.name, func(t *testing.T) {
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
		})
	}
}

func TestJoinUserBroadcast(t *testing.T) {
	testUser1 := &user.UserMock{IdMock: "3d8b173a-b5a5-360a-a07a-582535995e13"}
	testUser2 := &user.UserMock{IdMock: "0c2551b4-e653-3db7-b522-9195b41c9dab"}
	channel1 := make(chan Room, 100)
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
		// fmt.Println("err:", err)
		room1 := <-channel1
		room2 := <-channel2
		assert.Equal(t, room1, room2)
	}
}

func TestLeaveUserBroadcast(t *testing.T) {
	testUser1 := &user.UserMock{IdMock: "5bbe0676-86ad-342c-96c5-89c2dd688e15"}
	testUser2 := &user.UserMock{IdMock: "0b058f2e-44ef-304f-a858-dd550182a93e"}
	channel1 := make(chan Room, 100)
	channel2 := make(chan Room, 100)
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

func TestRoomClose(t *testing.T) {
	testUser1 := &user.UserMock{IdMock: "20665a49-c52a-3d65-96b4-cfbc456e9031"}
	testUser2 := &user.UserMock{IdMock: "0eea27e9-65a4-33e3-93f5-ef12a42ccaee"}
	channel1 := make(chan Room)
	channel2 := make(chan Room)
	testRoomUser1 := &roomUser{u: testUser1, c: channel1}
	testRoomUser2 := &roomUser{u: testUser2, c: channel2}
	cases := []struct {
		beforeUsers []*roomUser
		afterStatus RoomStatus
	}{
		{
			beforeUsers: []*roomUser{testRoomUser1, testRoomUser2},
			afterStatus: Close,
		},
	}

	for _, c := range cases {
		room := &roomImpl{
			status:           Open,
			users:            c.beforeUsers,
			maxNumberOfUsers: 10,
		}
		_ = room.CloseRoom()
		room1 := <-channel1
		room2 := <-channel2

		// can't compare as mutex state
		assert.Equal(t, nil, room1)
		assert.Equal(t, nil, room2)
	}
}

func mockRoomUser(u user.User) *roomUser {
	return &roomUser{u: u, c: make(chan Room, 100)}
}
