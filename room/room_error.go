package room

import "errors"

var (
	RoomCloseErr        = errors.New("the room status is close")
	RoomUserNotFoundErr = errors.New("the room doesn't have specified userId")
	RoomCallbackErr     = errors.New("the room callback function failed")
)
