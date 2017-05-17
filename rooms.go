package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

var (
	// ErrRoomAlreadyExists signals a room by the given uuid
	// already exists within the given property.
	ErrRoomAlreadyExists = errors.New("room already exists")

	// ErrRoomNotFound signals a room by the given uuid
	// does not exists within the given property.
	ErrRoomNotFound = errors.New("room not found")
)

// NewRoom creates a new room with the given number and capacity.
func NewRoom(number string, capacity *Capacity) *Room {
	return &Room{
		UUID:     uuid.NewV4().String(),
		Number:   number,
		Capacity: capacity,
	}
}

// AddRoom adds the given room to the property.
// Returns an error when a room by the given UUID already exists.
func (p *Property) AddRoom(room *Room) error {
	if p.Rooms[room.UUID] != nil {
		return ErrRoomAlreadyExists
	}

	p.Rooms[room.UUID] = room
	return nil
}

// RemoveRoom removes the room with the given UUID from the property.
// Returns an error when the property doesn't contain a room with the
// given UUID.
func (p *Property) RemoveRoom(uuid string) error {
	if p.Rooms[uuid] == nil {
		return ErrRoomNotFound
	}

	return nil
}

// Room contains data together form the concept of a room (i.e. Hotel-room).
type Room struct {
	UUID     string
	Number   string
	Capacity *Capacity
}

// Capacity contains data about the capacity of a physical place (i.e. the capacity of a hotel-room).
type Capacity struct {
	Kids   int
	Adults int
	Total  int
}
