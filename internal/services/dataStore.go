package services

import (
	"regexp"
	"sw-game-datapad/internal/swData"
	"sw-game-datapad/internal/vendor"
)

var parseIdRegex *regexp.Regexp = regexp.MustCompile(`\d{1,2}`)

type characterDataStore struct {
	characters []*swData.Character
	planets    map[string]swData.Planet
	species    map[string]swData.Species
	starShips  map[string]swData.StarShip
}

func newCharacterDataStore(requestedCharacters []vendor.Character) characterDataStore {
	store := characterDataStore{
		characters: make([]*swData.Character, 0),
		planets:    make(map[string]swData.Planet),
		species:    make(map[string]swData.Species),
		starShips:  make(map[string]swData.StarShip),
	}
	for _, character := range requestedCharacters {
		newCharacter := swData.Character{
			Name: character.Name,
		}
		starshipIds := parseIds(character.Starships)
		species := parseIds(character.Species)
		homePlanet := parseIdRegex.FindString(character.Homeworld)
		newCharacter.Initialize(starshipIds, species, homePlanet)
		store.characters = append(store.characters, &newCharacter)

		for _, starship := range newCharacter.StarShips {
			store.starShips[starship.ID] = swData.StarShip{}
		}
		store.planets[newCharacter.HomePlanet.ID] = swData.Planet{}
		store.species[newCharacter.Species.ID] = swData.Species{}
	}
	return store
}
func (store *characterDataStore) PopulateCharacters() {
	for _, character := range store.characters {
		if planet, found := store.planets[character.HomePlanet.ID]; found {
			character.HomePlanet = planet
		}
		if species, found := store.species[character.Species.ID]; found {
			character.Species = species
		}
		shipArr := make([]swData.StarShip, 0)
		for _, starship := range character.StarShips {
			if ship, found := store.starShips[starship.ID]; found {
				shipArr = append(shipArr, ship)
			}
		}
		character.StarShips = shipArr
	}
}
func (store *characterDataStore) Characters() []swData.Character {
	characters := make([]swData.Character, 0)
	for _, character := range store.characters {
		characters = append(characters, *character)
	}
	return characters
}
func (store *characterDataStore) Planets() map[string]swData.Planet {
	return store.planets
}
func (store *characterDataStore) AddPlanet(id string, planet vendor.Planet) {
	store.planets[id] = swData.Planet{
		ID:         id,
		Name:       planet.Name,
		Climate:    planet.Climate,
		Population: planet.Population,
	}
}
func (store *characterDataStore) Species() map[string]swData.Species {
	return store.species
}
func (store *characterDataStore) AddSpecies(id string, species vendor.Species) {
	store.species[id] = swData.Species{
		ID:              id,
		Name:            species.Name,
		AverageLifespan: species.AverageLifespan,
		Language:        species.Language,
	}
}
func (store *characterDataStore) StarShips() map[string]swData.StarShip {
	return store.starShips
}
func (store *characterDataStore) AddStarShip(id string, starship vendor.StarShip) {
	store.starShips[id] = swData.StarShip{
		ID:            id,
		Name:          starship.Name,
		CargoCapacity: starship.CargoCapacity,
		Class:         starship.StarshipClass,
	}
}
func parseIds(urls []string) []string {
	ids := make([]string, 0)
	for _, raw := range urls {
		parsed := parseIdRegex.FindString(raw)
		if parsed != "" {
			ids = append(ids, parsed)
		}
	}
	return ids
}
