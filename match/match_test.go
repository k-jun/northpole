package match

import (
	pb "northpole/grpc"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	outId := uuid.New()
	match := New(outId)
	outMatch := &MatchImpl{
		id:                   outId,
		status:               pb.MatchStatus_Waiting,
		maxNumberOfUsers:     maxNumberOfUser,
		currentNumberOfUsers: 0,
	}
	assert.Equal(t, outMatch, match)
}

func TestJoin(t *testing.T) {
	cases := []struct {
		nowStatus     pb.MatchStatus
		nowMaxNOU     int64
		nowCurrentNOU int64
		outStatus     pb.MatchStatus
		outCurrentNOU int64
		outError      error
	}{
		{
			nowStatus:     pb.MatchStatus_Waiting,
			nowMaxNOU:     4,
			nowCurrentNOU: 2,
			outStatus:     pb.MatchStatus_Waiting,
			outCurrentNOU: 3,
			outError:      nil,
		},
	}

	for _, c := range cases {
		match := MatchImpl{
			id:                   uuid.New(),
			status:               c.nowStatus,
			currentNumberOfUsers: c.nowCurrentNOU,
			maxNumberOfUsers:     c.nowMaxNOU,
		}
		err := match.Join()
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.outStatus, match.status)
		assert.Equal(t, c.outCurrentNOU, match.currentNumberOfUsers)
	}
}

func TestLeave(t *testing.T) {
	cases := []struct {
		nowStatus     pb.MatchStatus
		nowMaxNOU     int64
		nowCurrentNOU int64
		outStatus     pb.MatchStatus
		outCurrentNOU int64
		outError      error
	}{
		{
			nowStatus:     pb.MatchStatus_Waiting,
			nowMaxNOU:     4,
			nowCurrentNOU: 2,
			outStatus:     pb.MatchStatus_Waiting,
			outCurrentNOU: 1,
			outError:      nil,
		},
	}

	for _, c := range cases {
		match := MatchImpl{
			id:                   uuid.New(),
			status:               c.nowStatus,
			currentNumberOfUsers: c.nowCurrentNOU,
			maxNumberOfUsers:     c.nowMaxNOU,
		}
		err := match.Leave()
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.outStatus, match.status)
		assert.Equal(t, c.outCurrentNOU, match.currentNumberOfUsers)
	}
}
