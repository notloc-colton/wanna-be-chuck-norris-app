package environment

var loadedEnv env

const (
	localEnvironment = "LOCAL"
	prodEnvironment  = "PROD"
)

type env struct {
	APP_ENVIRONMENT      string
	JOKE_URL             string
	NUMBER_CACHE_ENTRIES int
	RANDOM_NAME_URL      string
	SERVER_ADDRESS       string
	SERVER_PORT          string
}

func ENV() env {
	return loadedEnv
}
