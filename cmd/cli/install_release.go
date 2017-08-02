package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	"github.com/weiwei04/compass/pkg/api/services/compass"
	"google.golang.org/grpc"
	"io/ioutil"
	hapi "k8s.io/helm/pkg/proto/hapi/chart"
	"os"
	"time"
)

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
		if releaseName == "" {
			return fmt.Errorf("%s: %q must provide release name with --name",
				os.Args[0], ctx.Command.Name)
		}

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

		client, conn, err := defaultClient()
		if err != nil {
			return err
		}
		defer conn.Close()

		req := &compass.CreateCompassReleaseRequest{
			Chart:     chart,
			Name:      releaseName,
			Namespace: namespace,
			Values:    values,
		}
		resp, err := client.CreateCompassRelease(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(resp.Release.String())
		return nil
	},
}

var defaultClient = func() (compass.CompassServiceClient, *grpc.ClientConn, error) {
	compassAddr := os.Getenv("COMPASS_ADDR")
	if compassAddr == "" {
		compassAddr = "http://127.0.0.1:8900"
	}
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(compassAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	client := compass.NewCompassServiceClient(conn)
	return client, conn, nil
}
