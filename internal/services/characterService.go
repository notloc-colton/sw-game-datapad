package services

import (
	"sw-game-datapad/internal/swData"
	"sw-game-datapad/internal/vendor"
)

type CharacterService interface {
	GetCharacters() ([]swData.Character, error)
}
type characterService struct {
	Vendor vendor.VendorService
}

func NewCharacterService() *characterService {
	return &characterService{}
}
func (service *characterService) GetCharacters() ([]swData.Character, error) {
	return []swData.Character{
		{
			Name: "Luke",
			HomePlanet: swData.Planet{
				Name: "Tatooine",
			},
			Species: swData.Species{
				Name: "Human",
			},
		},
		{
			Name: "Leia",
			HomePlanet: swData.Planet{
				Name: "Alderaan",
			},
			Species: swData.Species{
				Name: "Human",
			},
		},
	}, nil
}
