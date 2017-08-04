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

var UpdateReleaseCommand = cli.Command{
	Name:  "update",
	Usage: "",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "values",
			Value: "",
			Usage: "-values values.yaml",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires release name as arguments",
				os.Args[0], ctx.Command.Name)
		}
		releaseName := ctx.Args().First()
		var values *hapi.Config
		valuesFileName := ctx.String("values")
		if valuesFileName == "" {
			return fmt.Errorf("%s: %q must provide values with --values",
				os.Args[0], ctx.Command.Name)
		}
		data, err := ioutil.ReadFile(valuesFileName)
		if err != nil {
			return err
		}
		values = &hapi.Config{Raw: string(data)}

		client := defaultClient()
		err = client.Connect()
		if err != nil {
			return err
		}
		defer client.Shutdown()

		req := &capi.UpdateReleaseRequest{
			Name:   releaseName,
			Values: values,
		}
		resp, err := client.UpdateRelease(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(resp.Release.String())
		return nil
	},
}
