package controllers

import (
	"net/http"
	"sw-game-datapad/internal/services"
	"sw-game-datapad/internal/swData"

	"github.com/gin-gonic/gin"
)

type CharacterResponse struct {
	Characters []swData.Character `json:"characters,omitempty"`
}
type CharactersRequest struct {
	Query string
}
type CharacterController struct {
	Service services.CharacterService
}

func NewCharacterController(service services.CharacterService) CharacterController {
	return CharacterController{
		Service: service,
	}
}

func (controller *CharacterController) GetCharacters(c *gin.Context) {
	query := c.Query("query")
	if res, err := controller.Service.GetCharacters(query); err != nil {
		c.JSON(http.StatusInternalServerError, struct{ Error string }{
			Error: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
	// c.JSON(http.StatusOK, CharacterResponse{
	// 	Characters: []swData.Character{
	// 		{
	// 			Name: "Luke",
	// 			HomePlanet: swData.Planet{
	// 				Name: "Tatooine",
	// 			},
	// 			Species: swData.Species{
	// 				Name: "Human",
	// 			},
	// 		},
	// 		{
	// 			Name: "Leia",
	// 			HomePlanet: swData.Planet{
	// 				Name: "Alderaan",
	// 			},
	// 			Species: swData.Species{
	// 				Name: "Human",
	// 			},
	// 		},
	// 	},
	// })
}
