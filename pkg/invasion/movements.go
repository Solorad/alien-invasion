package invasion

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/solorad/alien-invasion/pkg/models"
)

const neighboursCount = 4

// GroundAliens puts all aliens on a map without city destroying and aliens kill logic.
// This method returns data structure for simple work with aliens - a map with city name as a key and array of present
// aliens in it.
func GroundAliens(alienNames []string, cities map[string]*models.City) map[string][]*models.Alien {
	// there might be a situation when 2 aliens starts from the same city.
	// We should track such cases and destroy the city and both aliens in such case
	aliensMap := make(map[string][]*models.Alien)
	for i := range alienNames {
		var city *models.City
		// utilize the fact that in go iteration via a map is always random, and we always get a random city
		for _, v := range cities {
			city = v
			break
		}
		aliensMap[city.Name] = append(aliensMap[city.Name], &models.Alien{
			Name:     alienNames[i],
			Position: city,
		})
	}
	return aliensMap
}

// CheckCitiesAndAliens function validates state after each iteration and initial ground.
// It returns a map with all alive cities with aliens
func CheckCitiesAndAliens(stepName string, aliensMap map[string][]*models.Alien) map[string][]*models.Alien {
	result := make(map[string][]*models.Alien, len(aliensMap))
	for name, v := range aliensMap {
		if len(v) > 1 {
			CityWasDestroyedMessage(stepName, name, v)
			v[0].Position.Alive = false
		} else {
			result[name] = v
		}
	}
	// we could also add validation for `trapped` aliens - check cities graph and find all such not connected aliens
	return result
}

func MoveAliens(aliensMap map[string][]*models.Alien) map[string][]*models.Alien {
	result := make(map[string][]*models.Alien, len(aliensMap))
	for cityName, al := range aliensMap {
		if len(al) < 1 {
			fmt.Printf("error in aliensMap logic. 0 aliens in a city %s\n", cityName)
			continue
		}
		alien := al[0]
		// randomly select for each alien a new city. If there's no city in that direction OR city is already destroyed - we go clockwise
		v := rand.Int31n(neighboursCount)
		moved := false
		for i := 0; i < neighboursCount; i++ {
			p := (i + int(v)) % neighboursCount
			neighbour := alien.Position.Neighbours[p]
			// validate that the current alien's city has a connection in that direction and city there is alive
			if neighbour != nil && neighbour.Alive {
				alien.Position = neighbour
				result[neighbour.Name] = append(result[neighbour.Name], alien)
				moved = true
				break
			}
		}
		// for trapped in a city without any connections
		if !moved {
			result[cityName] = append(result[cityName], alien)
		}
	}
	return result
}

func CityWasDestroyedMessage(stepName, cityName string, aliens []*models.Alien) {
	var alienNames []string
	for _, v := range aliens {
		alienNames = append(alienNames, v.Name)
	}
	fmt.Printf("[%s] %s has been destroyed by %s!\n", stepName, cityName, strings.Join(alienNames, " and "))
}
