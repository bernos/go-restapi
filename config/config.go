package config

import (
	"fmt"
	"github.com/jacobstr/confer"
	"github.com/unrolled/render"
	"log"
	"os"
)

type ApiConfiguration struct {
	*confer.Config
	Render *render.Render
}

func NewApiConfiguration() *ApiConfiguration {
	config := &ApiConfiguration{Config: confer.NewConfig()}
	appenv := os.Getenv("GO_APPENV")
	paths := []string{"application.yml"}

	if appenv != "" {
		paths = append(paths, fmt.Sprintf("application.%s.yml", appenv))
	}

	if err := config.ReadPaths(paths...); err != nil {
		log.Print(err)
	}

	return config
}
