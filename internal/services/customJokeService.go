package services

import (
	"context"
	"fmt"
	"wanna-be-chuck-norris-app/internal/vendors/jokes"
	"wanna-be-chuck-norris-app/internal/vendors/randomName"
	"wanna-be-chuck-norris-app/pkg/logger"
)

type CustomJokeService interface {
	GetJoke(context.Context) (string, error)
}

type customJokeService struct {
	randomNameVendor randomName.Service
	jokeVendor       jokes.Service
}

// TODO: Potentially will need to add datadog spans for cloud analytics; request context included
// for such reason
func (cjs *customJokeService) GetJoke(ctx context.Context) (string, error) {
	nameResponse, err := cjs.randomNameVendor.GetName()
	if err != nil {
		wrappedError := fmt.Errorf("error getting random name from vendor: %w", err)
		logger.Log(logger.LogLevelError, wrappedError.Error())
		return "", wrappedError
	}

	jokeResponse, err := cjs.jokeVendor.GetJoke(nameResponse.FirstName, nameResponse.LastName)
	if err != nil {
		wrappedError := fmt.Errorf("error getting joke from vendor: %w", err)
		logger.Log(logger.LogLevelError, wrappedError.Error())
		return "", wrappedError
	}

	return jokeResponse.Value.Joke, nil
}

func NewCustomJokeService(randomNameVendor randomName.Service, chuckJokeVendor jokes.Service) *customJokeService {
	return &customJokeService{
		randomNameVendor,
		chuckJokeVendor,
	}
}
