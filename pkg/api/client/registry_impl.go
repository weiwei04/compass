package client

import (
	"fmt"

	"k8s.io/helm/pkg/proto/hapi/chart"

	v1 "github.com/caicloud/helm-registry/pkg/rest/v1"
	"golang.org/x/net/context"
	yaml "gopkg.in/yaml.v2"
)

func NewHelmRegistryClient(addr string) Registry {
	return &helmRegistry{addr: addr}
}

type helmRegistry struct {
	addr   string
	client *v1.Client
}

func (r *helmRegistry) Connect() error { return nil }

func (r *helmRegistry) Shutdown() {}

func (r *helmRegistry) ListSpaces(ctx context.Context, req *ListSpacesRequest) (*ListSpacesResponse, error) {
	listResp, err := r.client.ListSpaces(req.offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return &ListSpacesResponse{
		Spaces: listResp.Items,
		IsEnd:  req.offset+len(listResp.Items) >= listResp.Metadata.Total,
		offset: req.offset,
		limit:  req.Limit,
	}, nil
}

// 列取myspace下的所有chart
func (r *helmRegistry) ListCharts(ctx context.Context, req *ListChartsRequest) (*ListChartsResponse, error) {
	listResp, err := r.client.ListCharts(req.Space, req.offset, req.Limit)
	if err != nil {
		return nil, err
	}

	return &ListChartsResponse{
		Charts: listResp.Items,
		IsEnd:  req.offset+len(listResp.Items) >= listResp.Metadata.Total,
		space:  req.Space,
		offset: req.offset,
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

// 获取mysapce/mychart:ver的依赖说明
//GetChartRequiremens(context.Context, *GetChartRequirementsRequest) (*GetChartRequirementsResponse, error)
// 获取myspace/mychart:ver的README.md
func (r *helmRegistry) GetChartReadme(context.Context, *GetChartReadmeRequest) (*GetChartReadmeResponse, error) {
	return nil, fmt.Errorf("unimplemented")
}

// 推送myspace/mychart:ver
//PushChart(context.Context, *PushChartRequest) (*PushChartResponse, error)
