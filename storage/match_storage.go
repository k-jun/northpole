package storage

import (
	"northpole/match"
	"sync"

	"github.com/google/uuid"
)

type MatchStorage interface {
	Find(match.Match) (match.Match, error)
	FindFirst() (match.Match, error)
	Add(m match.Match) error
	Remove(m match.Match) error
}

type matchStorageImpl struct {
	sync.RWMutex
	matches map[uuid.UUID]match.Match
}

func NewMatchStorage() MatchStorage {
	return &matchStorageImpl{
		matches: map[uuid.UUID]match.Match{},
	}
}

func (ms *matchStorageImpl) Find(m match.Match) (match.Match, error) {
	ms.RLock()
	defer ms.RUnlock()

	if m.ID() == uuid.Nil {
		return nil, MatchStorageBadParameter
	}
	m = ms.matches[m.ID()]
	if m == nil {
		return nil, MatchStorageMatchNotFound
	}

	return m, nil
}

func (ms *matchStorageImpl) FindFirst() (match.Match, error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, match := range ms.matches {
		if match.IsAvailabel() {
			return match, nil
		}
	}
	return nil, MatchStorageMatchNotFound
}

func (ms *matchStorageImpl) Add(m match.Match) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.matches[m.ID()] != nil {
		return MatchStorageMatchAlreadyExistErr
	}
	ms.matches[m.ID()] = m

	return nil
}

func (ms *matchStorageImpl) Remove(m match.Match) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.matches[m.ID()] == nil {
		return MatchStorageMatchNotFound
	}
	delete(ms.matches, m.ID())

	return nil
}
