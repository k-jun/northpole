package match

// import (
// 	"northpole/user"
// 	"testing"
//
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestNew(t *testing.T) {
// 	outId := uuid.New()
// 	_ = New(outId)
// }
//
// func TestJoin(t *testing.T) {
// 	userId := uuid.New()
// 	cases := []struct {
// 		beforeStatus MatchStatus
// 		beforeMaxNOU int
// 		beforeUsers  []user.User
// 		inUser       user.User
// 		afterStatus  MatchStatus
// 		afterUsers   []user.User
// 		outError     error
// 	}{
// 		{
// 			beforeStatus: Availabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{},
// 			inUser:       user.New(userId),
// 			afterStatus:  Availabel,
// 			afterUsers:   []user.User{user.New(userId)},
// 			outError:     nil,
// 		},
// 		{
// 			beforeStatus: Availabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{user.New(userId), user.New(userId), user.New(userId)},
// 			inUser:       user.New(userId),
// 			afterStatus:  Unavailabel,
// 			afterUsers:   []user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
// 			outError:     nil,
// 		},
// 		{
// 			beforeStatus: Unavailabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{user.New(userId), user.New(userId), user.New(userId), user.New(userId)},
// 			inUser:       user.New(userId),
// 			afterStatus:  Unavailabel,
// 			afterUsers:   []user.User{},
// 			outError:     MatchUnavailableErr,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		channel := make(chan Match)
// 		match := MatchImpl{
// 			id:               uuid.New(),
// 			status:           c.beforeStatus,
// 			users:            c.beforeUsers,
// 			maxNumberOfUsers: c.beforeMaxNOU,
// 			channel:          channel,
// 		}
// 		err := match.JoinUser(c.inUser)
// 		if err != nil && err == c.outError {
// 			continue
// 		}
// 		assert.Equal(t, c.outError, err)
// 		assert.Equal(t, c.afterStatus, match.status)
// 		assert.Equal(t, c.afterUsers, match.users)
// 		for i := 0; i < len(match.users); i++ {
// 			<-match.channel
// 		}
// 	}
// }
//
// func TestLeave(t *testing.T) {
// 	userId := uuid.New()
// 	cases := []struct {
// 		beforeStatus MatchStatus
// 		beforeMaxNOU int
// 		beforeUsers  []user.User
// 		inUser       user.User
// 		afterStatus  MatchStatus
// 		outUsers     []user.User
// 		outError     error
// 	}{
// 		{
// 			beforeStatus: Availabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{user.New(userId)},
// 			inUser:       user.New(userId),
// 			afterStatus:  Unavailabel,
// 			outUsers:     []user.User{},
// 			outError:     nil,
// 		},
// 		{
// 			beforeStatus: Availabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{user.New(userId), user.New(userId), user.New(userId)},
// 			inUser:       user.New(userId),
// 			afterStatus:  Availabel,
// 			outUsers:     []user.User{user.New(userId), user.New(userId)},
// 			outError:     nil,
// 		},
// 		{
// 			beforeStatus: Availabel,
// 			beforeMaxNOU: 4,
// 			beforeUsers:  []user.User{user.New(userId)},
// 			inUser:       user.New(uuid.New()),
// 			afterStatus:  Unavailabel,
// 			outUsers:     []user.User{},
// 			outError:     MatchUserNotFound,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		channel := make(chan Match)
// 		match := MatchImpl{
// 			channel:          channel,
// 			id:               uuid.New(),
// 			status:           c.beforeStatus,
// 			users:            c.beforeUsers,
// 			maxNumberOfUsers: c.beforeMaxNOU,
// 		}
// 		err := match.LeaveUser(c.inUser)
// 		if err != nil && err == c.outError {
// 			continue
// 		}
// 		assert.Equal(t, c.outError, err)
// 		assert.Equal(t, c.afterStatus, match.status)
// 		assert.Equal(t, c.outUsers, match.users)
// 		for i := 0; i < len(match.users); i++ {
// 			<-match.channel
// 		}
// 	}
// }
//
// func TestIsAvailabel(t *testing.T) {
// 	cases := []struct {
// 		beforeStatus MatchStatus
// 		outBool      bool
// 	}{
// 		{
// 			beforeStatus: Availabel,
// 			outBool:      true,
// 		},
// 		{
// 			beforeStatus: Unavailabel,
// 			outBool:      false,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		match := MatchImpl{status: c.beforeStatus}
// 		assert.Equal(t, c.outBool, match.IsAvailabel())
// 	}
// }
