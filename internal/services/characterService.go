package services

import (
	"sw-game-datapad/internal/swData"
	"sw-game-datapad/internal/vendor"
)

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
		store := newCharacterDataStore(res)
		service.PopulatePlanets(&store)
		service.PopulateSpecies(&store)
		service.PopulateStarships(&store)
		store.PopulateCharacters()
		return store.Characters(), nil
	}
}
func (service *characterService) PopulatePlanets(store *characterDataStore) {
	for planetId := range store.Planets() {
		if planet, err := service.Vendor.GetPlanet(planetId); err == nil {
			store.AddPlanet(planetId, *planet)
		}
	}
}
func (service *characterService) PopulateSpecies(store *characterDataStore) {
	for speciesId := range store.Species() {
		if species, err := service.Vendor.GetSpecies(speciesId); err == nil {
			store.AddSpecies(speciesId, *species)
		}
	}
}

func (service *characterService) PopulateStarships(store *characterDataStore) {
	for shipId := range store.StarShips() {
		if starship, err := service.Vendor.GetStarShip(shipId); err == nil {
			store.AddStarShip(shipId, *starship)
		}
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
