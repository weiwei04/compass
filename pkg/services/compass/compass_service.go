package compass

import (
	api "github.com/weiwei04/compass/pkg/api/services/compass"
	context "golang.org/x/net/context"
)

var _ api.CompassServiceServer = &CompassServer{}

func (s *CompassServer) CreateCompassRelease(ctx context.Context, in *api.CreateCompassReleaseRequest) (*api.CreateCompassReleaseResponse, error) {
	return nil, nil
}
