// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// AssetID is an ID for Asset.
type AssetID ID

// NewAssetID generates a new AssetId.
func NewAssetID() AssetID {
	return AssetID(New())
}

// AssetIDFrom generates a new AssetID from a string.
func AssetIDFrom(i string) (nid AssetID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = AssetID(did)
	return
}

// MustAssetID generates a new AssetID from a string, but panics if the string cannot be parsed.
func MustAssetID(i string) AssetID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return AssetID(did)
}

// AssetIDFromRef generates a new AssetID from a string ref.
func AssetIDFromRef(i *string) *AssetID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := AssetID(*did)
	return &nid
}

// AssetIDFromRefID generates a new AssetID from a ref of a generic ID.
func AssetIDFromRefID(i *ID) *AssetID {
	if i == nil {
		return nil
	}
	nid := AssetID(*i)
	return &nid
}

// ID returns a domain ID.
func (d AssetID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d AssetID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d AssetID) GoString() string {
	return "id.AssetID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d AssetID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d AssetID) Ref() *AssetID {
	d2 := d
	return &d2
}

// CopyRef returns a copy of a reference.
func (d *AssetID) CopyRef() *AssetID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *AssetID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *AssetID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *AssetID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *AssetID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = AssetIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *AssetID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *AssetID) UnmarshalText(text []byte) (err error) {
	*d, err = AssetIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d AssetID) IsNil() bool {
	return ID(d).IsNil()
}

// AssetIDToKeys converts IDs into a string slice.
func AssetIDToKeys(ids []AssetID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// AssetIDsFrom converts a string slice into a ID slice.
func AssetIDsFrom(ids []string) ([]AssetID, error) {
	dids := make([]AssetID, 0, len(ids))
	for _, i := range ids {
		did, err := AssetIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// AssetIDsFromID converts a generic ID slice into a ID slice.
func AssetIDsFromID(ids []ID) []AssetID {
	dids := make([]AssetID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, AssetID(i))
	}
	return dids
}

// AssetIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func AssetIDsFromIDRef(ids []*ID) []AssetID {
	dids := make([]AssetID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, AssetID(*i))
		}
	}
	return dids
}

// AssetIDsToID converts a ID slice into a generic ID slice.
func AssetIDsToID(ids []AssetID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// AssetIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func AssetIDsToIDRef(ids []*AssetID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// AssetIDSet represents a set of AssetIDs
type AssetIDSet struct {
	m map[AssetID]struct{}
	s []AssetID
}

// NewAssetIDSet creates a new AssetIDSet
func NewAssetIDSet() *AssetIDSet {
	return &AssetIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *AssetIDSet) Add(p ...AssetID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[AssetID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []AssetID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *AssetIDSet) AddRef(p *AssetID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *AssetIDSet) Has(p AssetID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *AssetIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *AssetIDSet) All() []AssetID {
	if s == nil {
		return nil
	}
	return append([]AssetID{}, s.s...)
}

// Clone returns a cloned set
func (s *AssetIDSet) Clone() *AssetIDSet {
	if s == nil {
		return NewAssetIDSet()
	}
	s2 := NewAssetIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *AssetIDSet) Merge(s2 *AssetIDSet) *AssetIDSet {
	if s == nil {
		return nil
	}
	s3 := s.Clone()
	if s2 == nil {
		return s3
	}
	s3.Add(s2.s...)
	return s3
}
