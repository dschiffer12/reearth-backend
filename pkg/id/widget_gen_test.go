// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewWidgetID(t *testing.T) {
	id := NewWidgetID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestWidgetIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
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
				WidgetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result WidgetID
				err    error
			}{
				WidgetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result WidgetID
				err    error
			}{
				WidgetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := WidgetIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustWidgetID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
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
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.shouldPanic {
				assert.Panics(tt, func() { MustBeID(tc.input) })
				return
			}
			result := MustWidgetID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestWidgetIDFromRef(t *testing.T) {
	testCases := []struct {
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
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := WidgetIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestWidgetIDFromRefID(t *testing.T) {
	id := New()

	subId := WidgetIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestWidgetID_ID(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestWidgetID_String(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestWidgetID_GoString(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.WidgetID("+id.String()+")")
}

func TestWidgetID_RefString(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestWidgetID_Ref(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestWidgetID_CopyRef(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestWidgetID_IDRef(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestWidgetID_StringRef(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestWidgetID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestWidgetID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &WidgetID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestWidgetID_MarshalText(t *testing.T) {
	id := New()
	subId := WidgetIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestWidgetID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &WidgetID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestWidgetID_IsNil(t *testing.T) {
	subId := WidgetID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *WidgetIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestWidgetIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, WidgetIDToKeys(tc.input))
		})
	}

}

func TestWidgetIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
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
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.expected.err != nil {
				_, err := WidgetIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := WidgetIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestWidgetIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := WidgetIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestWidgetIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := WidgetIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestWidgetIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := WidgetIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestWidgetIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustWidgetID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustWidgetID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustWidgetID(id3.String())

	testCases := []struct {
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
			input:    []*WidgetID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*WidgetID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := WidgetIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
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
	t.Parallel()

	testCases := []struct {
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
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
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
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
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
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewWidgetIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestWidgetIDSet_AddRef(t *testing.T) {
	t.Parallel()

	WidgetId := MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
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
			input: &WidgetId,
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewWidgetIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestWidgetIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			WidgetIDSet
			WidgetID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				WidgetIDSet
				WidgetID
			}{WidgetIDSet: WidgetIDSet{}, WidgetID: MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				WidgetIDSet
				WidgetID
			}{WidgetIDSet: WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, WidgetID: MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				WidgetIDSet
				WidgetID
			}{WidgetIDSet: WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, WidgetID: MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.WidgetIDSet.Has(tc.input.WidgetID))
		})
	}
}

func TestWidgetIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    WidgetIDSet
		expected WidgetIDSet
	}{
		{
			name:  "Empty Set",
			input: WidgetIDSet{},
			expected: WidgetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: WidgetIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			set := tc.input
			p := &set
			p.Clear()
			assert.Equal(tt, tc.expected, *p)
		})
	}
}

func TestWidgetIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *WidgetIDSet
		expected []WidgetID
	}{
		{
			name: "Empty slice",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{},
				s: nil,
			},
			expected: make([]WidgetID, 0),
		},
		{
			name: "1 element",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.All())
		})
	}
}

func TestWidgetIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
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
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			clone := tc.input.Clone()
			assert.Equal(tt, tc.expected, clone)
			assert.False(tt, tc.input == clone)
		})
	}
}

func TestWidgetIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *WidgetIDSet
			b *WidgetIDSet
		}
		expected *WidgetIDSet
	}{
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
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WidgetIDSet{},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
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
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WidgetIDSet{
					m: map[WidgetID]struct{}{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{}},
					s: []WidgetID{MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &WidgetIDSet{
				m: map[WidgetID]struct{}{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
				},
				s: []WidgetID{
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWidgetID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}