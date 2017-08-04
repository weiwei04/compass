package clicommands

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	capi "github.com/weiwei04/compass/pkg/api/client"
	hapi "k8s.io/helm/pkg/proto/hapi/chart"
)

var UpgradeReleaseCommand = cli.Command{
	Name:  "upgrade",
	Usage: "",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "-name RELEASE_NAME",
		},
		cli.StringFlag{
			Name:  "values",
			Value: "",
			Usage: "-values values.yaml",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires chart pkg as arguments",
				os.Args[0], ctx.Command.Name)
		}
		chart := ctx.Args().First()
		releaseName := ctx.String("name")
		if releaseName == "" {
			return fmt.Errorf("%s: %q must provide release name with --name",
				os.Args[0], ctx.Command.Name)
		}
		var values *hapi.Config
		valuesFileName := ctx.String("values")
		if valuesFileName != "" {
			data, err := ioutil.ReadFile(valuesFileName)
			if err != nil {
				return err
			}
			values = &hapi.Config{Raw: string(data)}
		}

		client := defaultClient()
		err := client.Connect()
		if err != nil {
			return err
		}
		defer client.Shutdown()

		req := &capi.UpgradeReleaseRequest{
			Name:   releaseName,
			Chart:  chart,
			Values: values,
		}
		resp, err := client.UpgradeRelease(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(resp.Release.String())
		return nil
	},
}
