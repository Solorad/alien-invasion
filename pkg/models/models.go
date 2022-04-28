package models

type City struct {
	Name       string
	Alive      bool
	Neighbours [4]*City // array of length 4. Those are cities on North, East, South, and West correspondingly
}

type Alien struct {
	Name     string
	Position *City
}
