package storage

import (
	pb "northpole/grpc"
	"northpole/match"
	"sync"

	"github.com/google/uuid"
)

type MatchStorage interface {
	Find(id uuid.UUID) match.Match
	FindFirst() match.Match
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

func (ms *matchStorageImpl) Find(matchId uuid.UUID) match.Match {
	ms.RLock()
	defer ms.RUnlock()

	return ms.matches[matchId]
}

func (ms *matchStorageImpl) FindFirst() match.Match {
	for _, match := range ms.matches {
		if match.Status() == pb.MatchStatus_Availabel {
			return match
		}
	}
	return nil
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
