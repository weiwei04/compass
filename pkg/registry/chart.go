package registry

import (
	"fmt"
	"github.com/caicloud/helm-registry/pkg/rest/v1"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	//hapi "k8s.io/helm/pkg/proto/hapi/chart"
	//"k8s.io/helm/pkg/chartutil"
	//"encoding/json"
	"encoding/json"
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
			fmt.Println("-----------------------------------------")
		} else if len(parts) == 2 {
			res, err := c.ListVersions(parts[0], parts[1], 0, 1000)
			if err != nil {
				return err
			}
			fmt.Println("chart", parts[1], "in space", parts[0])
			for _, item := range res.Items {
				fmt.Println("\t\t", item)
			}
			fmt.Println("-----------------------------------------")
		} else {
			return fmt.Errorf("invalid name `%s` must be `space` or `space/chartname`", space)
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
	Name: "inspect",
	Usage: `inspect space/chart:ver
	  for example:
	    inspect myspace/mychart:0.0.1 (this will get myspace/mychart:0.0.1 metadata)`,
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
		c, err := v1.NewClient("http://helm-registry.ke-cs.dev.qiniu.io")
		if err != nil {
			return err
		}
		name := ctx.Args().First()
		parts := strings.Split(name, "/")
		if len(parts) != 2 {
			return fmt.Errorf("invalid name `%s` must be `space/chart:version`", name)
		}
		space := parts[0]
		chart := parts[1]
		parts = strings.Split(chart, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid name `%s` must be `space/chart:version`", name)
		}
		chartName := parts[0]
		chartVer := parts[1]
		raw, err := c.FetchVersionValues(space, chartName, chartVer)
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

		fmt.Println("chart", name, " values")
		fmt.Println()
		fmt.Println(string(data))
		fmt.Println("------------------------------")
		return nil
	},
}
