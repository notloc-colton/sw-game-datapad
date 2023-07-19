package swData

type Character struct {
	ID         string     `json:"-"`
	Name       string     `json:"name"`
	HomePlanet Planet     `json:"homePlanet,omitempty"`
	Species    Species    `json:"species,omitempty"`
	StarShips  []StarShip `json:"starships,omitempty"`
}

func (character *Character) Initialize(starShipIds []string, speciesIds []string, homePlanetId string) {
	character.StarShips = make([]StarShip, 0)
	for _, id := range starShipIds {
		if id != "" {
			character.StarShips = append(character.StarShips, StarShip{
				ID: id,
			})
		}
	}
	if len(speciesIds) > 0 {
		if speciesIds[0] != "" {
			character.Species = Species{
				ID: speciesIds[0],
			}
		}
	}
	if homePlanetId != "" {
		character.HomePlanet = Planet{ID: homePlanetId}
	}
}