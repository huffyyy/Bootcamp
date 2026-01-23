package models

type Location struct {
	locationId int
	address    string
}

func NewLocation(locationId int, address string) *Location {
	return &Location{
		locationId,
		address,
	}
}