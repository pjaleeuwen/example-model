package model

import "testing"

func TestNewRoomCreatesNewRoom(t *testing.T) {
	Before()

	someNumber := "1"
	someCapacity := &Capacity{
		Kids:   2,
		Adults: 2,
		Total:  4,
	}

	room := NewRoom(someNumber, someCapacity)

	if room == nil {
		t.Fatalf("Expected room, got nil")
	}

	if room.UUID == "" {
		t.Fatalf("Expected UUID, got empty string")
	}

	if room.Number != someNumber {
		t.Fatalf("Expected %v, got %v", someNumber, room.Number)
	}

	if room.Capacity.Kids != someCapacity.Kids {
		t.Fatalf("Expected %v, got %v", someCapacity.Kids, room.Capacity.Kids)
	}

	if room.Capacity.Adults != someCapacity.Adults {
		t.Fatalf("Expected %v, got %v", someCapacity.Adults, room.Capacity.Adults)
	}

	if room.Capacity.Total != someCapacity.Total {
		t.Fatalf("Expected %v, got %v", someCapacity.Total, room.Capacity.Total)
	}
}

func TestAddRoomAddsRoomToProperty(t *testing.T) {
	Before()

	somePropertyName := "Example property"
	someRoomNumber := "1"
	someRoomCapacity := &Capacity{
		Kids:   2,
		Adults: 2,
		Total:  4,
	}

	property := NewProperty(somePropertyName)
	room := NewRoom(someRoomNumber, someRoomCapacity)
	_ = property.AddRoom(room)

	result := property.Rooms[room.UUID]

	if result == nil {
		t.Fatalf("Expected room, got nil")
	}

	if result.Number != someRoomNumber {
		t.Fatalf("Expected %v, got %v", someRoomNumber, result.Number)
	}
}
