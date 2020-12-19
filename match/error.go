package match

import "errors"

var (
	MatchUnavailableErr = errors.New("the match status is unavailabel")
	MatchUserNotFound   = errors.New("the match doesn't have specified userId")
)
