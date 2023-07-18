package routes

import (
	"sw-game-datapad/internal/controllers"
	"sw-game-datapad/internal/server"
)

func AttachCharacterRoutes(srv *server.Server, controller controllers.CharacterController) {
	srv.AttachRoute().GET("/characters", controller.GetCharacters)
}