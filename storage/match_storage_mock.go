package storage

import (
	"northpole/match"
)

type MatchStorageMock struct {
	MatchMock match.Match
	ErrorMock error
}

func (ms MatchStorageMock) Find(m match.Match) (match.Match, error) {
	return ms.MatchMock, ms.ErrorMock
}

func (ms MatchStorageMock) FindFirst() (match.Match, error) {
	return ms.MatchMock, ms.ErrorMock
}

func (ms MatchStorageMock) Add(m match.Match) error {
	return ms.ErrorMock
}

func (ms MatchStorageMock) Remove(m match.Match) error {
	return ms.ErrorMock
}
