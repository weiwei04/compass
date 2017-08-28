package clicommands

import (
	"github.com/golang/glog"
)

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/urfave/cli"
	api "github.com/weiwei04/compass/pkg/api/client"
	hapi "k8s.io/helm/pkg/proto/hapi/chart"
)

var ReleaseCommand = cli.Command{
	Name:  "release",
	Usage: "",
	Subcommands: []cli.Command{
		installReleaseCommand,
		listReleasesCommand,
		listReleaseHistoryCommand,
		getReleaseStatusCommand,
		getReleaseContentCommand,
		updateReleaseCommand,
		upgradeReleaseCommand,
	},
}

func printYaml(data interface{}) error {
	raw_data, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(raw_data))
	return nil
}

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
		namespace := ctx.String("namespace")
		var values *hapi.Config
		valuesFileName := ctx.String("values")
		if valuesFileName != "" {
			data, err := ioutil.ReadFile(valuesFileName)
			if err != nil {
				glog.Errorf("install %s failed, err:%s", releaseName, err)
				return err
			}
			values = &hapi.Config{Raw: string(data)}
		}

		client, _ := defaultReleaseClient()

		req := &api.CreateReleaseRequest{
			Chart:     chart,
			Name:      releaseName,
			Namespace: namespace,
			Values:    values,
		}
		_, err := client.CreateRelease(context.Background(), req)
		if err != nil {
			glog.Errorf("CreateRelease(name:%s, namespace:%s) failed, err:%s",
				req.Name, req.Namespace, err)
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("deployed")
		return nil
	},
}

var listReleasesCommand = cli.Command{
	Name:  "list",
	Usage: "list [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
	},
	Action: func(ctx *cli.Context) error {
		client, _ := defaultReleaseClient()
		req := &api.ListReleasesRequest{
			Limit:     100,
			Namespace: ctx.String("namespace"),
		}
		resp, err := client.ListReleases(context.Background(), req)
		if err != nil {
			glog.Errorf("ListReleases(namespace:%s) failed, err:%s",
				req.Namespace, err)
			return err
		}
		for _, release := range resp.Releases {
			fmt.Println("release:", release.Name, ", chart:", release.Chart.Metadata.Name)
		}
		return nil
	},
}

var listReleaseHistoryCommand = cli.Command{
	Name:  "history",
	Usage: "history RELEASE_NAME [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires release name as arguments",
				os.Args[0], ctx.Command.Name)
		}
		releaseName := ctx.Args().First()
		client, _ := defaultReleaseClient()
		req := &api.GetReleaseHistoryRequest{
			Name:      releaseName,
			Namespace: ctx.String("namespace"),
			Max:       10,
		}
		resp, err := client.GetReleaseHistory(context.Background(), req)
		if err != nil {
			glog.Errorf("GetReleaseHistory(name:%s, namespace:%s) failed, err:%s",
				req.Name, req.Namespace, err)
			return err
		}
		for _, release := range resp.Releases {
			fmt.Println("version:", release.Chart.Metadata.Version, ", chart:", release.Chart.Metadata.Name)
		}
		return nil
	},
}

var getReleaseStatusCommand = cli.Command{
	Name:  "status",
	Usage: "status RELEASE_NAME [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
		cli.IntFlag{
			Name:  "version, v",
			Value: 0,
			Usage: "-v version",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires release_name as arguments",
				os.Args[0], ctx.Command.Name)
		}
		releaseName := ctx.Args().First()
		namespace := ctx.String("namespace")
		client, _ := defaultReleaseClient()
		req := &api.GetReleaseStatusRequest{
			Name:      releaseName,
			Namespace: namespace,
			Version:   int32(ctx.Int("version")),
		}
		resp, err := client.GetReleaseStatus(context.Background(), req)
		if err != nil {
			glog.Errorf("GetReleaseStatus(name:%s, namespace:%s, version:%d) failed, err:%s",
				req.Name, req.Namespace, req.Version, err)
			return err
		}
		return printYaml(resp.Info)
	},
}

var getReleaseContentCommand = cli.Command{
	Name:  "content",
	Usage: "content RELEASE_NAME [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
		cli.IntFlag{
			Name:  "version, v",
			Value: 0,
			Usage: "-v version",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires release_name as arguments",
				os.Args[0], ctx.Command.Name)
		}
		releaseName := ctx.Args().First()
		namespace := ctx.String("namespace")
		client, _ := defaultReleaseClient()
		req := &api.GetReleaseContentRequest{
			Name:      releaseName,
			Namespace: namespace,
			Version:   int32(ctx.Int("version")),
		}
		resp, err := client.GetReleaseContent(context.Background(), req)
		if err != nil {
			glog.Errorf("GetReleaseContent(name:%s, namespace:%s, version:%d) failed, err:%s",
				req.Name, req.Namespace, req.Version)
			return err
		}
		return printYaml(resp.Release.Info)
	},
}

var updateReleaseCommand = cli.Command{
	Name:  "update",
	Usage: "update RELEASE_NAME [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "values",
			Value: "",
			Usage: "-values values.yaml",
		},
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires release name as arguments",
				os.Args[0], ctx.Command.Name)
		}
		releaseName := ctx.Args().First()
		var values *hapi.Config
		valuesFileName := ctx.String("values")
		if valuesFileName == "" {
			return fmt.Errorf("%s: %q must provide values with --values",
				os.Args[0], ctx.Command.Name)
		}
		data, err := ioutil.ReadFile(valuesFileName)
		if err != nil {
			glog.Errorf("update cmd read values file failed, err:%s", err)
			return err
		}
		values = &hapi.Config{Raw: string(data)}

		client, _ := defaultReleaseClient()
		req := &api.UpdateReleaseRequest{
			Name:      releaseName,
			Namespace: ctx.String("namespace"),
			Values:    values,
		}
		_, err = client.UpdateRelease(context.Background(), req)
		if err != nil {
			glog.Errorf("UpdateRelease(name:%s, namespace:%s) failed, err:%s",
				req.Name, req.Namespace, err)
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("update release success")
		return nil
	},
}

var upgradeReleaseCommand = cli.Command{
	Name:  "upgrade",
	Usage: "upgrade CHART_NAME [OPTIONS]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "-name RELEASE_NAME",
		},
		cli.StringFlag{
			Name:  "values",
			Value: "",
			Usage: "-values values.yaml",
		},
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "-n namespace",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() != 1 {
			return fmt.Errorf("%s: %q requires chart pkg as arguments",
				os.Args[0], ctx.Command.Name)
		}
		chart := ctx.Args().First()
		releaseName := ctx.String("name")
		if releaseName == "" {
			return fmt.Errorf("%s: %q must provide release name with --name",
				os.Args[0], ctx.Command.Name)
		}
		var values *hapi.Config
		valuesFileName := ctx.String("values")
		if valuesFileName != "" {
			data, err := ioutil.ReadFile(valuesFileName)
			if err != nil {
				glog.Errorf("upgrade cmd read values file failed, err:%s", err)
				return err
			}
			values = &hapi.Config{Raw: string(data)}
		}

		client, _ := defaultReleaseClient()
		req := &api.UpgradeReleaseRequest{
			Name:      releaseName,
			Namespace: ctx.String("namespace"),
			Chart:     chart,
			Values:    values,
		}
		_, err := client.UpgradeRelease(context.Background(), req)
		if err != nil {
			glog.Errorf("UpgradeRelease(name:%s, namespace:%s) failed, err:%s",
				req.Name, req.Namespace, err)
			return err
		}
		// TODO: prettyPrint resp
		fmt.Println("upgrade release success")
		return nil
	},
}
