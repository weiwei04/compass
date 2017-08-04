package clicommands

import (
	"os"

	capi "github.com/weiwei04/compass/pkg/api/client"
)

var defaultClient = func() capi.CompassClient {
	compassAddr := os.Getenv("COMPASS_ADDR")
	if compassAddr == "" {
		compassAddr = "http://127.0.0.1:8900"
	}

	return capi.NewCompassClient(compassAddr)
}
