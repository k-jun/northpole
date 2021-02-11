package storage

import (
	"sync"

	"github.com/k-jun/northpole/room"
)

type RoomStorage interface {
	Find(room.Room) (room.Room, error)
	FindFirst() (room.Room, error)
	Add(m room.Room) error
	Remove(m room.Room) error
}

type RoomStorageImpl struct {
	sync.RWMutex
	rooms map[string]room.Room
}

func NewRoomStorage() RoomStorage {
	return &RoomStorageImpl{
		rooms: map[string]room.Room{},
	}
}

func (ms *RoomStorageImpl) Find(r room.Room) (room.Room, error) {
	ms.RLock()
	defer ms.RUnlock()

	if r.ID() == "" {
		return nil, RoomStorageBadParameter
	}
	r = ms.rooms[r.ID()]
	if r == nil {
		return nil, RoomStorageRoomNotFound
	}

	return r, nil
}

func (ms *RoomStorageImpl) FindFirst() (room.Room, error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, r := range ms.rooms {
		if r.IsOpen() {
			return r, nil
		}
	}
	return nil, RoomStorageRoomNotFound
}

func (ms *RoomStorageImpl) Add(r room.Room) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.rooms[r.ID()] != nil {
		return RoomStorageRoomAlreadyExistErr
	}
	ms.rooms[r.ID()] = r

	return nil
}

func (ms *RoomStorageImpl) Remove(r room.Room) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.rooms[r.ID()] == nil {
		return RoomStorageRoomNotFound
	}
	delete(ms.rooms, r.ID())

	return nil
}
