package storage

import (
	"sync"

	"github.com/k-jun/northpole/room"

	"github.com/google/uuid"
)

type RoomStorage interface {
	Find(room.Room) (room.Room, error)
	FindFirst() (room.Room, error)
	Add(m room.Room) error
	Remove(m room.Room) error
}

type RoomStorageImpl struct {
	sync.RWMutex
	rooms map[uuid.UUID]room.Room
}

func NewRoomStorage() RoomStorage {
	return &RoomStorageImpl{
		rooms: map[uuid.UUID]room.Room{},
	}
}

func (ms *RoomStorageImpl) Find(m room.Room) (room.Room, error) {
	ms.RLock()
	defer ms.RUnlock()

	if m.ID() == uuid.Nil {
		return nil, RoomStorageBadParameter
	}
	m = ms.rooms[m.ID()]
	if m == nil {
		return nil, RoomStorageRoomNotFound
	}

	return m, nil
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

func (ms *RoomStorageImpl) Add(m room.Room) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.rooms[m.ID()] != nil {
		return RoomStorageRoomAlreadyExistErr
	}
	ms.rooms[m.ID()] = m

	return nil
}

func (ms *RoomStorageImpl) Remove(m room.Room) error {
	ms.Lock()
	defer ms.Unlock()

	if ms.rooms[m.ID()] == nil {
		return RoomStorageRoomNotFound
	}
	delete(ms.rooms, m.ID())

	return nil
}
