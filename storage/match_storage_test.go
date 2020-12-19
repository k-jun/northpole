package storage

import (
	pb "northpole/grpc"
	"northpole/match"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	testMatchId := uuid.New()
	mockMatch := match.New(testMatchId)
	cases := []struct {
		nowMatchs map[uuid.UUID]match.Match
		inMatchId uuid.UUID
		outMatch  match.Match
	}{
		{
			nowMatchs: map[uuid.UUID]match.Match{
				testMatchId: mockMatch,
			},
			inMatchId: testMatchId,
			outMatch:  mockMatch,
		},
		{
			nowMatchs: map[uuid.UUID]match.Match{
				testMatchId: mockMatch,
			},
			inMatchId: uuid.New(),
			outMatch:  nil,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.nowMatchs}
		match := ms.Find(c.inMatchId)
		assert.Equal(t, c.outMatch, match)
	}
}

func TestFindFirst(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	unavailabel_match := match.MatchMock{MockStatus: pb.MatchStatus_Unavailabel}
	availabel_match := match.MatchMock{MockStatus: pb.MatchStatus_Availabel}
	cases := []struct {
		nowMatchs map[uuid.UUID]match.Match
		outMatch  match.Match
	}{
		{
			nowMatchs: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
				uuid2: availabel_match,
			},
			outMatch: availabel_match,
		},
		{
			nowMatchs: map[uuid.UUID]match.Match{
				uuid1: unavailabel_match,
				uuid2: unavailabel_match,
			},
			outMatch: nil,
		},
	}

	for _, c := range cases {
		ms := matchStorageImpl{matches: c.nowMatchs}
		match := ms.FindFirst()
		assert.Equal(t, c.outMatch, match)
	}
}

func TestAdd(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	unavailabel_match := match.MatchMock{MockID: uuid2}
	availabel_match := match.MatchMock{MockID: uuid2}
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
	availabel_match := match.MatchMock{MockID: uuid1}
	unavailabel_match := match.MatchMock{MockID: uuid2}
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
