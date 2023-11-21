package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"wanna-be-chuck-norris-app/internal/controllers"
	"wanna-be-chuck-norris-app/internal/environment"
	"wanna-be-chuck-norris-app/internal/routes"
	"wanna-be-chuck-norris-app/internal/server"
	"wanna-be-chuck-norris-app/internal/services"
	"wanna-be-chuck-norris-app/internal/vendors/jokes"
	"wanna-be-chuck-norris-app/internal/vendors/randomName"
	"wanna-be-chuck-norris-app/pkg/httpclient"
	"wanna-be-chuck-norris-app/pkg/logger"
)

func main() {
	if err := environment.LoadEnv(); err != nil {
		panic(fmt.Sprintf("could not load env! (%v)", err))
	}
	srv := server.NewServer()
	routes.AttachCustomJokeRoutes(&srv, controllers.
		NewCustomJokerController(services.
			NewCustomJokeService(
				randomName.NewService(
					environment.ENV().RANDOM_NAME_URL,
					httpclient.StandardClient(),
					environment.ENV().NUMBER_CACHE_ENTRIES),
				jokes.NewService(
					environment.ENV().JOKE_URL,
					httpclient.StandardClient()))))
	go func() {
		addr := environment.ENV().SERVER_ADDRESS
		port := environment.ENV().SERVER_PORT
		fullAddress := fmt.Sprintf("%s:%s", addr, port)
		logger.Log(logger.LogLevelInfo, "Starting server", map[string]string{
			"address": addr,
			"port":    port,
		})
		if err := srv.ListenAndServe(fullAddress); err != nil {
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
