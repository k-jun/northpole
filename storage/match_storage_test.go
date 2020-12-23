package storage

import (
	// pb "northpole/grpc"
	"northpole/match"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	availabel_match := match.MatchMock{IDMock: uuid1}
	unavailabel_match := match.MatchMock{IDMock: uuid2}
	cases := []struct {
		beforeMatches map[uuid.UUID]match.Match
		inMatch       match.Match
		outMatch      match.Match
		outError      error
	}{
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid1: availabel_match,
			},
			inMatch:  availabel_match,
			outMatch: availabel_match,
			outError: nil,
		},
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid1: availabel_match,
			},
			inMatch:  unavailabel_match,
			outMatch: nil,
			outError: MatchStorageMatchNotFound,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.beforeMatches}
		m, err := ms.Find(c.inMatch)
		assert.Equal(t, c.outMatch, m)
		assert.Equal(t, c.outError, err)
	}
}

func TestFindFirst(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	unavailabel_match := match.MatchMock{StatusMock: match.Unavailabel}
	availabel_match := match.MatchMock{StatusMock: match.Availabel}
	cases := []struct {
		beforeMatchs map[uuid.UUID]match.Match
		outMatch     match.Match
		outError     error
	}{
		{
			beforeMatchs: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
				uuid2: availabel_match,
			},
			outMatch: availabel_match,
			outError: nil,
		},
		{
			beforeMatchs: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
				uuid2: unavailabel_match,
			},
			outMatch: nil,
			outError: MatchStorageMatchNotFound,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.beforeMatchs}
		match, err := ms.FindFirst()
		assert.Equal(t, c.outMatch, match)
		assert.Equal(t, c.outError, err)
	}
}

func TestAdd(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	unavailabel_match := match.MatchMock{IDMock: uuid2}
	availabel_match := match.MatchMock{IDMock: uuid2}
	cases := []struct {
		beforeMatches map[uuid.UUID]match.Match
		afterMatches  map[uuid.UUID]match.Match
		inMatch       match.Match
		outError      error
	}{
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
			},
			afterMatches: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
				uuid2: availabel_match,
			},
			inMatch:  availabel_match,
			outError: nil,
		},
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid2: availabel_match,
			},
			afterMatches: map[uuid.UUID]match.Match{},
			inMatch:      availabel_match,
			outError:     MatchStorageMatchAlreadyExistErr,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.beforeMatches}
		err := ms.Add(c.inMatch)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterMatches, ms.matches)
	}
}

func TestRemove(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	availabel_match := match.MatchMock{IDMock: uuid1}
	unavailabel_match := match.MatchMock{IDMock: uuid2}
	cases := []struct {
		beforeMatches map[uuid.UUID]match.Match
		afterMatches  map[uuid.UUID]match.Match
		inMatch       match.Match
		outError      error
	}{
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid1: availabel_match,
			},
			afterMatches: map[uuid.UUID]match.Match{},
			inMatch:      availabel_match,
			outError:     nil,
		},
		{
			beforeMatches: map[uuid.UUID]match.Match{
				uuid1: availabel_match,
			},
			afterMatches: map[uuid.UUID]match.Match{},
			inMatch:      unavailabel_match,
			outError:     MatchStorageMatchNotFound,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.beforeMatches}
		err := ms.Remove(c.inMatch)
		if err != nil && err == c.outError {
			continue
		}
		assert.Equal(t, c.outError, err)
		assert.Equal(t, c.afterMatches, ms.matches)
	}
}
