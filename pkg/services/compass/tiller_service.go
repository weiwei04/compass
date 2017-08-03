package compass

import (
	"io"
	//"log"

	context "golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
	//"k8s.io/helm/pkg/version"
)

func newContext() context.Context {
	md := metadata.Pairs("x-helm-api-client", "2.5.0")
	return metadata.NewContext(context.TODO(), md)
}

func (s *CompassServer) ListReleases(req *tiller.ListReleasesRequest, stream tiller.ReleaseService_ListReleasesServer) error {
	ctx := newContext()
	cli, err := s.tiller.ListReleases(ctx, req)
	if err != nil {
		return err
	}

	for {
		msg, err := cli.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(msg)
		if err != nil {
			return err
		}
	}
}

// GetReleasesStatus retrieves status information for the specified release.
func (s *CompassServer) GetReleaseStatus(ctx context.Context, req *tiller.GetReleaseStatusRequest) (*tiller.GetReleaseStatusResponse, error) {
	return s.tiller.GetReleaseStatus(ctx, req)
}

// // GetReleaseContent retrieves the release content (chart + value) for the specified release.
func (s *CompassServer) GetReleaseContent(ctx context.Context, req *tiller.GetReleaseContentRequest) (*tiller.GetReleaseContentResponse, error) {
	return s.tiller.GetReleaseContent(ctx, req)
}

// // UpdateRelease updates release content.
func (s *CompassServer) UpdateRelease(ctx context.Context, req *tiller.UpdateReleaseRequest) (*tiller.UpdateReleaseResponse, error) {
	return s.tiller.UpdateRelease(ctx, req)
}

// // InstallRelease requests installation of a chart as a new release.
func (s *CompassServer) InstallRelease(ctx context.Context, req *tiller.InstallReleaseRequest) (*tiller.InstallReleaseResponse, error) {
	return s.tiller.InstallRelease(ctx, req)
}

// // UninstallRelease requests deletion of a named release.
func (s *CompassServer) UninstallRelease(ctx context.Context, req *tiller.UninstallReleaseRequest) (*tiller.UninstallReleaseResponse, error) {
	return s.tiller.UninstallRelease(ctx, req)
}

// // GetVersion returns the current version of the server.
func (s *CompassServer) GetVersion(ctx context.Context, req *tiller.GetVersionRequest) (*tiller.GetVersionResponse, error) {
	return s.tiller.GetVersion(ctx, req)
}

// // RollbackRelease rolls back a release to a previous version.
func (s *CompassServer) RollbackRelease(ctx context.Context, req *tiller.RollbackReleaseRequest) (*tiller.RollbackReleaseResponse, error) {
	return s.tiller.RollbackRelease(ctx, req)
}

// // ReleaseHistory retrieves a releasse's history.
func (s *CompassServer) GetHistory(ctx context.Context, req *tiller.GetHistoryRequest) (*tiller.GetHistoryResponse, error) {
	return s.tiller.GetHistory(ctx, req)
}

// // RunReleaseTest executes the tests defined of a named release
func (s *CompassServer) RunReleaseTest(req *tiller.TestReleaseRequest, stream tiller.ReleaseService_RunReleaseTestServer) error {
	ctx := newContext()
	cli, err := s.tiller.RunReleaseTest(ctx, req)
	if err != nil {
		return err
	}

	for {
		msg, err := cli.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(msg)
		if err != nil {
			return err
		}
	}
}
