package jsonapi

// Document ...
type Document struct {
	// Data
	Data interface{}

	// Included
	Included map[string]Resource

	// References
	Resources map[string]map[string]struct{}
	Links     map[string]Link

	// Relationships where data has to be included in payload
	RelData map[string][]string

	// Top-level members
	Meta map[string]interface{}

	// Errors
	Errors []Error

	// Internal
	PrePath string
}

// NewDocument ...
func NewDocument() *Document {
	return &Document{
		Included:  map[string]Resource{},
		Resources: map[string]map[string]struct{}{},
		Links:     map[string]Link{},
		RelData:   map[string][]string{},
		Meta:      map[string]interface{}{},
	}
}

// Include ...
func (d *Document) Include(res Resource) {
	id, typ := res.IDAndType()
	key := typ + " " + id

	if len(d.Included) == 0 {
		d.Included = map[string]Resource{}
	}

	if dres, ok := d.Data.(Resource); ok {
		// Check resource
		rid, rtype := dres.IDAndType()
		rkey := rid + " " + rtype

		if rkey == key {
			return
		}
	} else if col, ok := d.Data.(Collection); ok {
		// Check Collection
		_, ctyp := col.Sample().IDAndType()
		if ctyp == typ {
			for i := 0; i < col.Len(); i++ {
				rid, rtype := col.Elem(i).IDAndType()
				rkey := rid + " " + rtype

				if rkey == key {
					return
				}
			}
		}
	}

	// Check already included resources
	if _, ok := d.Included[key]; ok {
		return
	}

	d.Included[key] = res
}
