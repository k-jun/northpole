package storage

import "errors"

var (
	MatchStorageMatchAlreadyExistErr = errors.New("the match id have already exist in the storage")
	MatchStorageMatchNotFound        = errors.New("the match id doesn't exist in the storage")
	MatchStorageBadParameter         = errors.New("provided match parameter is invalid")
)
