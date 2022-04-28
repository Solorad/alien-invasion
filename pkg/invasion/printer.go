package invasion

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/solorad/alien-invasion/pkg/models"
)

func PrintOutResult(cities map[string]*models.City) {
	fmt.Println("========== Cities ==========")
	for _, c := range cities {
		if c.Alive {
			paths := []string{c.Name}
			for i := 0; i < 4; i++ {
				if c.Neighbours[i] != nil && c.Neighbours[i].Alive {
					var directionName string
					switch i {
					case 0:
						directionName = "north"
					case 1:
						directionName = "east"
					case 2:
						directionName = "south"
					case 3:
						directionName = "west"
					}
					paths = append(paths, fmt.Sprintf("%s=%s", directionName, c.Neighbours[i].Name))
				}
			}
			fmt.Println(strings.Join(paths, " "))
		}
	}
}

func PrintAliensPosition(aliensMap map[string][]*models.Alien) {
	pos := make(map[string][]string)
	for cityName, aliens := range aliensMap {
		for _, a := range aliens {
			pos[cityName] = append(pos[cityName], a.Name)
		}
	}
	marshal, err := json.Marshal(pos)
	if err != nil {
		fmt.Printf("error on aliens marshalling: %v\n", err)
	} else {
		fmt.Printf("%s\n", string(marshal))
	}
}
