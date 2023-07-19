package services

import "sw-game-datapad/internal/vendor"

type PlanetCall struct {
	Id       string
	Response vendor.Planet
	Error    error
}
type SpeciesCall struct {
	Id       string
	Response vendor.Species
	Error    error
}
type StarshipsCall struct {
	Id       string
	Response vendor.StarShip
	Error    error
}
