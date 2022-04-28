package invasion

import (
	"testing"

	"github.com/solorad/alien-invasion/pkg/models"

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

func TestBuildCityFromLine(t *testing.T) {
	line := "San Francisco north=Batam east=Bandung west=Lille"
	cityMap := make(map[string]*models.City)
	city, err := buildCityFromLine(cityMap, line)
	assert.Nil(t, err)
	assert.Equal(t, "San Francisco", city.Name)
	assert.True(t, city.Alive)
	assert.Equal(t, "Batam", city.Neighbours[0].Name)
	assert.Equal(t, "Bandung", city.Neighbours[1].Name)
	assert.Nil(t, city.Neighbours[2])
	assert.Equal(t, "Lille", city.Neighbours[3].Name)
}
