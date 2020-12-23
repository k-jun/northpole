package main

import (
	"northpole/match"
	"northpole/storage"
	"northpole/user"
)

type NorthPole interface {
	JoinPublicMatch(user.User) (match.Match, error)
	JoinPrivateMatch(user.User, match.Match) (match.Match, error)
	CreatePrivateMatch(user.User) (match.Match, error)
	LeavePublicMatch(user.User, match.Match) (match.Match, error)
	LeavePrivateMatch(user.User, match.Match) (match.Match, error)
}

type NorthPoleImpl struct {
	publicMatchStorage  storage.MatchStorage
	privateMatchStorage storage.MatchStorage
}

// type matchUsecaseImpl struct {
// 	publicMatchStorage  storage.MatchStorage
// 	privateMatchStorage storage.MatchStorage
// }
//
// func NewMatchUsecase(pubms storage.MatchStorage, prims storage.MatchStorage) MatchUsecase {
// 	return &matchUsecaseImpl{
// 		publicMatchStorage:  pubms,
// 		privateMatchStorage: prims,
// 	}
// }
//
// func (mu *matchUsecaseImpl) JoinPublicMatch(u *user.User) (match.Match, error) {
// 	m, err := mu.publicMatchStorage.FindFirst()
// 	if err != nil {
// 		if err == storage.MatchStorageMatchNotFound {
// 			m = match.New(utils.NewUUID())
// 			if err = mu.publicMatchStorage.Add(m); err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			return nil, err
// 		}
// 	}
//
// 	if err = m.JoinUser(u); err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
//
// func (mu *matchUsecaseImpl) JoinPrivateMatch(u *user.User, m match.Match) (match.Match, error) {
// 	m, err := mu.privateMatchStorage.Find(m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = m.JoinUser(u)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
//
// func (mu *matchUsecaseImpl) CreatePrivateMatch(u *user.User) (match.Match, error) {
// 	m := match.New(utils.NewUUID())
// 	if err := mu.privateMatchStorage.Add(m); err != nil {
// 		return nil, err
// 	}
// 	if err := m.JoinUser(u); err != nil {
// 		return nil, err
// 	}
//
// 	return m, nil
// }
//
// func (mu *matchUsecaseImpl) LeavePublicMatch(u *user.User, m match.Match) (match.Match, error) {
// 	m, err := mu.publicMatchStorage.Find(m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err = m.LeaveUser(u); err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
//
// func (mu *matchUsecaseImpl) LeavePrivateMatch(u *user.User, m match.Match) (match.Match, error) {
// 	m, err := mu.privateMatchStorage.Find(m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err = m.LeaveUser(u); err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
