package swData

type Character struct {
	Name       string `json:"name"`
	HomePlanet Planet `json:"homePlanet,omitempty"`
	Species    `json:"species,omitempty"`
	StarShips  []StarShip `json:"starships,omitempty"`
}
