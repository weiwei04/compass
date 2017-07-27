package registry

import (
	"fmt"

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
		fmt.Println("list namepsace")
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
		fmt.Println("create namespace ", ns)
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
		fmt.Println("delete namespace ", ns)
		return nil
	},
}
