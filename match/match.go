package match

import (
	"sync"

	"github.com/google/uuid"

	pb "northpole/grpc"
	"northpole/user"
)

type Match interface {
	JoinUser(u *user.User) error
	LeaveUser(u *user.User) error
}

type MatchImpl struct {
	sync.Mutex
	cond *sync.Cond

	id               uuid.UUID
	status           pb.MatchStatus
	maxNumberOfUsers int64
	users            []*user.User
}

var (
	maxNumberOfUser int64 = 4
)

func New(id uuid.UUID) Match {
	lock := new(sync.Mutex)
	cond := sync.NewCond(lock)
	return &MatchImpl{
		cond:             cond,
		id:               id,
		status:           pb.MatchStatus_Waiting,
		maxNumberOfUsers: maxNumberOfUser,
		users:            []*user.User{},
	}
}

func (m *MatchImpl) JoinUser(inUser *user.User) error {
	m.Lock()
	defer m.Unlock()
	if int64(len(m.users)) >= m.maxNumberOfUsers {
		return MatchMaxNumberOfUsersErr
	}

	m.users = append(m.users, inUser)
	if int64(len(m.users)) >= m.maxNumberOfUsers {
		m.status = pb.MatchStatus_Start
	}
	m.cond.Broadcast()

	return nil
}

func (m *MatchImpl) LeaveUser(outUser *user.User) error {
	m.Lock()
	defer m.Unlock()
	if int64(len(m.users)) == m.maxNumberOfUsers && m.status == pb.MatchStatus_Start {
		return MatchAlreadyStartErr
	}

	found := false
	for i, user := range m.users {
		if user.ID() == outUser.ID() {
			m.users[i] = m.users[0]
			m.users = m.users[1:]
			found = true
			break
		}
	}

	if !found {
		return MatchUserNotFound
	}
	m.cond.Broadcast()

	return nil
}
