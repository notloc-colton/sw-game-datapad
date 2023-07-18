package controllers

import (
	"net/http"
	"sw-game-datapad/internal/swData"

	"github.com/gin-gonic/gin"
)

type CharacterResponse struct {
	Characters []swData.Character `json:"characters,omitempty"`
}
type CharacterController struct{}

func NewCharacterController() CharacterController {
	return CharacterController{}
}

func (controller *CharacterController) GetCharacters(c *gin.Context) {
	c.JSON(http.StatusOK, CharacterResponse{
		Characters: []swData.Character{
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
		},
	})
}
