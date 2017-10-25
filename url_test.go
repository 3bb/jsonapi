package jsonapi

import (
	"net/url"
	"testing"

	"github.com/kkaribu/tchek"
)

func TestParseURL(t *testing.T) {
	// Registry
	reg := NewMockRegistry()

	tests := []struct {
		url           string
		expectedURL   URL
		expectedError bool
	}{

		{
			// 0
			url:           ``,
			expectedURL:   URL{},
			expectedError: true,
		}, {
			// 1
			url:           `https://example.com`,
			expectedURL:   URL{},
			expectedError: true,
		}, {
			// 2
			url: `mocktypes1`,
			expectedURL: URL{
				Path: `
					/mocktypes1
					?fields%5Bmocktypes1%5D=bool%2Cint%2Cint16%2Cint32%2Cint64%2Cint8%2Cstr%2Ctime%2Cto-many%2Cto-many-from-many%2Cto-many-from-one%2Cto-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint%2Cuint16%2Cuint32%2Cuint8
					&page%5Bsize%5D=10
					&page%5Bnumber%5D=1
				`,
				PathSlice:  []string{"mocktypes1"},
				Route:      "/mocktypes1",
				Type:       "col",
				FromFilter: FromFilter{},
				ResType:    "mocktypes1",
				ResID:      "",
				RelKind:    "",
				IsCol:      true,
			},
			expectedError: false,
		}, {
			// 3
			url: `https://api.example.com/mocktypes1`,
			expectedURL: URL{
				Scheme: "https",
				Host:   "api.example.com",
				Path: `
					/mocktypes1
					?fields%5Bmocktypes1%5D=bool%2Cint%2Cint16%2Cint32%2Cint64%2Cint8%2Cstr%2Ctime%2Cto-many%2Cto-many-from-many%2Cto-many-from-one%2Cto-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint%2Cuint16%2Cuint32%2Cuint8
					&page%5Bsize%5D=10
					&page%5Bnumber%5D=1
				`,
				PathSlice:  []string{"mocktypes1"},
				Route:      "/mocktypes1",
				Type:       "col",
				FromFilter: FromFilter{},
				ResType:    "mocktypes1",
				ResID:      "",
				RelKind:    "",
				IsCol:      true,
			},
			expectedError: false,
		}, {
			// 4
			url: `https://example.com/mocktypes1/mc1-1`,
			expectedURL: URL{
				Scheme: "https",
				Host:   "example.com",
				Path: `
					/mocktypes1/mc1-1
					?fields%5Bmocktypes1%5D=bool%2Cint%2Cint16%2Cint32%2Cint64%2Cint8%2Cstr%2Ctime%2Cto-many%2Cto-many-from-many%2Cto-many-from-one%2Cto-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint%2Cuint16%2Cuint32%2Cuint8
				`,
				PathSlice:  []string{"mocktypes1", "mc1-1"},
				Route:      "/mocktypes1/:id",
				Type:       "res",
				FromFilter: FromFilter{},
				ResType:    "mocktypes1",
				ResID:      "mc1-1",
				RelKind:    "",
				IsCol:      false,
			},
			expectedError: false,
		}, {
			// 5
			url: `https://example.com/mocktypes1/mc1-1/to-one`,
			expectedURL: URL{
				Scheme: "https",
				Host:   "example.com",
				Path: `
					/mocktypes1/mc1-1/to-one
					?fields%5Bmocktypes2%5D=boolptr%2Cint16ptr%2Cint32ptr%2Cint64ptr%2Cint8ptr%2Cintptr%2Cstrptr%2Ctimeptr%2Cto-many-from-many%2Cto-many-from-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint16ptr%2Cuint32ptr%2Cuint8ptr%2Cuintptr
				`,
				PathSlice: []string{"mocktypes1", "mc1-1", "to-one"},
				Route:     "/mocktypes1/:id/to-one",
				Type:      "related",
				FromFilter: FromFilter{
					Type: "mocktypes1",
					ID:   "mc1-1",
					Name: "to-one",
				},
				ResType: "mocktypes2",
				ResID:   "",
				RelKind: "related",
				Rel: Rel{
					Name:         "to-one",
					Type:         "mocktypes2",
					ToOne:        true,
					InverseName:  "",
					InverseType:  "mocktypes1",
					InverseToOne: false,
				},
				IsCol: false,
			},
			expectedError: false,
		}, {
			// 6
			url: `https://example.com/mocktypes1/mc1-1/relationships/to-many-from-one`,
			expectedURL: URL{
				Scheme: "https",
				Host:   "example.com",
				Path: `
					/mocktypes1/mc1-1/relationships/to-many-from-one
					?fields%5Bmocktypes2%5D=boolptr%2Cint16ptr%2Cint32ptr%2Cint64ptr%2Cint8ptr%2Cintptr%2Cstrptr%2Ctimeptr%2Cto-many-from-many%2Cto-many-from-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint16ptr%2Cuint32ptr%2Cuint8ptr%2Cuintptr
					&page%5Bsize%5D=10
					&page%5Bnumber%5D=1
				`,
				PathSlice: []string{"mocktypes1", "mc1-1", "relationships", "to-many-from-one"},
				Route:     "/mocktypes1/:id/relationships/to-many-from-one",
				Type:      "self",
				FromFilter: FromFilter{
					Type:        "mocktypes1",
					ID:          "mc1-1",
					Name:        "to-many-from-one",
					InverseName: "to-one-from-many",
				},
				ResType: "mocktypes2",
				ResID:   "",
				RelKind: "self",
				Rel: Rel{
					Name:         "to-many-from-one",
					Type:         "mocktypes2",
					ToOne:        false,
					InverseName:  "to-one-from-many",
					InverseType:  "mocktypes1",
					InverseToOne: true,
				},
				IsCol: true,
			},
			expectedError: false,
		}, {
			// 7
			url: `/mocktypes1/mc1-1/relationships/to-many-from-one`,
			expectedURL: URL{
				Path: `
					/mocktypes1/mc1-1/relationships/to-many-from-one
					?fields%5Bmocktypes2%5D=boolptr%2Cint16ptr%2Cint32ptr%2Cint64ptr%2Cint8ptr%2Cintptr%2Cstrptr%2Ctimeptr%2Cto-many-from-many%2Cto-many-from-one%2Cto-one-from-many%2Cto-one-from-one%2Cuint16ptr%2Cuint32ptr%2Cuint8ptr%2Cuintptr
					&page%5Bsize%5D=10
					&page%5Bnumber%5D=1
				`,
				PathSlice: []string{"mocktypes1", "mc1-1", "relationships", "to-many-from-one"},
				Route:     "/mocktypes1/:id/relationships/to-many-from-one",
				Type:      "self",
				FromFilter: FromFilter{
					Type:        "mocktypes1",
					ID:          "mc1-1",
					Name:        "to-many-from-one",
					InverseName: "to-one-from-many",
				},
				ResType: "mocktypes2",
				ResID:   "",
				RelKind: "self",
				Rel: Rel{
					Name:         "to-many-from-one",
					Type:         "mocktypes2",
					ToOne:        false,
					InverseName:  "to-one-from-many",
					InverseType:  "mocktypes1",
					InverseToOne: true,
				},
				IsCol: true,
			},
			expectedError: false,
		}, {
			// 8
			url: `/mocktypes1/mc1-1/relationships/to-many-from-one?fields[mocktypes2]=boolptr%2Cint8ptr`,
			expectedURL: URL{
				Path: `
					/mocktypes1/mc1-1/relationships/to-many-from-one
					?fields%5Bmocktypes2%5D=boolptr%2Cint8ptr
					&page%5Bsize%5D=10
					&page%5Bnumber%5D=1
				`,
				PathSlice: []string{"mocktypes1", "mc1-1", "relationships", "to-many-from-one"},
				Route:     "/mocktypes1/:id/relationships/to-many-from-one",
				Type:      "self",
				FromFilter: FromFilter{
					Type:        "mocktypes1",
					ID:          "mc1-1",
					Name:        "to-many-from-one",
					InverseName: "to-one-from-many",
				},
				ResType: "mocktypes2",
				ResID:   "",
				RelKind: "self",
				Rel: Rel{
					Name:         "to-many-from-one",
					Type:         "mocktypes2",
					ToOne:        false,
					InverseName:  "to-one-from-many",
					InverseType:  "mocktypes1",
					InverseToOne: true,
				},
				IsCol: true,
			},
			expectedError: false,
		},
	}

	for n, test := range tests {
		u, _ := url.Parse(tchek.MakeOneLineNoSpaces(test.url))
		url, err := ParseURL(reg, u)
		tchek.ErrorExpected(t, n, test.expectedError, err)

		test.expectedURL.Path = tchek.MakeOneLineNoSpaces(test.expectedURL.Path)

		if !test.expectedError {
			url.Params = nil
			tchek.AreEqual(
				t, n,
				test.expectedURL,
				*url,
			)
		}
	}
}

func TestParseParams(t *testing.T) {
	// Registry
	reg := NewMockRegistry()

	tests := []struct {
		url            string
		resType        string
		expectedParams Params
		expectedError  bool
	}{

		{
			// 0
			url: `/`,
			expectedParams: Params{
				Fields:       map[string][]string{"": nil},
				Attrs:        map[string][]Attr{},
				Rels:         map[string][]Rel{},
				RelData:      map[string][]string{},
				AttrFilters:  map[string]AttrFilter{},
				RelFilters:   map[string]RelFilter{},
				SortingRules: []string{},
				PageSize:     10,
				PageNumber:   1,
				Include:      [][]Rel{},
			},
			expectedError: false,
		}, {
			// 1
			url: `?`,
			expectedParams: Params{
				Fields:       map[string][]string{"": nil},
				Attrs:        map[string][]Attr{},
				Rels:         map[string][]Rel{},
				RelData:      map[string][]string{},
				AttrFilters:  map[string]AttrFilter{},
				RelFilters:   map[string]RelFilter{},
				SortingRules: []string{},
				PageSize:     10,
				PageNumber:   1,
				Include:      [][]Rel{},
			},
			expectedError: false,
		}, {
			// 2
			url: `
				?include=
					to-many-from-one.to-one-from-many.to-one.to-many-from-many,
					to-one-from-one.to-many-from-many
				&sort=to-many,str,,-bool
				&page[number]=3
				&sort=uint8
				&include=
					to-many-from-one,
					to-many-from-many
				&page[size]=50
			`,
			resType: "mocktypes1",
			expectedParams: Params{
				Fields: map[string][]string{
					"mocktypes1": reg.Types["mocktypes1"].Fields,
					"mocktypes2": reg.Types["mocktypes2"].Fields,
				},
				Attrs:        map[string][]Attr{},
				Rels:         map[string][]Rel{},
				RelData:      map[string][]string{},
				AttrFilters:  map[string]AttrFilter{},
				RelFilters:   map[string]RelFilter{},
				SortingRules: []string{"str", "-bool", "uint8"},
				PageSize:     50,
				PageNumber:   3,
				Include: [][]Rel{
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-many-from-many"],
					},
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-many-from-one"],
						reg.Types["mocktypes2"].Rels["to-one-from-many"],
						reg.Types["mocktypes1"].Rels["to-one"],
						reg.Types["mocktypes2"].Rels["to-many-from-many"],
					},
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-one-from-one"],
						reg.Types["mocktypes2"].Rels["to-many-from-many"],
					},
				},
			},
			expectedError: false,
		}, {
			// 3
			url: `
				?include=
					to-many-from-one.to-one-from-many.to-one.to-many-from-many%2C
					to-one-from-one.to-many-from-many
				&sort=to-many%2Cstr,%2C%2C-bool
				&page[number]=3
				&sort=uint8
				&include=
					to-many-from-one,
					to-many-from-many
				&page[size]=50
			`,
			resType: "mocktypes1",
			expectedParams: Params{
				Fields: map[string][]string{
					"mocktypes1": reg.Types["mocktypes1"].Fields,
					"mocktypes2": reg.Types["mocktypes2"].Fields,
				},
				Attrs:        map[string][]Attr{},
				Rels:         map[string][]Rel{},
				RelData:      map[string][]string{},
				AttrFilters:  map[string]AttrFilter{},
				RelFilters:   map[string]RelFilter{},
				SortingRules: []string{"str", "-bool", "uint8"},
				PageSize:     50,
				PageNumber:   3,
				Include: [][]Rel{
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-many-from-many"],
					},
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-many-from-one"],
						reg.Types["mocktypes2"].Rels["to-one-from-many"],
						reg.Types["mocktypes1"].Rels["to-one"],
						reg.Types["mocktypes2"].Rels["to-many-from-many"],
					},
					[]Rel{
						reg.Types["mocktypes1"].Rels["to-one-from-one"],
						reg.Types["mocktypes2"].Rels["to-many-from-many"],
					},
				},
			},
			expectedError: false,
		},
	}

	for n, test := range tests {
		u, err := url.Parse(tchek.MakeOneLineNoSpaces(test.url))
		tchek.UnintendedError(err)

		params, err := parseParams(reg, test.resType, u)
		tchek.ErrorExpected(t, n, test.expectedError, err)

		// Set Attrs and Rels
		for resType, fields := range test.expectedParams.Fields {
			for _, field := range fields {
				if res, ok := reg.Types[resType]; ok {
					if _, ok := res.Attrs[field]; ok {
						test.expectedParams.Attrs[resType] = append(test.expectedParams.Attrs[resType], res.Attrs[field])
					} else if _, ok := reg.Types[resType].Rels[field]; ok {
						test.expectedParams.Rels[resType] = append(test.expectedParams.Rels[resType], res.Rels[field])
					}
				}
			}
		}

		if !test.expectedError {
			// data, _ := json.MarshalIndent(test.expectedParams, "", "\t")
			// fmt.Printf("EXPECTED:\n%s\n", data)
			// data, _ = json.MarshalIndent(params, "", "\t")
			// fmt.Printf("PROVIDED:\n%s\n", data)
			tchek.AreEqual(
				t, n,
				test.expectedParams,
				*params,
			)
		}
	}
}
