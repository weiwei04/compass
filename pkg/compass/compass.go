package compass

import context "golang.org/x/net/context"
import hapi_services_tiller "k8s.io/helm/pkg/proto/hapi/services"
import services "github.com/weiwei04/compass/pkg/proto/compass/services"

type CompassServer struct {
}

func NewCompassServer() *CompassServer {
	return &CompassServer{}
}

var _ services.CompassServiceServer = &CompassServer{}

func (s *CompassServer) ListReleases(context.Context, *hapi_services_tiller.ListReleasesRequest) (*hapi_services_tiller.ListReleasesResponse, error) {
	return nil, nil
}

// GetReleasesStatus retrieves status information for the specified release.
func (s *CompassServer) GetReleaseStatus(context.Context, *hapi_services_tiller.GetReleaseStatusRequest) (*hapi_services_tiller.GetReleaseStatusResponse, error) {
	return nil, nil
}

// // GetReleaseContent retrieves the release content (chart + value) for the specified release.
func (s *CompassServer) GetReleaseContent(context.Context, *hapi_services_tiller.GetReleaseContentRequest) (*hapi_services_tiller.GetReleaseContentResponse, error) {
	return nil, nil
}

// // UpdateRelease updates release content.
func (s *CompassServer) UpdateRelease(context.Context, *hapi_services_tiller.UpdateReleaseRequest) (*hapi_services_tiller.UpdateReleaseResponse, error) {
	return nil, nil
}

// // InstallRelease requests installation of a chart as a new release.
func (s *CompassServer) InstallRelease(context.Context, *hapi_services_tiller.InstallReleaseRequest) (*hapi_services_tiller.InstallReleaseResponse, error) {
	return nil, nil
}

// // UninstallRelease requests deletion of a named release.
func (s *CompassServer) UninstallRelease(context.Context, *hapi_services_tiller.UninstallReleaseRequest) (*hapi_services_tiller.UninstallReleaseResponse, error) {
	return nil, nil
}

// // GetVersion returns the current version of the server.
func (s *CompassServer) GetVersion(context.Context, *hapi_services_tiller.GetVersionRequest) (*hapi_services_tiller.GetVersionResponse, error) {
	return nil, nil
}

// // RollbackRelease rolls back a release to a previous version.
func (s *CompassServer) RollbackRelease(context.Context, *hapi_services_tiller.RollbackReleaseRequest) (*hapi_services_tiller.RollbackReleaseResponse, error) {
	return nil, nil
}

// // ReleaseHistory retrieves a releasse's history.
func (s *CompassServer) GetHistory(context.Context, *hapi_services_tiller.GetHistoryRequest) (*hapi_services_tiller.GetHistoryResponse, error) {
	return nil, nil
}

// // RunReleaseTest executes the tests defined of a named release
func (s *CompassServer) RunReleaseTest(context.Context, *hapi_services_tiller.TestReleaseRequest) (*hapi_services_tiller.TestReleaseResponse, error) {
	return nil, nil
}
