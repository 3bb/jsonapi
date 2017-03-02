package jsonapi

// Collection ...
type Collection interface {
	Elem(i int) Resource
	Add(r Resource)
	Sample() Resource

	// JSON
	MarshalJSONOptions(opts *Options) ([]byte, error)
	UnmarshalJSON(payload []byte) error
}