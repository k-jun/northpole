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
	ID() uuid.UUID
	Status() pb.MatchStatus
	Channel() chan Match
}

type MatchImpl struct {
	sync.Mutex

	id               uuid.UUID
	status           pb.MatchStatus
	maxNumberOfUsers int
	users            []*user.User
	channel          chan Match
}

var (
	maxNumberOfUser = 4
)

func New(id uuid.UUID) Match {
	return &MatchImpl{
		id:               id,
		status:           pb.MatchStatus_Availabel,
		maxNumberOfUsers: maxNumberOfUser,
		users:            []*user.User{},
		channel:          make(chan Match, maxNumberOfUser),
	}
}

func (m *MatchImpl) ID() uuid.UUID {
	return m.id
}

func (m *MatchImpl) Status() pb.MatchStatus {
	return m.status
}

func (m *MatchImpl) Channel() chan Match {
	return m.channel
}

func (m *MatchImpl) JoinUser(inUser *user.User) error {
	m.Lock()
	defer m.Unlock()

	if m.status != pb.MatchStatus_Availabel {
		return MatchUnavailableErr
	}

	m.users = append(m.users, inUser)
	if len(m.users) >= m.maxNumberOfUsers {
		m.status = pb.MatchStatus_Unavailabel
	}

	go m.broadcast(*m)
	return nil
}

func (m *MatchImpl) LeaveUser(outUser *user.User) error {
	m.Lock()
	defer m.Unlock()

	if m.status != pb.MatchStatus_Availabel {
		return MatchUnavailableErr
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

	if len(m.users) == 0 {
		m.status = pb.MatchStatus_Unavailabel
		close(m.channel)
	} else {
		go m.broadcast(*m)
	}

	return nil
}

func (m *MatchImpl) broadcast(match MatchImpl) {
	for i := 0; i < len(m.users); i++ {
		m.channel <- &match
	}
}
