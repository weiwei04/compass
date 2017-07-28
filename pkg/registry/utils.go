package registry

import (
	"fmt"
	"os"
	"strings"

	"github.com/caicloud/helm-registry/pkg/rest/v1"
	"github.com/urfave/cli"
)

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

var registryAddr string

func init() {
	registryAddr = os.Getenv("HELM_REGISTRY_ADDR")
	if registryAddr == "" {
		registryAddr = "http://localhost:8900"
	}
}

func defaultClient() (*v1.Client, error) {
	return v1.NewClient(registryAddr)
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
