package match

import (
	"github.com/google/uuid"

	pb "northpole/grpc"
)

type Match interface {
	Join() error
	Leave() error
}

type MatchImpl struct {
	id                   uuid.UUID
	status               pb.MatchStatus
	maxNumberOfUsers     int64
	currentNumberOfUsers int64
	// channel             chan pb.MatchStatus
}

func (m *MatchImpl) Join() error {
	if m.currentNumberOfUsers >= m.maxNumberOfUsers {
		return MatchMaxNumberOfUsersErr
	}

	m.currentNumberOfUsers += 1
	if m.currentNumberOfUsers == m.maxNumberOfUsers {
		m.status = pb.MatchStatus_Start
	}

	return nil

}

func (m *MatchImpl) Leave() error {
	if m.currentNumberOfUsers == m.maxNumberOfUsers && m.status == pb.MatchStatus_Start {
		return MatchAlreadyStartErr
	}

	m.currentNumberOfUsers -= 1
	return nil
}

var (
	maxNumberOfUser int64 = 4
)

func New(id uuid.UUID) Match {
	return &MatchImpl{
		id:                   id,
		status:               pb.MatchStatus_Waiting,
		maxNumberOfUsers:     maxNumberOfUser,
		currentNumberOfUsers: 0,
	}
}
