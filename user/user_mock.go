package user

var _ User = &UserMock{}

type UserMock struct {
	IdMock    string
	ErrorMock error
}

func (u *UserMock) ID() string {
	return u.IdMock
}
