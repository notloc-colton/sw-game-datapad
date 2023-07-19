package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sw-game-datapad/internal/controllers"
	"sw-game-datapad/internal/routes"
	"sw-game-datapad/internal/server"
	"sw-game-datapad/internal/services"
	"sw-game-datapad/internal/vendor"
	"sw-game-datapad/pkg/logger"
	"syscall"
	"time"
)

func main() {
	srv := server.NewServer()
	routes.AttachCharacterRoutes(&srv,
		controllers.
			NewCharacterController(services.
				NewCharacterService(vendor.
					NewVendorService("https://swapi.dev/api/"))))
	go func() {
		addr := "127.0.0.1"
		port := "8080"
		fullAddress := fmt.Sprintf("%s:%s", addr, port)
		logger.Log(logger.LogLevelInfo, "Starting server", map[string]string{
			"address": addr,
			"port":    port,
		})
		if err := srv.ListenAndServe(fullAddress); err != nil {
			fmt.Println(err)
			logger.Log(logger.LogLevelFatal, "an error was encountered while running server", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	<-quit
	logger.Log(logger.LogLevelInfo, "shutting down application")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log(logger.LogLevelFatal, "server could not shutdown gracefully; forcing shutdown", err)
	}
	select {
	case <-ctx.Done():
		logger.Log(logger.LogLevelInfo, "timeout of 5 seconds.")
	}
	logger.Log(logger.LogLevelInfo, "server exiting")
}
