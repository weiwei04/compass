package client

import (
	"bytes"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"

	v1 "github.com/caicloud/helm-registry/pkg/rest/v1"
	"golang.org/x/net/context"
	yaml "gopkg.in/yaml.v2"
)

const (
	readmeFile       = "README.md"
	requirementsFile = "requirements.yaml"
)

func NewHelmRegistryClient(addr string) Registry {
	return &helmRegistry{addr: addr}
}

type helmRegistry struct {
	addr   string
	client *v1.Client
}

func (r *helmRegistry) Connect() error {
	var err error
	r.client, err = v1.NewClient(r.addr)
	return err
}

func (r *helmRegistry) Shutdown() {}

func (r *helmRegistry) ListSpaces(ctx context.Context, req *ListSpacesRequest) (*ListSpacesResponse, error) {
	listResp, err := r.client.ListSpaces(req.Start, req.Limit)
	if err != nil {
		return nil, err
	}
	return &ListSpacesResponse{
		Spaces: listResp.Items,
		IsEnd:  req.Start+len(listResp.Items) >= listResp.Metadata.Total,
		offset: req.Start,
		limit:  req.Limit,
	}, nil
}

func (r *helmRegistry) CreateSpace(ctx context.Context, req *CreateSpaceRequest) (*CreateSpaceResponse, error) {
	createResp, err := r.client.CreateSpace(req.Space)
	if err != nil {
		return nil, err
	}
	return &CreateSpaceResponse{
		Space: req.Space,
		Link:  createResp.Link,
	}, err
}

func (r *helmRegistry) DeleteSpace(ctx context.Context, req *DeleteSpaceRequest) (*DeleteSpaceResponse, error) {
	err := r.client.DeleteSpace(req.Space)
	if err != nil {
		return nil, err
	}
	return &DeleteSpaceResponse{}, nil
}

// 列取myspace下的所有chart
func (r *helmRegistry) ListCharts(ctx context.Context, req *ListChartsRequest) (*ListChartsResponse, error) {
	listResp, err := r.client.ListCharts(req.Space, req.Start, req.Limit)
	if err != nil {
		return nil, err
	}

	return &ListChartsResponse{
		Charts: listResp.Items,
		IsEnd:  req.Start+len(listResp.Items) >= listResp.Metadata.Total,
		space:  req.Space,
		offset: req.Start,
		limit:  req.Limit,
	}, nil
}

// 列取mysapce/mychart的所有版本
func (r *helmRegistry) ListChartVersions(ctx context.Context, req *ListChartVersionsRequest) (*ListChartVersionsResponse, error) {
	listResp, err := r.client.ListVersions(req.Space, req.Chart, req.offset, req.Limit)
	if err != nil {
		return nil, err
	}

	return &ListChartVersionsResponse{
		Versions: listResp.Items,
		IsEnd:    req.offset+len(listResp.Items) >= listResp.Metadata.Total,
		space:    req.Space,
		chart:    req.Chart,
		offset:   req.offset,
		limit:    req.Limit,
	}, nil
}

// 获取myspace/mychart:ver的metadata
func (r *helmRegistry) GetChartMetadata(ctx context.Context, req *GetChartMetadataRequest) (*GetChartMetadataResponse, error) {
	tmpResp, err := r.client.FetchVersionMetadata(req.Space, req.Chart, req.Version)
	if err != nil {
		return nil, err
	}

	resp := GetChartMetadataResponse{
		Metadata:     &tmpResp.Metadata,
		Dependencies: make([]*chart.Metadata, len(tmpResp.Dependencies), len(tmpResp.Dependencies)),
	}
	for i := range tmpResp.Dependencies {
		dep := tmpResp.Dependencies[i]
		resp.Dependencies[i] = &dep.Metadata
	}
	return &resp, nil
}

// 获取myspace/mychart:ver的values
func (r *helmRegistry) GetChartValues(ctx context.Context, req *GetChartValuesRequest) (*GetChartValuesResponse, error) {
	raw, err := r.client.FetchVersionValues(req.Space, req.Chart, req.Version)
	if err != nil {
		return nil, err
	}
	values := map[string]interface{}{}
	err = yaml.Unmarshal(raw, &values)
	if err != nil {
		return nil, err
	}

	return &GetChartValuesResponse{
		Values: values,
	}, nil
}

func (r *helmRegistry) fetchChartFile(ctx context.Context, space, chart, ver, file string) ([]byte, error) {
	data, err := r.client.DownloadVersion(space, chart, ver)
	if err != nil {
		return nil, err
	}
	ch, err := chartutil.LoadArchive(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for _, f := range ch.Files {
		if f.TypeUrl == file {
			return f.Value, nil
		}
	}
	return []byte{}, nil
}

// 获取myspace/mychart:ver的README.md
func (r *helmRegistry) GetChartReadme(ctx context.Context, req *GetChartReadmeRequest) (*GetChartReadmeResponse, error) {
	data, err := r.fetchChartFile(ctx, req.Space, req.Chart, req.Version, readmeFile)
	if err != nil {
		return nil, err
	}
	return &GetChartReadmeResponse{data}, nil
}

func (r *helmRegistry) GetChartRequirements(ctx context.Context, req *GetChartRequirementsRequest) (*GetChartRequirementsResponse, error) {
	data, err := r.fetchChartFile(ctx, req.Space, req.Chart, req.Version, requirementsFile)
	if err != nil {
		return nil, err
	}
	var deps []*chartutil.Dependency
	if len(data) > 0 {
		r := &chartutil.Requirements{}
		err := yaml.Unmarshal(data, r)
		if err != nil {
			return nil, err
		}
		deps = r.Dependencies
	}
	return &GetChartRequirementsResponse{deps}, nil
}

// 推送myspace/mychart:ver
func (r *helmRegistry) PushChart(ctx context.Context, req *PushChartRequest) (*PushChartResponse, error) {
	pushResp, err := r.client.UploadChart(req.Space, req.Data)
	if err != nil {
		return nil, err
	}
	return &PushChartResponse{
		Space:   pushResp.Space,
		Chart:   pushResp.Chart,
		Version: pushResp.Version,
		Link:    pushResp.Link,
	}, nil
}

func (r *helmRegistry) FetchChart(ctx context.Context, req *FetchChartRequest) (*FetchChartResponse, error) {
	data, err := r.client.DownloadVersion(req.Space, req.Chart, req.Version)
	if err != nil {
		return nil, err
	}
	return &FetchChartResponse{
		Data: data,
	}, nil
}
