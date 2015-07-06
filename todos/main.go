package todos

import (
	"github.com/bernos/go-restapi/config"
)

func Configure(c *config.ApiConfiguration) {
	ConfigureRoutes(c)
}
