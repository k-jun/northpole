package room

import (
	"sync"

	"github.com/k-jun/northpole/user"
)

type RoomStatus string

var (
	Open  RoomStatus = "open"
	Close RoomStatus = "close"
)

type Room interface {
	ID() string
	JoinUser(u user.User) (chan Room, error)
	LeaveUser(u user.User) error
	CloseRoom() error

	// getter
	IsOpen() bool
	CurrentNumberOfUsers() int
	MaxNumberOfUsers() int
}

type roomImpl struct {
	sync.Mutex

	id               string
	status           RoomStatus
	maxNumberOfUsers int
	users            []*roomUser
	callback         func(string) error
}

type roomUser struct {
	u user.User
	c chan Room
}

func New(id string, mnou int, callback func(string) error) Room {
	return &roomImpl{
		id:               id,
		status:           Open,
		maxNumberOfUsers: mnou,
		users:            []*roomUser{},
		callback:         callback,
	}
}

func (m *roomImpl) ID() string {
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
	m.broadcast()

	if len(m.users) >= m.maxNumberOfUsers {
		if m.callback != nil {
			if err := m.callback(m.id); err != nil {
				return nil, RoomCallbackErr
			}
		}
		m.status = Close
		// let all users leave
		for _, ru := range m.users {
			close(ru.c)
		}
	}
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
		go m.broadcast()
	}

	return nil
}

func (m *roomImpl) CloseRoom() error {
	if m.callback != nil {
		if err := m.callback(m.id); err != nil {
			return RoomCallbackErr
		}
	}
	// let all users leave
	for _, ru := range m.users {
		close(ru.c)
	}
	m.status = Close
	return nil
}

func (m *roomImpl) broadcast() {
	if m.status == Close {
		return
	}
	for i := 0; i < len(m.users); i++ {
		m.users[i].c <- m
	}
}
