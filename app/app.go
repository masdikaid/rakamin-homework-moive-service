package app

import (
	"github.com/masdikaid-rakamin-homework-movie-service/config"
)

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
