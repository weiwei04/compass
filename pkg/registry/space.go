package registry

import (
	"fmt"
	"github.com/caicloud/helm-registry/pkg/rest/v1"
	"github.com/urfave/cli"
)

var SpaceCommand = cli.Command{
	Name:  "space",
	Usage: "",
	Subcommands: []cli.Command{
		listSpaceCommand,
		createSpaceCommand,
		deleteSpaceCommand,
	},
}

var listSpaceCommand = cli.Command{
	Name:  "list",
	Usage: "",
	Action: func(ctx *cli.Context) error {
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}
		res, err := c.ListSpaces(0, 10000)
		if err != nil {
			return err
		}
		fmt.Println("spaces:")
		for _, item := range res.Items {
			fmt.Println("\t\t", item)
		}
		fmt.Println("-----------------------------------------")
		return nil
	},
}

var createSpaceCommand = cli.Command{
	Name:      "create",
	Usage:     "",
	ArgsUsage: `<namespace>`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		ns := ctx.Args().First()
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}
		res, err := c.CreateSpace(ns)
		if err != nil {
			return err
		}
		fmt.Println("created space", res.Name, res.Link)
		return nil
	},
}

var deleteSpaceCommand = cli.Command{
	Name:  "delete",
	Usage: "",
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		ns := ctx.Args().First()
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}
		return c.DeleteSpace(ns)
	},
}
