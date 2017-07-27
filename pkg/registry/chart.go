package registry

import (
	"fmt"

	"github.com/urfave/cli"
)

var ChartCommand = cli.Command{
	Name:  "chart",
	Usage: "",
	Subcommands: []cli.Command{
		listChartCommand,
		pushChartCommand,
		inspectChartCommand,
	},
}

var listChartCommand = cli.Command{
	Name: "list",
	Usage: `list namespace for example list library
  list namespace/chartname for example list library/redis`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		name := ctx.Args().First()
		fmt.Println("list ", name)
		return nil
	},
}

var pushChartCommand = cli.Command{
	Name:  "push",
	Usage: `push chart.tgz`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		name := ctx.Args().First()
		fmt.Println("push ", name)
		return nil
	},
}

var inspectChartCommand = cli.Command{
	Name:  "inspect",
	Usage: `inspect namespace/chartname:version for example library/redis:0.1.0`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		name := ctx.Args().First()
		fmt.Println("inspect ", name)
		return nil
	},
}
