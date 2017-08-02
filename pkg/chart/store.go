package chart

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	v1 "github.com/caicloud/helm-registry/pkg/rest/v1"
	"k8s.io/helm/pkg/chartutil"
	hapi "k8s.io/helm/pkg/proto/hapi/chart"
)

var (
	ErrNotFound = errors.New("not found")
)

type Store interface {
	Get(chart string) (*hapi.Chart, error)
}

var _ Store = helmRegistryStore{}

type helmRegistryStore struct {
	client *v1.Client
}

func NewHelmRegistryStore(addr string) (Store, error) {
	client, err := v1.NewClient(addr)
	if err != nil {
		return nil, err
	}
	return helmRegistryStore{client}, nil
}

func splitSpaceChartVer(arg string) (string, string, string, error) {
	parts := strings.Split(arg, "/")
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid chart `%s` must be `space/chart:ver`", arg)
	}
	space := parts[0]
	chart := parts[1]
	parts = strings.Split(chart, ":")
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid chart `%s` must be `space/chartver`", arg)
	}
	name := parts[0]
	ver := parts[1]
	return space, name, ver, nil
}

func (r helmRegistryStore) Get(fullname string) (*hapi.Chart, error) {
	space, chart, ver, err := splitSpaceChartVer(fullname)
	if err != nil {
		return nil, err
	}
	data, err := r.client.DownloadVersion(space, chart, ver)
	if err != nil {
		return nil, err
	}
	return chartutil.LoadArchive(bytes.NewReader(data))
}
