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
		beforeStatus pb.MatchStatus
		beforeMaxNOU int
		beforeUsers  []*user.User
		inUser       *user.User
		afterStatus  pb.MatchStatus
		afterUsers   []*user.User
		outError     error
	}{
		{
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{},
			inUser:       user.New(userId),
			afterStatus:  pb.MatchStatus_Availabel,
			afterUsers:   []*user.User{user.New(userId)},
			outError:     nil,
		},
		{
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId)},
			inUser:       user.New(userId),
			afterStatus:  pb.MatchStatus_Unavailabel,
			afterUsers:   []*user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
			outError:     nil,
		},
		{
			beforeStatus: pb.MatchStatus_Unavailabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
			inUser:       user.New(userId),
			afterStatus:  pb.MatchStatus_Unavailabel,
			afterUsers:   []*user.User{},
			outError:     MatchUnavailableErr,
		},
	}

	for _, c := range cases {
		channel := make(chan Match)
		match := MatchImpl{
			id:               uuid.New(),
			status:           c.beforeStatus,
			users:            c.beforeUsers,
			maxNumberOfUsers: c.beforeMaxNOU,
			channel:          channel,
		}
		err := match.JoinUser(c.inUser)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterStatus, match.status)
		assert.Equal(t, c.afterUsers, match.users)
		for i := 0; i < len(match.users); i++ {
			<-match.channel
		}
	}
}

func TestLeave(t *testing.T) {
	userId := uuid.New()
	cases := []struct {
		beforeStatus pb.MatchStatus
		beforeMaxNOU int
		beforeUsers  []*user.User
		inUser       *user.User
		afterStatus  pb.MatchStatus
		outUsers     []*user.User
		outError     error
	}{
		{
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{user.New(userId)},
			inUser:       user.New(userId),
			afterStatus:  pb.MatchStatus_Unavailabel,
			outUsers:     []*user.User{},
			outError:     nil,
		},
		{
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{user.New(userId), user.New(userId), user.New(userId)},
			inUser:       user.New(userId),
			afterStatus:  pb.MatchStatus_Availabel,
			outUsers:     []*user.User{user.New(userId), user.New(userId)},
			outError:     nil,
		},
		{
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 4,
			beforeUsers:  []*user.User{user.New(userId)},
			inUser:       user.New(uuid.New()),
			afterStatus:  pb.MatchStatus_Unavailabel,
			outUsers:     []*user.User{},
			outError:     MatchUserNotFound,
		},
	}

	for _, c := range cases {
		channel := make(chan Match)
		match := MatchImpl{
			channel:          channel,
			id:               uuid.New(),
			status:           c.beforeStatus,
			users:            c.beforeUsers,
			maxNumberOfUsers: c.beforeMaxNOU,
		}
		err := match.LeaveUser(c.inUser)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterStatus, match.status)
		assert.Equal(t, c.outUsers, match.users)
		for i := 0; i < len(match.users); i++ {
			<-match.channel
		}
	}
}

func TestMatchInfo(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	cases := []struct {
		beforeId     uuid.UUID
		beforeStatus pb.MatchStatus
		beforeMaxNOU int
		beforeUsers  []*user.User
		outMatchInfo *pb.MatchInfo
	}{
		{
			beforeId:     uuid1,
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 3,
			beforeUsers:  []*user.User{user.New(uuid2)},
			outMatchInfo: &pb.MatchInfo{
				Id:                   uuid1.String(),
				Status:               pb.MatchStatus_Availabel,
				MaxNumberOfUsers:     3,
				CurrentNumberOfUsers: 1,
			},
		},
		{
			beforeId:     uuid1,
			beforeStatus: pb.MatchStatus_Availabel,
			beforeMaxNOU: 3,
			beforeUsers:  []*user.User{user.New(uuid2), user.New(uuid2)},
			outMatchInfo: &pb.MatchInfo{
				Id:                   uuid1.String(),
				Status:               pb.MatchStatus_Availabel,
				MaxNumberOfUsers:     3,
				CurrentNumberOfUsers: 2,
			},
		},
	}

	for _, c := range cases {
		match := MatchImpl{
			id:               c.beforeId,
			status:           c.beforeStatus,
			users:            c.beforeUsers,
			maxNumberOfUsers: c.beforeMaxNOU,
		}
		assert.Equal(t, c.outMatchInfo, match.MatchInfo())
	}

}
