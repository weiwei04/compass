package compass

import (
	pb "github.com/weiwei04/compass/pkg/api/services/compass"
	context "golang.org/x/net/context"
)

type fakeCompassServer struct{}

func newFakeCompassServer() *fakeCompassServer {
	return &fakeCompassServer{}
}

var _ pb.CompassServiceServer = &fakeCompassServer{}

func (s *fakeCompassServer) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*pb.CreateReleaseResponse, error) {
	return pb.NewFakeCreateReleaseResponse(req), nil
}

func (s *fakeCompassServer) UpdateRelease(ctx context.Context, req *pb.UpdateReleaseRequest) (*pb.UpdateReleaseResponse, error) {
	return pb.NewFakeUpdateReleaseResponse(req), nil
}

func (s *fakeCompassServer) UpgradeRelease(ctx context.Context, req *pb.UpgradeReleaseRequest) (*pb.UpgradeReleaseResponse, error) {
	return pb.NewFakeUpgradeReleaseResponse(req), nil
}

func (s *fakeCompassServer) DeleteRelease(ctx context.Context, req *pb.DeleteReleaseRequest) (*pb.DeleteReleaseResponse, error) {
	return pb.NewFakeDeleteReleaseResponse(req), nil
}

func (s *fakeCompassServer) ListReleases(ctx context.Context, req *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	return pb.NewFakeListReleasesResponse(req), nil
}

func (s *fakeCompassServer) GetReleaseStatus(ctx context.Context, req *pb.GetReleaseStatusRequest) (*pb.GetReleaseStatusResponse, error) {
	return pb.NewFakeGetReleaseStatusResponse(req), nil
}

func (s *fakeCompassServer) GetReleaseContent(ctx context.Context, req *pb.GetReleaseContentRequest) (*pb.GetReleaseContentResponse, error) {
	return pb.NewFakeGetReleaseContentResponse(req), nil
}

func (s *fakeCompassServer) GetReleaseHistory(ctx context.Context, req *pb.GetReleaseHistoryRequest) (*pb.GetReleaseHistoryResponse, error) {
	return pb.NewFakeGetReleaseHistoryResponse(req), nil
}

func (s *fakeCompassServer) RollbackRelease(ctx context.Context, req *pb.RollbackReleaseRequest) (*pb.RollbackReleaseResponse, error) {
	return pb.NewFakeRollbackReleaseResponse(req), nil
}

func (s *fakeCompassServer) RunReleaseTest(req *pb.TestReleaseRequest, stream pb.CompassService_RunReleaseTestServer) error {
	return stream.Send(pb.NewFakeTestReleaseResponse(req))
}
