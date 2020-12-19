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
			inUser:    user.New(userId),
			outStatus: pb.MatchStatus_Availabel,
			outUsers:  []*user.User{user.New(userId)},
			outError:  nil,
		},
		{
			nowStatus: pb.MatchStatus_Availabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId)},
			inUser:    user.New(userId),
			outStatus: pb.MatchStatus_Unavailabel,
			outUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
			outError:  nil,
		},
		{
			nowStatus: pb.MatchStatus_Unavailabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
			inUser:    user.New(userId),
			outStatus: pb.MatchStatus_Unavailabel,
			outUsers:  []*user.User{},
			outError:  MatchUnavailableErr,
		},
	}

	for _, c := range cases {
		channel := make(chan Match)
		match := MatchImpl{
			id:               uuid.New(),
			status:           c.nowStatus,
			users:            c.nowUsers,
			maxNumberOfUsers: c.nowMaxNOU,
			channel:          channel,
		}
		err := match.JoinUser(c.inUser)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.outStatus, match.status)
		assert.Equal(t, c.outUsers, match.users)
		for i := 0; i < len(match.users); i++ {
			<-match.channel
		}
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
			nowUsers:  []*user.User{user.New(userId)},
			inUser:    user.New(userId),
			outStatus: pb.MatchStatus_Unavailabel,
			outUsers:  []*user.User{},
			outError:  nil,
		},
		{
			nowStatus: pb.MatchStatus_Availabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId)},
			inUser:    user.New(userId),
			outStatus: pb.MatchStatus_Availabel,
			outUsers:  []*user.User{user.New(userId), user.New(userId)},
			outError:  nil,
		},
		{
			nowStatus: pb.MatchStatus_Availabel,
			nowMaxNOU: 4,
			nowUsers:  []*user.User{user.New(userId)},
			inUser:    user.New(uuid.New()),
			outStatus: pb.MatchStatus_Unavailabel,
			outUsers:  []*user.User{},
			outError:  MatchUserNotFound,
		},
	}

	for _, c := range cases {
		channel := make(chan Match)
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
		for i := 0; i < len(match.users); i++ {
			<-match.channel
		}
	}
}
