package match

import "errors"

var (
	MatchMaxNumberOfUsersErr = errors.New("the match reached the max number of users")
	MatchAlreadyStartErr     = errors.New("the match have already started")
	MatchUserNotFound        = errors.New("the match doesn't have specified userId")
)
