package match

import (
	pb "northpole/grpc"
	"northpole/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	outId := uuid.New()
	_ = New(outId)
}

func TestJoin(t *testing.T) {
	userId := uuid.New()
	cases := []struct {
		nowStatus pb.MatchStatus
		nowMaxNOU int
		nowUsers  []*user.User
		inUser    *user.User
		outStatus pb.MatchStatus
		outUsers  []*user.User
		outError  error
	}{
		{
			nowStatus: pb.MatchStatus_Availabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{},
			inUser:    user.New(userId, "Major"),
			outStatus: pb.MatchStatus_Availabel,
			outUsers:  []*user.User{user.New(userId, "Major")},
			outError:  nil,
		},
	}

	for _, c := range cases {
		match := MatchImpl{
			id:               uuid.New(),
			status:           c.nowStatus,
			users:            c.nowUsers,
			maxNumberOfUsers: c.nowMaxNOU,
		}
		err := match.JoinUser(c.inUser)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.outStatus, match.status)
		assert.Equal(t, c.outUsers, match.users)
	}
}

func TestLeave(t *testing.T) {
	userId := uuid.New()
	cases := []struct {
		nowStatus pb.MatchStatus
		nowMaxNOU int
		nowUsers  []*user.User
		inUser    *user.User
		outStatus pb.MatchStatus
		outUsers  []*user.User
		outError  error
	}{
		{
			nowStatus: pb.MatchStatus_Availabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{user.New(userId, "Major")},
			inUser:    user.New(userId, "Major"),
			outStatus: pb.MatchStatus_Unavailabel,
			outUsers:  []*user.User{},
			outError:  nil,
		},
	}

	for _, c := range cases {
		channel := make(chan MatchImpl)
		match := MatchImpl{
			channel:          channel,
			id:               uuid.New(),
			status:           c.nowStatus,
			users:            c.nowUsers,
			maxNumberOfUsers: c.nowMaxNOU,
		}
		err := match.LeaveUser(c.inUser)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.outStatus, match.status)
		assert.Equal(t, c.outUsers, match.users)
	}
}
