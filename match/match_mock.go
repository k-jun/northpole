package match

import (
	pb "northpole/grpc"
	"northpole/user"

	"github.com/google/uuid"
)

type MatchMock struct {
	ErrorMock  error
	StatusMock pb.MatchStatus
	IDMock     uuid.UUID
}

func (m MatchMock) JoinUser(u *user.User) error {
	return m.ErrorMock
}

func (m MatchMock) LeaveUser(u *user.User) error {
	return m.ErrorMock
}

func (m MatchMock) ID() uuid.UUID {
	return m.IDMock
}

func (m MatchMock) Status() pb.MatchStatus {
	return m.StatusMock
}

func (m MatchMock) Channel() chan Match {
	return nil
}

func (m MatchMock) MatchInfo() *pb.MatchInfo {
	return nil
}
