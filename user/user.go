package user

import (
	"github.com/google/uuid"
)

type User interface {
	ID() uuid.UUID
}

type UserImpl struct {
	id uuid.UUID
}

func (u *UserImpl) ID() uuid.UUID {
	return u.id
}

func New(id uuid.UUID) User {
	return &UserImpl{
		id: id,
	}
}
