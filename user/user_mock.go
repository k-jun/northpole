package user

import "github.com/google/uuid"

var _ User = &UserMock{}

type UserMock struct {
	IdMock    uuid.UUID
	ErrorMock error
}

func (u *UserMock) ID() uuid.UUID {
	return u.IdMock
}
