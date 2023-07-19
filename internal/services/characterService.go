package services

import (
	"context"
	"fmt"

	"sw-game-datapad/internal/swData"
	"sw-game-datapad/internal/vendor"
	"sw-game-datapad/pkg/logger"
	"time"
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

func (service *characterService) GetCharacters(query string) ([]swData.Character, error) {
	if res, err := service.Vendor.GetCharacters(query); err != nil {
		return nil, err
	} else {
		planetCalls := make(chan PlanetCall)
		speciesCalls := make(chan SpeciesCall)
		starshipsCalls := make(chan StarshipsCall)

		store := newCharacterDataStore(res)
		service.populatePlanets(&store, planetCalls)
		service.populateSpecies(&store, speciesCalls)
		service.populateStarships(&store, starshipsCalls)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		totalCalls := len(store.Planets()) + len(store.Species()) + len(store.StarShips())
		received := 0
		for received < totalCalls {
			select {
			case call := <-planetCalls:
				received++
				if call.Error == nil {
					store.AddPlanet(call.Id, call.Response)
				}
			case call := <-speciesCalls:
				received++
				if call.Error == nil {
					store.AddSpecies(call.Id, call.Response)
				}
			case call := <-starshipsCalls:
				received++
				if call.Error == nil {
					store.AddStarShip(call.Id, call.Response)
				}
			case <-ctx.Done():
				errMsg := "vendor timed out"
				logger.Log(logger.LogLevelError, errMsg)
				return nil, fmt.Errorf(errMsg)
			}
		}
		store.PopulateCharacters()
		return store.Characters(), nil
	}
}
func (service *characterService) populatePlanets(store *characterDataStore, calls chan PlanetCall) {
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
func (service *characterService) populateSpecies(store *characterDataStore, calls chan SpeciesCall) {
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

func (service *characterService) populateStarships(store *characterDataStore, calls chan StarshipsCall) {
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
