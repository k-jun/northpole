package match

import (
	pb "northpole/grpc"
	"northpole/user"

	"github.com/google/uuid"
)

type MatchMock struct {
	MockError  error
	MockStatus pb.MatchStatus
	MockID     uuid.UUID
}

func (m MatchMock) JoinUser(u *user.User) error {
	return m.MockError
}

func (m MatchMock) LeaveUser(u *user.User) error {
	return m.MockError
}

func (m MatchMock) ID() uuid.UUID {
	return m.MockID
}

func (m MatchMock) Status() pb.MatchStatus {
	return m.MockStatus
}

func (m MatchMock) Channel() chan Match {
	return nil
}
