package compass

import (
	api "github.com/weiwei04/compass/pkg/api/services/compass"
	"go.uber.org/zap"
	context "golang.org/x/net/context"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
)

var _ api.CompassServiceServer = &CompassServer{}

func (s *CompassServer) CreateCompassRelease(ctx context.Context, req *api.CreateCompassReleaseRequest) (*api.CreateCompassReleaseResponse, error) {
	var resp api.CreateCompassReleaseResponse
	chart, err := s.registry.Get(req.GetChart())
	if err != nil {
		s.logger.Debug("get chart failed",
			zap.String("chart", req.GetChart()),
			zap.Error(err))
		return &resp, err
	}

	// TODO: improve this
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

	// TODO: wrap ctx with newContext
	tillerResp, err := s.tiller.InstallRelease(newContext(), tillerReq)
	if tillerResp != nil {
		resp.Release = tillerResp.Release
	}
	return &resp, err
}

func (s *CompassServer) UpdateCompassRelease(ctx context.Context, req *api.UpdateCompassReleaseRequest) (*api.UpdateCompassReleaseResponse, error) {
	getReq := &tiller.GetReleaseContentRequest{
		Name: req.Name,
		//Version: req.Version,
	}
	getResp, err := s.tiller.GetReleaseContent(newContext(), getReq)
	if err != nil {
		return nil, err
	}
	chart := getResp.Release.Chart
	updateReq := &tiller.UpdateReleaseRequest{
		Name:   req.Name,
		Chart:  chart,
		Values: req.Values,
	}

	var resp api.UpdateCompassReleaseResponse
	updateResp, err := s.tiller.UpdateRelease(newContext(), updateReq)
	if updateResp != nil {
		resp.Release = updateResp.Release
	}

	return &resp, err
}

func (s *CompassServer) UpgradeCompassRelease(ctx context.Context, req *api.UpgradeCompassReleaseRequest) (*api.UpgradeCompassReleaseResponse, error) {
	chart, err := s.registry.Get(req.GetChart())
	if err != nil {
		s.logger.Debug("get chart failed",
			zap.String("chart", req.GetChart()),
			zap.Error(err))
		return nil, err
	}

	upgradeReq := &tiller.UpdateReleaseRequest{
		Name:   req.Name,
		Chart:  chart,
		Values: req.Values,
	}

	// TODO: wrap ctx with newContext
	var resp api.UpgradeCompassReleaseResponse
	upgradeResp, err := s.tiller.UpdateRelease(newContext(), upgradeReq)
	if upgradeResp != nil {
		resp.Release = upgradeResp.Release
	}
	return &resp, err
}
