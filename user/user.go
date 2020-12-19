package user

import "github.com/google/uuid"

type User struct {
	id   uuid.UUID
	name string
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func New(id uuid.UUID, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}
