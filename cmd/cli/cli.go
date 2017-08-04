package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/weiwei04/compass/pkg/clicommands"
)

const (
	usage = `
Compass(github.com/weiwei04/compass) is a front end for tiller, it implements all tiller grpc api and integrate with Helm-Registry(github.com/caicloud/helm-registry) and provide a extra grpc api: InstallCompassRelease, it will download chart from helm-registry and send a InstallReleaseRequest to tiller

Avaiable Commands:
  compass install [CHART] [OPTIONS]
  `
)

func main() {
	app := cli.NewApp()
	app.Name = "compass"
	app.Usage = usage
	app.Commands = []cli.Command{
		clicommands.InstallReleaseCommand,
		clicommands.UpdateReleaseCommand,
		clicommands.UpgradeReleaseCommand,
	}
	app.Run(os.Args)
}
