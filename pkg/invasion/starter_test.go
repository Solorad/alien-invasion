package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCityMap(t *testing.T) {
	cityMap, err := buildCityMap([]string{
		"Foo north=Bar west=Baz south=Qu-ux",
		"Bar south=Foo west=Bee",
		"Bee east=Bar",
		"Baz east=Foo",
	})
	assert.Nil(t, err)
	fooCity := cityMap["Foo"]
	assert.NotNil(t, fooCity)
	assert.Equal(t, "Foo", fooCity.Name)
	assert.True(t, fooCity.Alive)
	assert.NotNil(t, fooCity.Neighbours[0])
	assert.Equal(t, "Bar", fooCity.Neighbours[0].Name)
	assert.Nil(t, fooCity.Neighbours[1])
	assert.NotNil(t, fooCity.Neighbours[2])
	assert.Equal(t, "Qu-ux", fooCity.Neighbours[2].Name)
	assert.NotNil(t, fooCity.Neighbours[3])
	assert.Equal(t, "Baz", fooCity.Neighbours[3].Name)
}
