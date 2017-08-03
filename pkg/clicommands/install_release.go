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

var ReleaseCommand = cli.Command{
	Name:  "install",
	Usage: "",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "-name RELEASENAME",
		},
		cli.StringFlag{
			Name:  "namespace",
			Value: "default",
			Usage: "-namespace NAMESPACE",
		},
		cli.StringFlag{
			Name:  "values",
			Value: "",
			Usage: "-values values.yaml",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires space/chart:ver as arguments",
				os.Args[0], ctx.Command.Name)
		}
		chart := ctx.Args().First()
		releaseName := ctx.String("name")
		if releaseName == "" {
			return fmt.Errorf("%s: %q must provide release name with --name",
				os.Args[0], ctx.Command.Name)
		}

		namespace := ctx.String("namespace")

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

		req := &capi.CreateReleaseRequest{
			Chart:     chart,
			Name:      releaseName,
			Namespace: namespace,
			Values:    values,
		}
		resp, err := client.CreateRelease(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(resp.Release.String())
		return nil
	},
}

var defaultClient = func() capi.CompassClient {
	compassAddr := os.Getenv("COMPASS_ADDR")
	if compassAddr == "" {
		compassAddr = "http://127.0.0.1:8900"
	}

	return capi.NewCompassClient(compassAddr)
}
