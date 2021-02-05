package user

type User interface {
	ID() string
}

type UserImpl struct {
	id string
}

func (u *UserImpl) ID() string {
	return u.id
}

func New(id string) User {
	return &UserImpl{
		id: id,
	}
}
