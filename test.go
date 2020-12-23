package main

import (
	"errors"
	"northpole/match"
	"northpole/storage"
	"northpole/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJoinPublicMatch(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                  *user.User
		inMatch                 match.Match
		beforeMatchStorageError error
		beforeMatchError        error
		outError                error
	}{
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        nil,
			outError:                nil,
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: errors.New(""),
			beforeMatchError:        nil,
			outError:                errors.New(""),
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        errors.New(""),
			outError:                errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := match.MatchMock{IDMock: uuid2, ErrorMock: c.beforeMatchError}
		ms := storage.MatchStorageMock{MatchMock: m, ErrorMock: c.beforeMatchStorageError}
		u := matchUsecaseImpl{publicMatchStorage: ms}

		match, err := u.JoinPublicMatch(c.inUser)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, match.ID())
	}
}

func TestJoinPrivateMatch(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                  *user.User
		inMatch                 match.Match
		beforeMatchStorageError error
		beforeMatchError        error
		outError                error
	}{
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        nil,
			outError:                nil,
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: errors.New(""),
			beforeMatchError:        nil,
			outError:                errors.New(""),
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        errors.New(""),
			outError:                errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := match.MatchMock{IDMock: uuid2, ErrorMock: c.beforeMatchError}
		ms := storage.MatchStorageMock{MatchMock: m, ErrorMock: c.beforeMatchStorageError}
		u := matchUsecaseImpl{privateMatchStorage: ms}

		match, err := u.JoinPrivateMatch(c.inUser, c.inMatch)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, match.ID())
	}
}

func TestCreatePrivateMatch(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                  *user.User
		beforeMatchStorageError error
		outError                error
	}{
		{
			inUser:                  user.New(uuid1),
			beforeMatchStorageError: nil,
			outError:                nil,
		},
		{
			inUser:                  user.New(uuid1),
			beforeMatchStorageError: errors.New(""),
			outError:                errors.New(""),
		},
	}

	for _, c := range cases {
		ms := storage.MatchStorageMock{ErrorMock: c.beforeMatchStorageError}
		u := matchUsecaseImpl{privateMatchStorage: ms}

		match, err := u.CreatePrivateMatch(c.inUser)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.NotEqual(t, uuid.Nil, match.ID())
	}
}

func TestLeavePublicMatch(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                  *user.User
		inMatch                 match.Match
		beforeMatchStorageError error
		beforeMatchError        error
		outError                error
	}{
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        nil,
			outError:                nil,
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: errors.New(""),
			beforeMatchError:        nil,
			outError:                errors.New(""),
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        errors.New(""),
			outError:                errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := match.MatchMock{IDMock: uuid2, ErrorMock: c.beforeMatchError}
		ms := storage.MatchStorageMock{MatchMock: m, ErrorMock: c.beforeMatchStorageError}
		u := matchUsecaseImpl{publicMatchStorage: ms}

		match, err := u.LeavePublicMatch(c.inUser, c.inMatch)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, match.ID())
	}
}

func TestLeavePrivateMatch(t *testing.T) {
	uuid1 := uuid.New()
	cases := []struct {
		inUser                  *user.User
		inMatch                 match.Match
		beforeMatchStorageError error
		beforeMatchError        error
		outError                error
	}{
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        nil,
			outError:                nil,
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: errors.New(""),
			beforeMatchError:        nil,
			outError:                errors.New(""),
		},
		{
			inUser:                  user.New(uuid1),
			inMatch:                 match.MatchMock{},
			beforeMatchStorageError: nil,
			beforeMatchError:        errors.New(""),
			outError:                errors.New(""),
		},
	}

	for _, c := range cases {
		uuid2 := uuid.New()
		m := match.MatchMock{IDMock: uuid2, ErrorMock: c.beforeMatchError}
		ms := storage.MatchStorageMock{MatchMock: m, ErrorMock: c.beforeMatchStorageError}
		u := matchUsecaseImpl{privateMatchStorage: ms}

		match, err := u.LeavePrivateMatch(c.inUser, c.inMatch)
		if err != nil && err.Error() == c.outError.Error() {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, uuid2, match.ID())
	}
}
