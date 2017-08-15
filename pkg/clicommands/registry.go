package clicommands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/urfave/cli"
	api "github.com/weiwei04/compass/pkg/api/client"
	yaml "gopkg.in/yaml.v2"
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
		c, err := defaultRegistryClient()
		resp, err := c.ListSpaces(nil, &api.ListSpacesRequest{
			Limit: 10000,
		})
		if err != nil {
			return err
		}
		fmt.Println("spaces:")
		for _, space := range resp.Spaces {
			fmt.Println("\t\t", space)
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
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		resp, err := c.CreateSpace(nil, &api.CreateSpaceRequest{Space: ns})
		if err != nil {
			return err
		}
		fmt.Println("created space", resp.Space, resp.Link)
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
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		_, err = c.DeleteSpace(nil, &api.DeleteSpaceRequest{Space: ns})
		return err
	},
}

type chartValues map[string]interface{}

var ChartCommand = cli.Command{
	Name:  "chart",
	Usage: "",
	Subcommands: []cli.Command{
		listChartCommand,
		pushChartCommand,
		inspectChartCommand,
		showChartFileCommand,
		fetchChartCommand,
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
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}

		space, name, err := splitSpaceChart(arg)
		if err != nil {
			return err
		}
		if name == "" {
			resp, err := c.ListCharts(nil, &api.ListChartsRequest{
				Space: space,
				Limit: 1000})
			if err != nil {
				return err
			}
			fmt.Println(arg)
			for _, chart := range resp.Charts {
				fmt.Println("\t\t", chart)
			}
			fmt.Println("-----------------------------------------")
		} else {
			resp, err := c.ListChartVersions(nil, &api.ListChartVersionsRequest{
				Space: space,
				Chart: name,
				Limit: 1000,
			})
			if err != nil {
				return err
			}
			fmt.Println(arg)
			for _, ver := range resp.Versions {
				fmt.Println("\t\t", ver)
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
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		resp, err := c.PushChart(nil, &api.PushChartRequest{
			Space: space,
			Data:  chartPkg,
		})
		if err != nil {
			return err
		}
		fmt.Println(resp.Chart, "at", resp.Link)
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
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		arg := ctx.Args().First()
		space, name, ver, err := splitSpaceChartVer(arg)
		if err != nil {
			return err
		}
		resp, err := c.GetChartMetadata(nil, &api.GetChartMetadataRequest{
			Space:   space,
			Chart:   name,
			Version: ver,
		})
		if err != nil {
			return err
		}
		var data []byte
		switch ctx.String("output") {
		case "yaml":
			data, _ = yaml.Marshal(resp)
		case "json":
			data, _ = json.MarshalIndent(resp, "", "  ")
		default:
			return fmt.Errorf("unsupported output format %s", ctx.String("output"))
		}
		fmt.Println(arg, "metadata")
		fmt.Println(string(data))
		fmt.Println("-----------------------------------------")
		return nil
	},
}

func showChartValues(client api.Registry, space, chart, version string) ([]byte, error) {
	resp, err := client.GetChartValues(nil, &api.GetChartValuesRequest{
		Space:   space,
		Chart:   chart,
		Version: version,
	})
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(resp.Values)
}

func showChartReadme(client api.Registry, space, chart, version string) ([]byte, error) {
	resp, err := client.GetChartReadme(nil, &api.GetChartReadmeRequest{
		Space:   space,
		Chart:   chart,
		Version: version,
	})
	if err != nil {
		return nil, err
	}
	return resp.Readme, err
}

func showChartDeps(client api.Registry, space, chart, version string) ([]byte, error) {
	resp, err := client.GetChartRequirements(nil, &api.GetChartRequirementsRequest{
		Space:   space,
		Chart:   chart,
		Version: version,
	})
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(resp.Dependencies)
}

var showChartFileCommand = cli.Command{
	Name: "show",
	Usage: `show space/chart:ver
	  for example:
	    values myspace/mychart:0.0.1 (this will get myspace/mychart:0.0.1 values.yaml content)`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "readme",
			Usage: `-f readme|values|deps`,
		},
	},
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		arg := ctx.Args().First()
		space, name, ver, err := splitSpaceChartVer(arg)
		if err != nil {
			return err
		}
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		var data []byte
		switch ctx.String("file") {
		case "readme":
			data, err = showChartReadme(c, space, name, ver)
		case "values":
			data, err = showChartValues(c, space, name, ver)
		case "deps":
			data, err = showChartDeps(c, space, name, ver)
		default:
			err = fmt.Errorf("unsupported file %s", ctx.String("file"))
		}
		if err != nil {
			return err
		}
		fmt.Println(arg, " ", ctx.String("file"))
		fmt.Println()
		fmt.Println(string(data))
		fmt.Println("-----------------------------------------")
		return nil
	},
}

var fetchChartCommand = cli.Command{
	Name: "fetch",
	Usage: `fetch space/chart:ver
    for example:
    fetch myspace/mychart:0.0.1 (this will download myspace/mychart:0.0.1 chart package)`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "./",
			Usage: `-o dir`,
		},
	},
	Action: func(ctx *cli.Context) error {
		if err := checkArgs(ctx, 1, exactArgs); err != nil {
			return err
		}
		c, err := defaultRegistryClient()
		if err != nil {
			return err
		}
		arg := ctx.Args().First()
		space, name, ver, err := splitSpaceChartVer(arg)
		if err != nil {
			return err
		}
		resp, err := c.FetchChart(nil, &api.FetchChartRequest{
			Space:   space,
			Chart:   name,
			Version: ver,
		})
		if err != nil {
			return err
		}
		tarName := name + "-" + ver + ".tgz"
		dir := ctx.String("output")
		fileName := filepath.Join(dir, tarName)
		return ioutil.WriteFile(fileName, resp.Data, 0644)
	},
}
