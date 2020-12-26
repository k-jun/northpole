package storage

import "errors"

var (
	RoomStorageRoomAlreadyExistErr = errors.New("the room id have already exist in the storage")
	RoomStorageRoomNotFound        = errors.New("the room id doesn't exist in the storage")
	RoomStorageBadParameter        = errors.New("provided room parameter is invalid")
)
