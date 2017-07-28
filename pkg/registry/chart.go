package registry

import (
	"fmt"
	"github.com/caicloud/helm-registry/pkg/rest/v1"
	"github.com/urfave/cli"
	"strings"
	"io/ioutil"
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
		space := ctx.Args().First()
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}

		parts := strings.Split(space, "/")
		if len(parts) == 1 {
			res, err := c.ListCharts(parts[0], 0, 1000)
			if err != nil {
				return err
			}
			fmt.Println("charts in space", parts[0])
			for _, item := range res.Items {
				fmt.Println("\t\t", item)
			}
			fmt.Println("---------------------------------------")
		} else if len(parts) == 2 {
			res, err := c.ListVersions(parts[0], parts[1], 0, 1000)
			if err != nil {
				return err
			}
			fmt.Println("chart", parts[1], "in space", parts[0])
			for _, item := range res.Items {
				fmt.Println("\t\t", item)
			}
			fmt.Println("---------------------------------------")
		} else {
			return fmt.Errorf("invalid name `%s` must be `space` or `space/chartname`", space)
		}

		return nil
	},
}

var pushChartCommand = cli.Command{
	Name:  "push",
	Usage: `push space chart.tgz`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 2, exactArgs); err != nil {
			return err
		}
		chart := ctx.Args().Get(1)
		chartPkg, err := ioutil.ReadFile(chart)
		if err != nil {
			return err
		}
		space := ctx.Args().First()
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}
		res, err := c.UploadChart(space, chartPkg)
		if err != nil {
			return err
		}
		fmt.Println("push chart", res.Chart, res.Link)
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
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err

		}
		name := ctx.Args().First()
		parts := strings.Split(name, "/")

		if len(parts) != 2 {
			return fmt.Errorf("invalid name `%s` must be `space` or `space/chartname:version`", name)
		}
		space := parts[0]
		chart := parts[1]
		parts = strings.Split(chart, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid name `%s` must be `space` or `space/chartname:version`", name)
		}
		chartName := parts[0]
		chartVer := parts[1]
		res, err := c.FetchVersionMetadata(space, chartName, chartVer)
		if err != nil {
			return err
		}
		fmt.Println(name, "metadata")
		fmt.Println(res.String())
		return nil
	},
}
