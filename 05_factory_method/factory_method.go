package factorymethod

import "fmt"

type IClub interface {
	setName(name string)
	setLocation(location string)
	getName() string
	getLocation() string
}

type Club struct {
	name     string
	location string
}

func (c *Club) setName(name string) {
	c.name = name
}

func (c *Club) getName() string {
	return c.name
}

func (c *Club) setLocation(location string) {
	c.location = location
}

func (c *Club) getLocation() string {
	return c.location
}

type Milan struct {
	Club
}

func newMilan() IClub {
	return &Milan{
		Club: Club{
			name:     "AC Milan",
			location: "Milan, Italy",
		},
	}
}

type Dortmund struct {
	Club
}

func newDortmund() IClub {
	return &Dortmund{
		Club: Club{
			name:     "Borussia Dortmund",
			location: "Dortmund, Germany",
		},
	}
}

func getClub(club string) (IClub, error) {
	if club == "Milan" {
		return newMilan(), nil
	}
	if club == "Dortmund" {
		return newDortmund(), nil
	}
	return nil, fmt.Errorf("Cannot get the club info.")
}
