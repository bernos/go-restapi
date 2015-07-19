package config

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacobstr/confer"
	"github.com/unrolled/render"

	log "github.com/Sirupsen/logrus"
)

type Configuration struct {
	*confer.Config
	Render *render.Render
	Router *mux.Router
}

func NewConfiguration() *Configuration {
	return &Configuration{Config: loadConfig()}
}

func loadConfig() *confer.Config {
	config := confer.NewConfig()
	appenv := os.Getenv("GO_APPENV")
	paths := []string{"application.yml"}

	if appenv != "" {
		paths = append(paths, fmt.Sprintf("application.%s.yml", appenv))
	}

	if err := config.ReadPaths(paths...); err != nil {
		log.Warn(err)
	}

	return config
}
