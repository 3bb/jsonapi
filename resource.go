package jsonapi

// Resource ...
type Resource interface {
	// Structure
	IDAndType() (string, string)
	Attrs() []Attr
	Rels() []Rel
	Attr(key string) Attr
	Rel(key string) Rel
	New() Resource

	// Read
	Get(key string) interface{}

	// Update
	SetID(id string)
	Set(key string, val interface{})

	// Read relationship
	GetToOneRel(key string) string
	GetToManyRel(key string) []string

	// Update relationship
	SetToOneRel(key string, rel string)
	SetToManyRel(key string, rels []string)

	// Validate
	Validate() []error

	// Copy
	Copy() Resource

	// JSON
	UnmarshalJSON(payload []byte) error
}
