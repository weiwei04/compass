package client

import (
	"golang.org/x/net/context"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
)

type ListSpacesRequest struct {
	Start int
	Limit int
}

type ListSpacesResponse struct {
	Spaces []string
	offset int
	limit  int
	IsEnd  bool
}

type CreateSpaceRequest struct {
	Space string
}

type CreateSpaceResponse struct {
	Space string
	Link  string
}

type DeleteSpaceRequest CreateSpaceRequest

type DeleteSpaceResponse struct{}

type ListChartsRequest struct {
	Space string
	Start int
	Limit int
}

type ListChartsResponse struct {
	Charts []string
	space  string
	offset int
	limit  int
	IsEnd  bool
}

type ListChartVersionsRequest struct {
	Space string
	Chart string
	Limit int
	Start int
}

type ListChartVersionsResponse struct {
	Versions []string
	IsEnd    bool
	space    string
	chart    string
	limit    int
	offset   int
}

type GetChartMetadataRequest struct {
	Space   string
	Chart   string
	Version string
}

type GetChartMetadataResponse struct {
	Metadata     *chart.Metadata
	Dependencies []*chart.Metadata
}

type GetChartRequirementsRequest GetChartMetadataRequest

type GetChartRequirementsResponse struct {
	Dependencies []*chartutil.Dependency
}

type GetChartValuesRequest GetChartMetadataRequest

type GetChartValuesResponse struct {
	Values map[string]interface{}
}

type GetChartReadmeRequest GetChartMetadataRequest

type GetChartReadmeResponse struct {
	Readme []byte
}

type PushChartRequest struct {
	Space string
	Data  []byte
}

type PushChartResponse struct {
	Space   string
	Chart   string
	Version string
	Link    string
}

type FetchChartRequest GetChartMetadataRequest

type FetchChartResponse struct {
	Data []byte
}

type DeleteVersionRequest GetChartMetadataRequest

type DeleteVersionResponse DeleteSpaceResponse

type DeleteChartRequest struct {
	Space string
	Chart string
}

type DeleteChartResponse DeleteSpaceResponse

type Registry interface {
	// 列取space
	ListSpaces(context.Context, *ListSpacesRequest) (*ListSpacesResponse, error)
	CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error)
	DeleteSpace(context.Context, *DeleteSpaceRequest) (*DeleteSpaceResponse, error)
	// 列取myspace下的所有chart
	ListCharts(context.Context, *ListChartsRequest) (*ListChartsResponse, error)
	// 列取mysapce/mychart的所有版本
	ListChartVersions(context.Context, *ListChartVersionsRequest) (*ListChartVersionsResponse, error)
	// 获取myspace/mychart:ver的metadata
	GetChartMetadata(context.Context, *GetChartMetadataRequest) (*GetChartMetadataResponse, error)
	// 获取myspace/mychart:ver的values
	GetChartValues(context.Context, *GetChartValuesRequest) (*GetChartValuesResponse, error)
	// 获取mysapce/mychart:ver的依赖说明
	GetChartRequirements(context.Context, *GetChartRequirementsRequest) (*GetChartRequirementsResponse, error)
	// 获取myspace/mychart:ver的README.md
	GetChartReadme(context.Context, *GetChartReadmeRequest) (*GetChartReadmeResponse, error)
	// 推送myspace/mychart:ver
	PushChart(context.Context, *PushChartRequest) (*PushChartResponse, error)

	FetchChart(context.Context, *FetchChartRequest) (*FetchChartResponse, error)

	// 删除某个版本的 chart
	DeleteVersion(context.Context, *DeleteVersionRequest) (*DeleteVersionResponse, error)
	// 删除全部 chart
	DeleteChart(context.Context, *DeleteChartRequest) (*DeleteChartResponse, error)
}
