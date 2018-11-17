package jsonapi

import "time"

// Field ...
type Field interface {
	IsNullable() bool
	GetString() string
	GetInt() int
	GetTime() time.Time
	SetString(v string)
	SetInt(v int)
	Validate() error
}

// String ...
type String struct {
	val string
}

// IsNullable ...
func (s *String) IsNullable() bool { return false }

// GetString ...
func (s *String) GetString() string { return s.val }

// GetInt ...
func (s *String) GetInt() int { panic("jsonapi: can't call GetInt on a string-based field") }

// GetTime ...
func (s *String) GetTime() int { panic("jsonapi: can't call GetTime on a string-based field") }

// SetString ...
func (s *String) SetString(val string) { s.val = val }

// SetInt ...
func (s *String) SetInt(_ int) { panic("jsonapi: can't call SetInt on a string-based field") }

// SetTime ...
func (s *String) SetTime(_ time.Time) { panic("jsonapi: can't call SetTime on a string-based field") }

// Validate ...
func (s *String) Validate() error { return nil }
