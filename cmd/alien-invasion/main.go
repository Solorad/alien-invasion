package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/solorad/alien-invasion/pkg/invasion"
)

const (
	maxCycles = 10_000
)

func main() {
	// 0. we use random function a lot, so put it here
	rand.Seed(time.Now().UnixNano())
	// 1.0 read aliens num from argument. I decided to use the simplest approach with pure arg get and
	// conversion to int. As an alternative to it was a `flag` library.
	alNum, err := invasion.GetAliensNum()
	if err != nil {
		log.Fatalf("error on aliens argument read: %v", err)
	}
	// 1.2 read city map.
	cities, err := invasion.GetCities()
	if err != nil {
		log.Fatalf("error on cities read: %v", err)
	}
	// 2. generate alien names
	aliens := invasion.BuildAlienNamesArray(alNum)
	// 3. put aliens on a map.
	aliensMap := invasion.GroundAliens(aliens, cities)
	// 3.a Check cities/aliens who might be destroyed
	aliensMap = invasion.CheckCitiesAndAliens("grounding", aliensMap)

	// 4. make 10.000 cycles of game
	for i := 0; i < maxCycles && len(aliensMap) > 0; i++ {
		aliensMap = invasion.MoveAliens(aliensMap)
		// after everyone made a move - check city map
		aliensMap = invasion.CheckCitiesAndAliens(fmt.Sprintf("step-%d", i+1), aliensMap)
	}

	// 5. print out city map
	invasion.PrintOutResult(cities)
}
