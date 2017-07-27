package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/weiwei04/compass/pkg/registry"
)

const (
	usage = `
Helm-Registry(github.com/caicloud/helm-registry) integration with Helm.

This provides tools for working with the helm-registry.

Available Commands:
  space list|create|delete   List/Create/Delete a Namespace
  chart list|push|delete|inspect    List/Push/Inspect a Chart`
)

func main() {
	app := cli.NewApp()
	app.Name = "helm-registry-plugin"
	app.Usage = usage
	app.Commands = []cli.Command{
		registry.SpaceCommand,
		registry.ChartCommand,
	}
	app.Run(os.Args)
}
