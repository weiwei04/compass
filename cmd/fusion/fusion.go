package main

import (
	"flag"
	"github.com/urfave/cli"
	"github.com/weiwei04/compass/pkg/clicommands"
	"os"
)

const (
	usage = `a cli tool for Compass,
	 Compass(github.com/weiwei04/compass) is a front end for tiller,
	 it implements all tiller grpc api and integrate with Helm-Registry(github.com/caicloud/helm-registry)
`
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	app := cli.NewApp()
	app.Name = "fusion"
	app.Version = "0.0.1"
	app.Usage = usage
	app.Commands = []cli.Command{
		clicommands.SpaceCommand,
		clicommands.ChartCommand,
		clicommands.ReleaseCommand,
	}
	app.Run(os.Args)
}
