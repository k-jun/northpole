package match

import "errors"

var (
	MatchMaxNumberOfUsersErr = errors.New("the match reached the max number of users")
	MatchAlreadyStartErr     = errors.New("the match have already started")
)
