package jsonapi

import (
	"testing"

	"github.com/kkaribu/tchek"
)

// TestDocument ...
func TestDocument(t *testing.T) {
	pl1 := Document{}

	tchek.AreEqual(t, "empty", nil, pl1.Data)
}
