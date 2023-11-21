package environment

// TODO: Add unit tests
import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnv loads env variables from shell (or from .env if APP_ENVIRONMENT is LOCAL)
func LoadEnv() error {
	appEnvironment := strings.ToUpper(os.Getenv("APP_ENVIRONMENT"))
	if appEnvironment == "" || appEnvironment == localEnvironment {
		fmt.Println("Loading from .env")
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	newEnv := env{
		APP_ENVIRONMENT: os.Getenv("APP_ENVIRONMENT"),
		JOKE_URL:        os.Getenv("JOKE_URL"),
		RANDOM_NAME_URL: os.Getenv("RANDOM_NAME_URL"),
		SERVER_ADDRESS:  os.Getenv("SERVER_ADDRESS"),
		SERVER_PORT:     os.Getenv("SERVER_PORT"),
	}
	if numCacheEntries, err := strconv.Atoi(os.Getenv("NUMBER_CACHE_ENTRIES")); err != nil {
		return err
	} else {
		newEnv.NUMBER_CACHE_ENTRIES = numCacheEntries
	}
	// TODO: Add some validations to make sure that essential envs were loaded
	loadedEnv = newEnv
	return nil
}
