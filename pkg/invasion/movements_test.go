package invasion

import (
	"testing"

	"github.com/solorad/alien-invasion/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCheckCitiesAndAliens(t *testing.T) {
	res := make(map[string][]*models.Alien)
	res = CheckCitiesAndAliens(res)
	assert.Equal(t, 0, len(res))

	res["City1"] = []*models.Alien{
		{
			Name: "alien-1",
			Position: &models.City{
				Name:       "City1",
				Alive:      true,
				Neighbours: [4]*models.City{},
			},
		},
	}
	res = CheckCitiesAndAliens(res)
	assert.Equal(t, 1, len(res))
	res["City2"] = []*models.Alien{
		{
			Name: "alien-2",
			Position: &models.City{
				Name:       "City2",
				Alive:      true,
				Neighbours: [4]*models.City{},
			},
		},
	}
	res = CheckCitiesAndAliens(res)
	assert.Equal(t, 2, len(res))
	res["City1"] = append(res["City1"], &models.Alien{
		Name: "alien-3",
		Position: &models.City{
			Name:       "City1",
			Alive:      true,
			Neighbours: [4]*models.City{},
		},
	})
	res = CheckCitiesAndAliens(res)
	assert.Equal(t, 1, len(res))
	assert.NotNil(t, res["City2"])
}
