package dataset

import (
	"strings"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

const (
	csvmock = `hoge,foo,bar,lat,lng
1,foo,bar,12,15`
)

func TestCSVParser(t *testing.T) {
	r := strings.NewReader(csvmock)
	p := NewCSVParser(r, "hoge.csv", ',')
	err := p.Init()
	assert.NoError(t, err)
	sceneID := id.NewSceneID()
	err = p.GuessSchema(sceneID)
	assert.NoError(t, err)

	schema, datasets, err := p.ReadAll()
	assert.NoError(t, err)

	assert.NotEmpty(t, schema)
	assert.Equal(t, "hoge.csv", schema.Name())
	assert.Equal(t, Source("file:///hoge.csv"), schema.Source())

	assert.Equal(t, 1, len(datasets))
	sfm := make(map[string]string)
	for _, sf := range schema.Fields() {
		sfm[sf.ID().String()] = sf.Name()
	}
	dsfm := make(map[string]interface{})
	for _, dsf := range datasets[0].Fields() {
		dsfm[sfm[dsf.Field().String()]] = dsf.Value().Interface()
	}
	latlng := map[string]interface{}{"lat": 12.0, "lng": 15.0}
	assert.Equal(t, map[string]interface{}{
		"hoge":     1.0,
		"foo":      "foo",
		"bar":      "bar",
		"location": latlng,
	}, dsfm)

}

func TestCSVParserCheckCompatible(t *testing.T) {
	r := strings.NewReader(csvmock)
	p := NewCSVParser(r, "hoge", ',')
	err := p.Init()
	assert.NoError(t, err)
	f1 := NewSchemaField().NewID().Name("hoge").Type(ValueTypeNumber).MustBuild()
	f2 := NewSchemaField().NewID().Name("foo").Type(ValueTypeString).MustBuild()
	f3 := NewSchemaField().NewID().Name("bar").Type(ValueTypeString).MustBuild()
	f4 := NewSchemaField().NewID().Name("location").Type(ValueTypeLatLng).MustBuild()
	fields := []*SchemaField{f1, f2, f3, f4}
	ds, err := NewSchema().NewID().Fields(fields).Build()
	assert.NoError(t, err)
	result := p.CheckCompatible(ds)
	assert.NoError(t, result)

}
