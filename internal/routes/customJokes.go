package routes

import (
	"wanna-be-chuck-norris-app/internal/controllers"
	"wanna-be-chuck-norris-app/internal/server"
)

// TODO: Add unit tests
func AttachCustomJokeRoutes(srv *server.Server, controller controllers.CustomJokeController) {
	srv.AttachRoute().GET("/joke", controller.GetJoke)
}
