package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

var (
	properties map[string]*Property

	PropertyNotFoundError      = errors.New("property not found")
	PropertyAlreadyExistsError = errors.New("property already exists")
)

func NewProperty(name string) *Property {
	return &Property{
		UUID:  uuid.NewV4().String(),
		Name:  name,
		Rooms: make(map[int]*Room),
	}
}

func GetProperty(uuid string) (*Property, error) {
	initProperties()
	property := properties[uuid]

	if property == nil {
		return nil, PropertyNotFoundError
	}

	return property, nil
}

func GetAllProperties() map[string]*Property {
	initProperties()
	return properties
}

func SaveProperty(property *Property) error {
	initProperties()
	if properties[property.UUID] != nil {
		return PropertyAlreadyExistsError
	} else {
		properties[property.UUID] = property
		return nil
	}
}

func initProperties() {
	if properties == nil {
		properties = make(map[string]*Property)
	}
}

type Property struct {
	UUID  string
	Name  string
	Rooms map[int]*Room
}

type Room struct {
	UUID     string
	Number   string
	Capacity *Capacity
}

type Capacity struct {
	Kids   int
	Adults int
	Total  int
}
