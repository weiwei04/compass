package compass

import (
	api "github.com/weiwei04/compass/pkg/api/services/compass"
	context "golang.org/x/net/context"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
)

var _ api.CompassServiceServer = &CompassServer{}

func (s *CompassServer) CreateCompassRelease(ctx context.Context, req *api.CreateCompassReleaseRequest) (*api.CreateCompassReleaseResponse, error) {
	var resp api.CreateCompassReleaseResponse
	chart, err := s.registry.Get(req.GetChart())
	if err != nil {
		return &resp, err
	}

	// TODO: merge values
	values := req.Values
	if values == nil {
		values = chart.Values
	}
	tillerReq := &tiller.InstallReleaseRequest{
		Chart:     chart,
		Values:    values,
		DryRun:    false,
		Name:      req.Name,
		Namespace: req.Namespace,
		ReuseName: false,
		Timeout:   req.Timeout,
	}

	tillerResp, err := s.tiller.InstallRelease(newContext(), tillerReq)
	if tillerResp != nil {
		resp.Release = tillerResp.Release
	}
	return &resp, err
}
