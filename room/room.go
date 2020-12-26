package room

import (
	"sync"

	"github.com/k-jun/northpole/user"

	"github.com/google/uuid"
)

type RoomStatus string

var (
	Open  RoomStatus = "open"
	Close RoomStatus = "close"
)

type Room interface {
	ID() uuid.UUID
	JoinUser(u user.User) (chan Room, error)
	LeaveUser(u user.User) error
	IsOpen() bool
	CurrentNumberOfUsers() int
	MaxNumberOfUsers() int
}

type roomImpl struct {
	sync.Mutex

	id               uuid.UUID
	status           RoomStatus
	maxNumberOfUsers int
	users            []*roomUser
}

type roomUser struct {
	u user.User
	c chan Room
}

func New(id uuid.UUID, mnou int) Room {
	return &roomImpl{
		id:               id,
		status:           Open,
		maxNumberOfUsers: mnou,
		users:            []*roomUser{},
	}
}

func (m *roomImpl) ID() uuid.UUID {
	return m.id
}

func (m *roomImpl) IsOpen() bool {
	return m.status == Open
}

func (m *roomImpl) CurrentNumberOfUsers() int {
	return len(m.users)
}

func (m *roomImpl) MaxNumberOfUsers() int {
	return m.maxNumberOfUsers
}

func (m *roomImpl) JoinUser(inUser user.User) (chan Room, error) {
	m.Lock()
	defer m.Unlock()

	if m.status != Open {
		return nil, RoomCloseErr
	}

	channel := make(chan Room, m.maxNumberOfUsers)
	m.users = append(m.users, &roomUser{u: inUser, c: channel})
	if len(m.users) >= m.maxNumberOfUsers {
		m.status = Close
	}

	go m.broadcast(*m)
	return channel, nil
}

func (m *roomImpl) LeaveUser(outUser user.User) error {
	m.Lock()
	defer m.Unlock()

	if m.status != Open {
		return RoomCloseErr
	}

	found := false
	for i, ru := range m.users {
		if ru.u.ID() == outUser.ID() {
			close(ru.c)
			m.users = append(m.users[:i], m.users[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return RoomUserNotFoundErr
	}

	if len(m.users) == 0 {
		m.status = Close
	} else {
		go m.broadcast(*m)
	}

	return nil
}

func (m *roomImpl) broadcast(room roomImpl) {
	// TODO check valiables address
	for i := 0; i < len(m.users); i++ {
		m.users[i].c <- &room
	}
}
