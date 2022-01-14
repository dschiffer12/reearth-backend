// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewWidgetID(t *testing.T) {
	id := NewWidgetID()
	assert.NotNil(t, id)
	u, err := ulid.Parse(id.String())
	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestWidgetIDFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			result WidgetID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result WidgetID
				err    error
			}{
				result: WidgetID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result WidgetID
				err    error
			}{
				result: WidgetID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result WidgetID
				err    error
			}{
				result: WidgetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := WidgetIDFrom(tt.input)
			assert.Equal(t, tt.expected.result, result)
			if tt.expected.err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestMustWidgetID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    WidgetID
	}{
		{
			name:        "Fail:Not valid string",
			input:       "testMustFail",
			shouldPanic: true,
		},
		{
			name:        "Fail:Not valid string",
			input:       "",
			shouldPanic: true,
		},
		{
			name:        "success:valid string",
			input:       "01f2r7kg1fvvffp0gmexgy5hxy",
			shouldPanic: false,
			expected:    WidgetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.shouldPanic {
				assert.Panics(t, func() { MustBeID(tt.input) })
				return
			}
			result := MustWidgetID(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWidgetIDFromRef(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *WidgetID
	}{
		{
			name:     "Fail:Not valid string",
			input:    "testMustFail",
			expected: nil,
		},
		{
			name:     "Fail:Not valid string",
			input:    "",
			expected: nil,
		},
		{
			name:     "success:valid string",
			input:    "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: &WidgetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := WidgetIDFromRef(&tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWidgetIDFromRefID(t *testing.T) {
	id := New()
	id2 := WidgetIDFromRefID(&id)
	assert.Equal(t, id.id, id2.id)
	assert.Nil(t, WidgetIDFromRefID(nil))
	assert.Nil(t, WidgetIDFromRefID(&ID{}))
}

func TestWidgetID_ID(t *testing.T) {
	id := New()
	id2 := WidgetIDFromRefID(&id)
	assert.Equal(t, id, id2.ID())
}

func TestWidgetID_String(t *testing.T) {
	id := New()
	id2 := WidgetIDFromRefID(&id)
	assert.Equal(t, id.String(), id2.String())
	assert.Equal(t, "", WidgetID{}.String())
}

func TestWidgetID_RefString(t *testing.T) {
	id := NewWidgetID()
	assert.Equal(t, id.String(), *id.RefString())
	assert.Nil(t, WidgetID{}.RefString())
}

func TestWidgetID_GoString(t *testing.T) {
	id := New()
	id2 := WidgetIDFromRefID(&id)
	assert.Equal(t, "WidgetID("+id.String()+")", id2.GoString())
	assert.Equal(t, "WidgetID()", WidgetID{}.GoString())
}

func TestWidgetID_Ref(t *testing.T) {
	id := NewWidgetID()
	assert.Equal(t, WidgetID(id), *id.Ref())
	assert.Nil(t, (&WidgetID{}).Ref())
}

func TestWidgetID_Contains(t *testing.T) {
	id := NewWidgetID()
	id2 := NewWidgetID()
	assert.True(t, id.Contains([]WidgetID{id, id2}))
	assert.False(t, WidgetID{}.Contains([]WidgetID{id, id2, {}}))
	assert.False(t, id.Contains([]WidgetID{id2}))
}

func TestWidgetID_CopyRef(t *testing.T) {
	id := NewWidgetID().Ref()
	id2 := id.CopyRef()
	assert.Equal(t, id, id2)
	assert.NotSame(t, id, id2)
	assert.Nil(t, (*WidgetID)(nil).CopyRef())
}

func TestWidgetID_IDRef(t *testing.T) {
	id := New()
	id2 := WidgetIDFromRefID(&id)
	assert.Equal(t, &id, id2.IDRef())
	assert.Nil(t, (&WidgetID{}).IDRef())
	assert.Nil(t, (*WidgetID)(nil).IDRef())
}

func TestWidgetID_StringRef(t *testing.T) {
	id := NewWidgetID()
	assert.Equal(t, id.String(), *id.StringRef())
	assert.Nil(t, (&WidgetID{}).StringRef())
	assert.Nil(t, (*WidgetID)(nil).StringRef())
}

func TestWidgetID_MarhsalJSON(t *testing.T) {
	id := NewWidgetID()
	res, err := id.MarhsalJSON()
	assert.Nil(t, err)
	exp, _ := json.Marshal(id.String())
	assert.Equal(t, exp, res)

	res, err = (&WidgetID{}).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*WidgetID)(nil).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestWidgetID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""
	id := MustWidgetID("01f3zhkysvcxsnzepyyqtq21fb")
	id2 := &WidgetID{}
	err := id2.UnmarhsalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, id, *id2)
}

func TestWidgetID_MarshalText(t *testing.T) {
	id := New()
	res, err := WidgetIDFromRefID(&id).MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)

	res, err = (&WidgetID{}).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*WidgetID)(nil).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestWidgetID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")
	id2 := &WidgetID{}
	err := id2.UnmarshalText(text)
	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", id2.String())
}

func TestWidgetID_IsNil(t *testing.T) {
	assert.True(t, WidgetID{}.IsNil())
	assert.False(t, NewWidgetID().IsNil())
}

func TestWidgetID_IsNilRef(t *testing.T) {
	assert.True(t, WidgetID{}.Ref().IsNilRef())
	assert.True(t, (*WidgetID)(nil).IsNilRef())
	assert.False(t, NewWidgetID().Ref().IsNilRef())
}

func TestWidgetIDsToStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []WidgetID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]WidgetID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, WidgetIDsToStrings(tt.input))
		})
	}
}

func TestWidgetIDsFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected struct {
			res []WidgetID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []WidgetID
				err error
			}{
				res: make([]WidgetID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []WidgetID
				err error
			}{
				res: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
				err: nil,
			},
		},
		{
			name: "multiple elements",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []WidgetID
				err error
			}{
				res: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
				err: nil,
			},
		},
		{
			name: "error",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"x",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []WidgetID
				err error
			}{
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := WidgetIDsFrom(tc.input)
			if tc.expected.err != nil {
				assert.Equal(t, tc.expected.err, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expected.res, res)
			}
		})
	}
}

func TestWidgetIDsFromID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []ID
		expected []WidgetID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]WidgetID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WidgetIDsFromID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWidgetIDsFromIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	tests := []struct {
		name     string
		input    []*ID
		expected []WidgetID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]WidgetID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []WidgetID{MustWidgetID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []WidgetID{
				MustWidgetID(id1.String()),
				MustWidgetID(id2.String()),
				MustWidgetID(id3.String()),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WidgetIDsFromIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWidgetIDsToID(t *testing.T) {
	tests := []struct {
		name     string
		input    []WidgetID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]WidgetID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WidgetIDsToID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWidgetIDsToIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id21 := MustWidgetID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id22 := MustWidgetID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	id23 := MustWidgetID(id3.String())

	tests := []struct {
		name     string
		input    []*WidgetID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*WidgetID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*WidgetID{&id21},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*WidgetID{&id21, &id22, &id23},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WidgetIDsToIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestNewWidgetIDSet(t *testing.T) {
	WidgetIdSet := NewWidgetIDSet()
	assert.NotNil(t, WidgetIdSet)
	assert.Empty(t, WidgetIdSet.m)
	assert.Empty(t, WidgetIdSet.s)
}

func TestWidgetIDSet_Add(t *testing.T) {
	tests := []struct {
		name     string
		input    []WidgetID
		expected *WidgetIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]WidgetID, 0),
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewWidgetIDSet()
			set.Add(tc.input...)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestWidgetIDSet_AddRef(t *testing.T) {
	tests := []struct {
		name     string
		input    *WidgetID
		expected *WidgetIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &WidgetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1").Ref(),
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewWidgetIDSet()
			set.AddRef(tc.input)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestWidgetIDSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		target   *WidgetIDSet
		input    WidgetID
		expected bool
	}{
		{
			name:     "Empty Set",
			target:   &WidgetIDSet{},
			input:    MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: false,
		},
		{
			name: "Set Contains the element",
			target: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			target: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.target.Has(tc.input))
		})
	}
}

func TestWidgetIDSet_Clear(t *testing.T) {
	tests := []struct {
		name     string
		input    *WidgetIDSet
		expected *WidgetIDSet
	}{
		{
			name:     "Empty set",
			input:    &WidgetIDSet{},
			expected: &WidgetIDSet{},
		},
		{
			name:     "Nil set",
			input:    nil,
			expected: nil,
		},
		{
			name: "Contains the element",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &WidgetIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.input.Clear()
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestWidgetIDSet_All(t *testing.T) {
	tests := []struct {
		name     string
		input    *WidgetIDSet
		expected []WidgetID
	}{
		{
			name: "Empty",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{},
				s: nil,
			},
			expected: make([]WidgetID, 0),
		},
		{
			name:     "Nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "1 element",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []WidgetID{
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.All())
		})
	}
}

func TestWidgetIDSet_Clone(t *testing.T) {
	tests := []struct {
		name     string
		input    *WidgetIDSet
		expected *WidgetIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewWidgetIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewWidgetIDSet(),
			expected: NewWidgetIDSet(),
		},
		{
			name: "1 element",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			clone := tc.input.Clone()
			assert.Equal(t, tc.expected, clone)
			assert.NotSame(t, tc.input, clone)
		})
	}
}

func TestWidgetIDSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			a *WidgetIDSet
			b *WidgetIDSet
		}
		expected *WidgetIDSet
	}{
		{
			name: "Nil Set",
			input: struct {
				a *WidgetIDSet
				b *WidgetIDSet
			}{
				a: &WidgetIDSet{
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: nil,
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "Empty Set",
			input: struct {
				a *WidgetIDSet
				b *WidgetIDSet
			}{
				a: &WidgetIDSet{},
				b: &WidgetIDSet{},
			},
			expected: &WidgetIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *WidgetIDSet
				b *WidgetIDSet
			}{
				a: &WidgetIDSet{
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WidgetIDSet{},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *WidgetIDSet
				b *WidgetIDSet
			}{
				a: &WidgetIDSet{
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WidgetIDSet{
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}
