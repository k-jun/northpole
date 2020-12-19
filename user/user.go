package user

import "github.com/google/uuid"

type User struct {
	id uuid.UUID
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func New(id uuid.UUID) *User {
	return &User{
		id: id,
	}
}
