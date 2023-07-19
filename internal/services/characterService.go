package services

import (
	// "context"

	"sw-game-datapad/internal/swData"
	"sw-game-datapad/internal/vendor"
	"sw-game-datapad/pkg/logger"
	// "time"
)

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
type CharacterService interface {
	GetCharacters(query string) ([]swData.Character, error)
}
type characterService struct {
	Vendor vendor.VendorService
}

func NewCharacterService(vendorService vendor.VendorService) *characterService {
	return &characterService{
		Vendor: vendorService,
	}
}

// Benchmarks
// "l" with everything 6.49, 5.42, 4.73
// "l" with planets ONLY 3.43, 2.27, 5.12
// "p" with planets & species 15.1, 4.56, 4.58
// "p" with everything 18.73, 17.44
func (service *characterService) GetCharacters(query string) ([]swData.Character, error) {
	if res, err := service.Vendor.GetCharacters(query); err != nil {
		return nil, err
	} else {

		store := newCharacterDataStore(res)
		service.populatePlanets(&store)
		service.populateSpecies(&store)
		service.populateStarships(&store)
		store.PopulateCharacters()
		return store.Characters(), nil
	}
}

// Benchmarks
// "l" with everything
// "l" with planets ONLY
// "p" with planets & species 1.6, 2.24, 2.0
// "p" with everything 2.87, 4.0, 5.73
func (service *characterService) GetCharacters2(query string) ([]swData.Character, error) {
	if res, err := service.Vendor.GetCharacters(query); err != nil {
		return nil, err
	} else {
		planetCalls := make(chan PlanetCall)
		speciesCalls := make(chan SpeciesCall)
		starshipsCalls := make(chan StarshipsCall)

		store := newCharacterDataStore(res)
		service.populatePlanetsChannel(&store, planetCalls)
		service.populateSpeciesChannel(&store, speciesCalls)
		service.populateStarshipsChannel(&store, starshipsCalls)

		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()

		// var call vendor.Planet
		totalCalls := len(store.Planets()) + len(store.Species()) + len(store.StarShips())
		received := 0
		for received < totalCalls {
			select {
			case call := <-planetCalls:
				logger.Log(logger.LogLevelInfo, "received PLANET call", call)
				received++
				if call.Error == nil {
					store.AddPlanet(call.Id, call.Response)
				}
			case call := <-speciesCalls:
				logger.Log(logger.LogLevelInfo, "received SPECIES call", call)
				received++
				if call.Error == nil {
					store.AddSpecies(call.Id, call.Response)
				}
			case call := <-starshipsCalls:
				logger.Log(logger.LogLevelInfo, "received STARSHIP call", call)
				received++
				if call.Error == nil {
					store.AddStarShip(call.Id, call.Response)
				}
			}
		}
		store.PopulateCharacters()
		return store.Characters(), nil
	}
}
func (service *characterService) populatePlanets(store *characterDataStore) {
	for planetId := range store.Planets() {
		if planet, err := service.Vendor.GetPlanet(planetId); err == nil {
			store.AddPlanet(planetId, *planet)
		}
	}
}
func (service *characterService) populatePlanetsChannel(store *characterDataStore, calls chan PlanetCall) {
	for planetId := range store.Planets() {
		go func(id string) {
			planet, err := service.Vendor.GetPlanet(id)
			calls <- PlanetCall{
				Id:       id,
				Response: *planet,
				Error:    err,
			}
		}(planetId)

	}
}
func (service *characterService) populateSpecies(store *characterDataStore) {
	for speciesId := range store.Species() {
		if species, err := service.Vendor.GetSpecies(speciesId); err == nil {
			store.AddSpecies(speciesId, *species)
		}
	}
}
func (service *characterService) populateSpeciesChannel(store *characterDataStore, calls chan SpeciesCall) {
	for speciesId := range store.Species() {
		go func(id string) {
			species, err := service.Vendor.GetSpecies(id)
			calls <- SpeciesCall{
				Id:       id,
				Response: *species,
				Error:    err,
			}
		}(speciesId)
	}
}

func (service *characterService) populateStarships(store *characterDataStore) {
	for shipId := range store.StarShips() {
		if starship, err := service.Vendor.GetStarShip(shipId); err == nil {
			store.AddStarShip(shipId, *starship)
		}
	}
}
func (service *characterService) populateStarshipsChannel(store *characterDataStore, calls chan StarshipsCall) {
	for speciesId := range store.StarShips() {
		go func(id string) {
			species, err := service.Vendor.GetStarShip(id)
			calls <- StarshipsCall{
				Id:       id,
				Response: *species,
				Error:    err,
			}
		}(speciesId)
	}
}
func (service *characterService) parseCharacters(characters []vendor.Character) []swData.Character {
	characterArr := make([]swData.Character, 0)
	for _, character := range characters {
		characterArr = append(characterArr, swData.Character{
			Name: character.Name,
			HomePlanet: swData.Planet{
				Name: character.Homeworld,
			},
		})
	}
	return characterArr
}
