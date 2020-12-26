package room

import (
	"northpole/user"
	"sync"

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

var (
	maxNumberOfUser = 4
)

func New(id uuid.UUID) Room {
	return &roomImpl{
		id:               id,
		status:           Open,
		maxNumberOfUsers: maxNumberOfUser,
		users:            []*roomUser{},
	}
}

func (m *roomImpl) ID() uuid.UUID {
	return m.id
}

func (m *roomImpl) IsOpen() bool {
	return m.status == Open
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
			m.users[i] = m.users[0]
			m.users = m.users[1:]
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
