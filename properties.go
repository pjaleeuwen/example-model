package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

var (
	properties map[string]*Property

	// ErrPropertyNotFound signals a property is not found by given UUID.
	ErrPropertyNotFound = errors.New("property not found")
	// ErrPropertyAlreadyExists signals a property by that UUID already exists.
	ErrPropertyAlreadyExists = errors.New("property already exists")
)

// NewProperty creates a new property with a random generated UUID.
func NewProperty(name string) *Property {
	return &Property{
		UUID:  uuid.NewV4().String(),
		Name:  name,
		Rooms: make(map[int]*Room),
	}
}

// GetProperty returns the property by the given UUID.
// Returns an error when no property by that UUID exists.
func GetProperty(uuid string) (*Property, error) {
	initProperties()
	property := properties[uuid]

	if property == nil {
		return nil, ErrPropertyNotFound
	}

	return property, nil
}

// GetAllProperties returns a map of all properties.
// Returns an empty map when no properties exists.
func GetAllProperties() map[string]*Property {
	initProperties()
	return properties
}

// SaveProperty stores the given property.
// Returns an error when a property by that UUID already exists.
func SaveProperty(property *Property) error {
	initProperties()
	if properties[property.UUID] != nil {
		return ErrPropertyAlreadyExists
	}

	properties[property.UUID] = property
	return nil
}

func initProperties() {
	if properties == nil {
		properties = make(map[string]*Property)
	}
}

// Property contains data together form the concept of a Property (i.e. Hotel).
type Property struct {
	UUID  string
	Name  string
	Rooms map[int]*Room
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
