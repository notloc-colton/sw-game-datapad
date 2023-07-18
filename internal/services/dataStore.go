package services

import "sw-game-datapad/internal/swData"
//Note make all these member variables maps that you initialize when making a new data store
type characterDataStore struct {
	RequestedPlanets   []swData.Planet
	RequestedSpecies   []swData.Species
	RequestedStarShips []swData.StarShip

}
