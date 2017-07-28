package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type chartValues map[string]interface{}

var ChartCommand = cli.Command{
	Name:  "chart",
	Usage: "",
	Subcommands: []cli.Command{
		listChartCommand,
		pushChartCommand,
		inspectChartCommand,
		getChartFileCommand,
	},
}

var listChartCommand = cli.Command{
	Name: "list",
	Usage: `list space[/chart]
	  for example:
	    list myspace (this will list chart under space:myspace)
	    list myspace/mychart (this will list versions of myspace/mychart)`,
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		arg := ctx.Args().First()
		c, err := defaultClient()
		if err != nil {
			return err
		}

		space, name, err := splitSpaceChart(arg)
		if err != nil {
			return err
		}
		if name == "" {
			res, err := c.ListCharts(space, 0, 1000)
			if err != nil {
				return err
			}
			fmt.Println(arg)
			for _, item := range res.Items {
				fmt.Println("\t\t", item)
			}
			fmt.Println("-----------------------------------------")
		} else {
			res, err := c.ListVersions(space, name, 0, 1000)
			if err != nil {
				return err
			}
			fmt.Println(arg)
			for _, item := range res.Items {
				fmt.Println("\t\t", item)
			}
			fmt.Println("-----------------------------------------")
		}
		return nil
	},
}

var pushChartCommand = cli.Command{
	Name: "push",
	Usage: `push space chart
	  for example:
	    push myspace mychart.tgz (this will push mychart.tgz under mysapce)`,
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
		c, err := defaultClient()
		if err != nil {
			return err
		}
		res, err := c.UploadChart(space, chartPkg)
		if err != nil {
			return err
		}
		fmt.Println(res.Chart, "at", res.Link)
		return nil
	},
}

var inspectChartCommand = cli.Command{
	Name: "inspect",
	Usage: `inspect space/chart:ver
	  for example:
	    inspect myspace/mychart:0.0.1 (this will get myspace/mychart:0.0.1 metadata)`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "yaml",
			Usage: "-o yaml|json",
		},
	},
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		c, err := defaultClient()
		if err != nil {
			return err
		}
		arg := ctx.Args().First()
		space, name, ver, err := splitSpaceChartVer(arg)
		if err != nil {
			return err
		}
		meta, err := c.FetchVersionMetadata(space, name, ver)
		if err != nil {
			return err
		}
		var data []byte
		switch ctx.String("output") {
		case "yaml":
			data, _ = yaml.Marshal(meta)
		case "json":
			data, _ = json.MarshalIndent(meta, "", "  ")
		default:
			return fmt.Errorf("unsupported output format %s", ctx.String("output"))
		}
		fmt.Println(arg, "metadata")
		fmt.Println(string(data))
		fmt.Println("-----------------------------------------")
		return nil
	},
}

var getChartFileCommand = cli.Command{
	Name: "values",
	Usage: `values space/chart:ver
	  for example:
	    values myspace/mychart:0.0.1 (this will get myspace/mychart:0.0.1 values.yaml content)`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "yaml",
			Usage: `-o yaml|json`,
		},
	},
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		c, err := defaultClient()
		if err != nil {
			return err
		}
		arg := ctx.Args().First()
		space, name, ver, err := splitSpaceChartVer(arg)
		if err != nil {
			return err
		}
		raw, err := c.FetchVersionValues(space, name, ver)
		if err != nil {
			return err
		}

		var data []byte
		switch ctx.String("output") {
		case "yaml":
			values := chartValues{}
			yaml.Unmarshal(raw, &values)
			data, _ = yaml.Marshal(values)
		case "json":
			values := chartValues{}
			json.Unmarshal(raw, &values)
			data, _ = json.MarshalIndent(values, "", "  ")
		default:
			return fmt.Errorf("unsupported output format %s", ctx.String("output"))
		}

		fmt.Println(arg, " values")
		fmt.Println()
		fmt.Println(string(data))
		fmt.Println("-----------------------------------------")
		return nil
	},
}
