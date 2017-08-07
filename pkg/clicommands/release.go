package clicommands

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	api "github.com/weiwei04/compass/pkg/api/client"
	hapi "k8s.io/helm/pkg/proto/hapi/chart"
)

var ReleaseCommand = cli.Command{
	Name:  "release",
	Usage: "",
	Subcommands: []cli.Command{
		installReleaseCommand,
		updateReleaseCommand,
		upgradeReleaseCommand,
	},
}

var installReleaseCommand = cli.Command{
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

		client, err := defaultReleaseClient()
		if err != nil {
			return err
		}
		defer client.Shutdown()

		req := &api.CreateReleaseRequest{
			Chart:     chart,
			Name:      releaseName,
			Namespace: namespace,
			Values:    values,
		}
		_, err = client.CreateRelease(context.Background(), req)
		if err != nil {
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("deployed")
		return nil
	},
}

var updateReleaseCommand = cli.Command{
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

		client, err := defaultReleaseClient()
		if err != nil {
			return err
		}
		defer client.Shutdown()

		req := &api.UpdateReleaseRequest{
			Name:   releaseName,
			Values: values,
		}
		_, err = client.UpdateRelease(context.Background(), req)
		if err != nil {
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("update release success")
		return nil
	},
}

var upgradeReleaseCommand = cli.Command{
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

		client, err := defaultReleaseClient()
		if err != nil {
			return err
		}
		defer client.Shutdown()

		req := &api.UpgradeReleaseRequest{
			Name:   releaseName,
			Chart:  chart,
			Values: values,
		}
		_, err = client.UpgradeRelease(context.Background(), req)
		if err != nil {
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("upgrade release success")
		return nil
	},
}
