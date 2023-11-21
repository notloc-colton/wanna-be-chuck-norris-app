package controllers

// TODO: Add unit tests
import (
	"net/http"
	"wanna-be-chuck-norris-app/internal/services"

	"github.com/gin-gonic/gin"
)

type CustomJokeController interface {
	GetJoke(ginContext *gin.Context)
}

type customJokeController struct {
	service services.CustomJokeService
}

func (cjc *customJokeController) GetJoke(ginContext *gin.Context) {
	if res, err := cjc.service.GetJoke(ginContext.Request.Context()); err != nil {
		ginContext.JSON(http.StatusInternalServerError, struct{ Error string }{
			Error: err.Error(),
		})
	} else {
		ginContext.JSON(http.StatusOK, struct{ Joke string }{Joke: res})
	}
}

func NewCustomJokerController(service services.CustomJokeService) *customJokeController {
	return &customJokeController{
		service,
	}
}
