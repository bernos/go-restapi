package todos

import (
	"github.com/bernos/go-restapi/config"
)

func Configure(c *config.Configuration) {
	ConfigureRoutes(c)
}
