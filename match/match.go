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

	id               uuid.UUID
	status           pb.MatchStatus
	maxNumberOfUsers int64
	users            []*user.User
	channel          chan MatchImpl
}

var (
	maxNumberOfUser int64 = 4
)

func New(id uuid.UUID) Match {
	return &MatchImpl{
		id:               id,
		status:           pb.MatchStatus_Waiting,
		maxNumberOfUsers: maxNumberOfUser,
		users:            []*user.User{},
		channel:          make(chan MatchImpl),
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
	go m.broadcast()

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
	go m.broadcast()

	return nil
}

func (m *MatchImpl) broadcast() {
	for i := 0; i < len(m.users); i++ {
		m.channel <- *m
	}
}
