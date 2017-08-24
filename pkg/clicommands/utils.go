package clicommands

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	api "github.com/weiwei04/compass/pkg/api/client"
)

func defaultRegistryClient() (api.Registry, error) {
	registryAddr := os.Getenv("HELM_REGISTRY_ADDR")
	if registryAddr == "" {
		registryAddr = "http://192.168.99.100:32588"
	}
	client := api.NewHelmRegistryClient(registryAddr)
	return client, nil
}

func defaultReleaseClient() (api.Release, error) {
	compassAddr := os.Getenv("COMPASS_ADDR")
	if compassAddr == "" {
		compassAddr = "192.168.99.100:32589"
	}
	client := api.NewReleaseClient(compassAddr)
	return client, nil
}

func splitSpaceChart(arg string) (string, string, error) {
	parts := strings.Split(arg, "/")
	if len(parts) == 1 {
		return parts[0], "", nil
	} else if len(parts) == 2 {
		return parts[0], parts[1], nil
	}
	return "", "", fmt.Errorf("invalid arg `%s` must be `space[/chart]`", arg)
}

func splitSpaceChartVer(arg string) (string, string, string, error) {
	parts := strings.Split(arg, "/")

	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid arg `%s` must be `space/chart:ver`", arg)
	}
	space := parts[0]
	chart := parts[1]
	parts = strings.Split(chart, ":")
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid arg `%s` must be `space/chartver`", arg)
	}
	name := parts[0]
	ver := parts[1]
	return space, name, ver, nil
}

const (
	exactArgs = iota
	minArgs
	maxArgs
)

func checkArgs(context *cli.Context, expected, checkType int) error {
	var err error
	cmdName := context.Command.Name
	switch checkType {
	case exactArgs:
		if context.NArg() != expected {
			err = fmt.Errorf("%s: %q requires exactly %d argument(s)", os.Args[0], cmdName, expected)
		}
	case minArgs:
		if context.NArg() < expected {
			err = fmt.Errorf("%s: %q requires a minimum of %d argument(s)", os.Args[0], cmdName, expected)
		}
	case maxArgs:
		if context.NArg() > expected {
			err = fmt.Errorf("%s: %q requires a maximum of %d argument(s)", os.Args[0], cmdName, expected)
		}
	}

	if err != nil {
		fmt.Printf("Incorrect Usage.\n\n")
		cli.ShowCommandHelp(context, cmdName)
		return err
	}
	return nil
}
