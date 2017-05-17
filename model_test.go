package model

import (
	"testing"
)

func TestMain(m *testing.M) {
	properties = nil
	m.Run()
}

func TestNewPropertyReturnsNewProperty(t *testing.T) {
	someName := "Some property name"
	property := NewProperty(someName)

	if property == nil {
		t.Fatal("Expected property, got nil")
	}

	if property.Name != someName {
		t.Fatalf("Expected '%s', got '%s'", someName, property.Name)
	}

	if property.UUID == "" {
		t.Fatalf("Expected UUID, got empty string")
	}

	if property.Rooms == nil {
		t.Fatalf("Expected empty map, got nil")
	}

	if len(property.Rooms) != 0 {
		t.Fatalf("Expected empty map, got non-empty map")
	}
}

func TestSavePropertyStoresPropertyInProperties(t *testing.T) {
	someName := "Some property name"
	given := NewProperty(someName)

	SaveProperty(given)

	result := properties[given.UUID]
	if result == nil {
		t.Fatalf("Expected %v, got nil", given)
	}

	if result.Name != given.Name {
		t.Fatalf("Expected '%s', got '%s'", given.Name, result.Name)
	}
}

func TestGetPropertyReturnsProperty(t *testing.T) {
	someName := "Some property name"
	given := NewProperty(someName)

	SaveProperty(given)

	result, err := GetProperty(given.UUID)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if result == nil {
		t.Fatalf("Expected %v, got nil", given)
	}

	if result.Name != given.Name {
		t.Fatalf("Expected '%s', got '%s'", given.Name, result.Name)
	}
}

func TestGetPropertyWithInvalidUUIDReturnsError(t *testing.T) {
	_, err := GetProperty("abc")
	if err != PropertyNotFoundError {
		t.Fatalf("Expected %v, got %v", PropertyNotFoundError, err)
	}
}

func TestGetPropertiesReturnAllProperties(t *testing.T) {
	someName := "Some property name"
	SaveProperty(NewProperty(someName + "1"))
	SaveProperty(NewProperty(someName + "2"))
	SaveProperty(NewProperty(someName + "3"))

	result := GetAllProperties()

	if result == nil {
		t.Fatalf("Expected map, got nil")
	}

	if len(result) != 3 {
		t.Fatalf("Expected 3, got %v", len(result))
	}

}
