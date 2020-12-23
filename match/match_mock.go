package match

// import (
// 	"northpole/user"
//
// 	"github.com/google/uuid"
// )
//
// var _ Match = MatchMock{}
//
// type MatchMock struct {
// 	ErrorMock  error
// 	StatusMock MatchStatus
// 	IDMock     uuid.UUID
// }
//
// func (m MatchMock) JoinUser(u user.User) error {
// 	return m.ErrorMock
// }
//
// func (m MatchMock) LeaveUser(u user.User) error {
// 	return m.ErrorMock
// }
//
// func (m MatchMock) ID() uuid.UUID {
// 	return m.IDMock
// }
//
// func (m MatchMock) IsAvailabel() bool {
// 	return m.StatusMock == Availabel
// }
//
// func (m MatchMock) Channel() chan Match {
// 	return nil
// }
