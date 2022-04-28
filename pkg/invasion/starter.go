package invasion

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/solorad/alien-invasion/pkg/models"
)

const (
	cityFile = "city_map.txt"
)

func GetCities() (map[string]*models.City, error) {
	// 1. read all lines from the file
	lines, err := getCityLines()
	if err != nil {
		return nil, err
	}
	// 2. parse lines and
	return buildCityMap(lines)
}

func buildCityMap(lines []string) (map[string]*models.City, error) {
	cityMap := make(map[string]*models.City)
	for _, line := range lines {
		c, err := buildCityFromLine(cityMap, line)
		if err != nil {
			return nil, err
		}
		cityMap[c.Name] = c
	}
	return cityMap, nil
}

func buildCityFromLine(cityMap map[string]*models.City, l string) (*models.City, error) {
	parts := strings.Split(l, " ")

	cityName := parts[0]
	sourceCity, ok := cityMap[cityName]
	if !ok {
		sourceCity = &models.City{
			Name:       cityName,
			Alive:      true,
			Neighbours: [4]*models.City{},
		}
	}
	for i := 1; i < len(parts); i++ {
		err := addNeighbours(cityMap, parts[i], sourceCity)
		if err != nil {
			return nil, err
		}
	}
	return sourceCity, nil
}

func addNeighbours(cityMap map[string]*models.City, neighbourInfo string, sourceCity *models.City) error {
	vals := strings.Split(neighbourInfo, "=")
	if len(vals) != 2 {
		return errors.New("invalid address pair `" + neighbourInfo + "`")
	}
	destCity, ok := cityMap[vals[1]]
	if !ok {
		destCity = &models.City{
			Name:       vals[1],
			Alive:      true,
			Neighbours: [4]*models.City{},
		}
		cityMap[vals[1]] = destCity
	}
	switch strings.ToLower(vals[0]) {
	case "north":
		sourceCity.Neighbours[0] = destCity
	case "east":
		sourceCity.Neighbours[1] = destCity
	case "south":
		sourceCity.Neighbours[2] = destCity
	case "west":
		sourceCity.Neighbours[3] = destCity
	}
	return nil
}

// Read a map file. I use an approach with buffer and read the file line by line
// another possible approach here - ioutil.ReadFile, but might be worse by performance on
// a huge map files
func getCityLines() ([]string, error) {
	file, err := os.Open(cityFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// length of a city line is much lower than 64Kb, so we can easily use `.Scan()` method
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetAliensNum() (int, error) {
	if len(os.Args) < 2 {
		return 0, errors.New("aliens number is not provided")
	}
	return strconv.Atoi(os.Args[1])
}
