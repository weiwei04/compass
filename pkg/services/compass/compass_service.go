package compass

import (
	"io"
	"time"

	"github.com/golang/glog"
	pb "github.com/weiwei04/compass/pkg/api/services/compass"
	"github.com/weiwei04/compass/pkg/chart"
	//"go.uber.org/zap"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
)

var _ pb.CompassServiceServer = &compassServer{}

func newContext() context.Context {
	md := metadata.Pairs("x-helm-api-client", "2.5.0")
	return metadata.NewContext(context.TODO(), md)
}

//func stripChartFiles(release *hapi_release.Release) *hapi_release.Release {
//	release.Chart.Files = nil
//	return release
//}
//
//// NOTE: protobuf.Any is a raw file, it can't Marshal into a proper json object
//func stripChartsFiles(releases []*hapi_release.Release) []*hapi_release.Release {
//	for i := range releases {
//		releases[i].Chart.Files = nil
//	}
//	return releases
//}

type compassServer struct {
	config Config
	//logger   *zap.Logger
	registry chart.Store
	conn     *grpc.ClientConn
	tiller   tiller.ReleaseServiceClient
}

func newCompassServer(ctx context.Context, config Config) (*compassServer, error) {
	s := &compassServer{
		config: config,
	}

	var err error

	// init conn to helm-registry
	if s.registry, err = chart.NewHelmRegistryStore(s.config.RegistryAddr); err != nil {
		glog.Errorf("NewHelmRegistryStore failed, err:", err)
		return nil, err
	}

	// init conn to tiller service
	opts := []grpc.DialOption{
		grpc.WithTimeout(10 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	for i := 0; i < 3; i++ {
		if s.conn, err = grpc.Dial(s.config.TillerAddr, opts...); err != nil {
			glog.Errorf("dial tiller failed, tiller:%s, err:%s", s.config.TillerAddr, err)
			return nil, err
		}
	}
	go func() {
		<-ctx.Done()
		if cerr := s.conn.Close(); cerr != nil {
			glog.Errorf("Failed to close conn to, tiller:%s, err:%s",
				s.config.TillerAddr,
				cerr)
		} else {
			glog.Infof("Closed conn to tiller:%s", s.config.TillerAddr)
		}
		//s.logger.Sync()
	}()
	s.tiller = tiller.NewReleaseServiceClient(s.conn)

	return s, nil
}

func (s *compassServer) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*pb.CreateReleaseResponse, error) {
	glog.V(4).Infof("CreateRelease(name:%s, chart:%s)",
		req.Name, req.GetChart())
	var resp pb.CreateReleaseResponse
	chart, err := s.registry.Get(req.GetChart())
	if err != nil {
		glog.V(3).Infof("get chart failed, chart:%s, err:%s", req.GetChart(), err)
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
		resp.Release = pb.CompassRelease(tillerResp.Release)
	}
	return &resp, err
}

func (s *compassServer) UpdateRelease(ctx context.Context, req *pb.UpdateReleaseRequest) (*pb.UpdateReleaseResponse, error) {
	glog.V(4).Infof("UpdateRelease(name:%s, version:%d, namespace:%s)",
		req.Name, req.Version, req.Namespace)
	getReq := &tiller.GetReleaseContentRequest{
		Name:    req.Name,
		Version: req.Version,
	}
	getResp, err := s.tiller.GetReleaseContent(newContext(), getReq)
	if err != nil {
		glog.V(3).Infof("GetReleaseContent(name:%s, version:%d) failed, err:%s",
			req.Name, req.Version, err)
		return nil, err
	}
	chart := getResp.Release.Chart
	updateReq := &tiller.UpdateReleaseRequest{
		Name:   req.Name,
		Chart:  chart,
		Values: req.Values,
	}

	var resp pb.UpdateReleaseResponse
	updateResp, err := s.tiller.UpdateRelease(newContext(), updateReq)
	if err != nil {
		resp.Release = pb.CompassRelease(updateResp.Release)
	}

	return &resp, err
}

func (s *compassServer) UpgradeRelease(ctx context.Context, req *pb.UpgradeReleaseRequest) (*pb.UpgradeReleaseResponse, error) {
	glog.V(4).Infof("UpgradeRelease(name:%s, chart:%s, namespace:%s)",
		req.Name, req.GetChart(), req.Namespace)
	chart, err := s.registry.Get(req.GetChart())
	if err != nil {
		glog.V(3).Infof("GetChart(chart:%s) from registry failed, err:%s", req.GetChart(), err)
		return nil, err
	}

	upgradeReq := &tiller.UpdateReleaseRequest{
		Name:   req.Name,
		Chart:  chart,
		Values: req.Values,
	}

	// TODO: wrap ctx with newContext
	var resp pb.UpgradeReleaseResponse
	upgradeResp, err := s.tiller.UpdateRelease(newContext(), upgradeReq)
	if err != nil {
		resp.Release = pb.CompassRelease(upgradeResp.Release)
	}
	return &resp, err
}

func (s *compassServer) DeleteRelease(ctx context.Context, req *pb.DeleteReleaseRequest) (*pb.DeleteReleaseResponse, error) {
	glog.V(4).Infof("DeleteRelease(name:%s, namespace:%s)",
		req.Name, req.Namespace)
	uninstallReq := &tiller.UninstallReleaseRequest{
		Name:         req.Name,
		DisableHooks: req.DisableHooks,
		Purge:        true,
		Timeout:      req.Timeout,
	}
	uninstallResp, err := s.tiller.UninstallRelease(newContext(), uninstallReq)
	if err != nil {
		glog.V(3).Infof("UninstallRelease(name:%s) failed, err:%s",
			req.Name, err)
		return nil, err
	}
	return &pb.DeleteReleaseResponse{
		Release: pb.CompassRelease(uninstallResp.Release),
		Info:    uninstallResp.Info,
	}, err
}

func (s *compassServer) ListReleases(ctx context.Context, req *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	glog.V(4).Infof("ListReleases(limit:%d, offset:%s, namespace:%s)",
		req.Limit, req.Offset, req.Namespace)
	listReq := &tiller.ListReleasesRequest{
		Limit:     req.Limit,
		Offset:    req.Offset,
		SortBy:    tiller.ListSort_NAME,
		SortOrder: tiller.ListSort_ASC,
		Namespace: req.Namespace,
	}

	cli, err := s.tiller.ListReleases(newContext(), listReq)
	if err != nil {
		glog.Errorf("ListReleases(limit:%d, offset:%s, namespace:%s) failed, err:%s",
			req.Limit, req.Offset, req.Namespace, err)
		return nil, err
	}

	var resp *pb.ListReleasesResponse
	for {
		msg, err := cli.Recv()
		if err == io.EOF {
			return resp, nil
		} else if err != nil {
			glog.Errorf("ListReleases stream.Recv failed, err:%s", err)
			return nil, err
		} else {
			resp = &pb.ListReleasesResponse{
				Count:    msg.Count,
				Next:     msg.Next,
				Total:    msg.Total,
				Releases: pb.CompassReleaseSlice(msg.Releases),
			}
		}
	}

	return resp, nil
}

func (s *compassServer) GetReleaseStatus(ctx context.Context, req *pb.GetReleaseStatusRequest) (*pb.GetReleaseStatusResponse, error) {
	glog.V(4).Infof("GetReleaseStatus(name:%s, version:%d, namespace:%s)",
		req.Name, req.Version, req.Namespace)
	statusReq := &tiller.GetReleaseStatusRequest{
		Name:    req.Name,
		Version: req.Version,
	}
	resp, err := s.tiller.GetReleaseStatus(newContext(), statusReq)
	if err != nil {
		glog.V(3).Infof("GetReleaseStatus(name:%s, version:%d) failed, err:%s",
			req.Name, req.Version, err)
		return nil, err
	}
	return &pb.GetReleaseStatusResponse{
		Name:      resp.Name,
		Info:      pb.CompassReleaseInfo(resp.Info),
		Namespace: resp.Namespace,
	}, nil
}

func (s *compassServer) GetReleaseContent(ctx context.Context, req *pb.GetReleaseContentRequest) (*pb.GetReleaseContentResponse, error) {
	glog.V(4).Infof("GetReleaseContent(name:%s, version:%d, namespace:%s)",
		req.Name, req.Version, req.Namespace)
	contentReq := &tiller.GetReleaseContentRequest{
		Name:    req.Name,
		Version: req.Version,
	}
	resp, err := s.tiller.GetReleaseContent(newContext(), contentReq)
	if err != nil {
		glog.V(3).Infof("GetReleaseContent(name:%s, version:%d) failed, err:%s",
			req.Name, req.Version, err)
		return nil, err
	}
	return &pb.GetReleaseContentResponse{
		Release: pb.CompassRelease(resp.Release),
	}, nil
}

func (s *compassServer) GetReleaseHistory(ctx context.Context, req *pb.GetReleaseHistoryRequest) (*pb.GetReleaseHistoryResponse, error) {
	glog.V(4).Infof("GetReleaseHistory(name:%s, namespace:%s, max:%d)",
		req.Name, req.Namespace, req.Max)
	historyReq := &tiller.GetHistoryRequest{
		Name: req.Name,
		Max:  req.Max,
	}
	resp, err := s.tiller.GetHistory(newContext(), historyReq)
	if err != nil {
		glog.V(3).Infof("GetHistory(name:%s, max:%d) failed, err:%s",
			req.Name, req.Max, err)
		return nil, err
	}
	return &pb.GetReleaseHistoryResponse{
		Releases: pb.CompassReleaseSlice(resp.Releases),
	}, nil
}

func (s *compassServer) RollbackRelease(ctx context.Context, req *pb.RollbackReleaseRequest) (*pb.RollbackReleaseResponse, error) {
	glog.V(4).Infof("RollbackRelease(name:%s, namespace:%s, version:%d)",
		req.Name, req.Namespace, req.Version)
	rollbackReq := &tiller.RollbackReleaseRequest{
		Name:         req.Name,
		DryRun:       req.DryRun,
		DisableHooks: req.DisableHooks,
		Version:      req.Version,
		Recreate:     req.Recreate,
		Timeout:      req.Timeout,
		Wait:         req.Wait,
		Force:        req.Force,
	}
	resp, err := s.tiller.RollbackRelease(newContext(), rollbackReq)
	if err != nil {
		glog.V(3).Infof("RollbackRelease(name:%s, version:%d) failed, err:%s",
			req.Name, req.Version, err)
	}
	return &pb.RollbackReleaseResponse{
		Release: pb.CompassRelease(resp.Release),
	}, nil
}

func (s *compassServer) RunReleaseTest(req *pb.TestReleaseRequest, stream pb.CompassService_RunReleaseTestServer) error {
	glog.V(4).Infof("RunReleaseTest(name:%s, namespace:%s)",
		req.Name, req.Namespace)
	testReq := &tiller.TestReleaseRequest{
		Name:    req.Name,
		Timeout: req.Timeout,
		Cleanup: req.Cleanup,
	}
	cli, err := s.tiller.RunReleaseTest(newContext(), testReq)
	if err != nil {
		return err
	}

	for {
		msg, err := cli.Recv()
		if err == io.EOF {
			return nil
		} else if err == nil {
			return err
		} else {
			resp := &pb.TestReleaseResponse{
				Msg:    msg.Msg,
				Status: msg.Status,
			}
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}

	return nil
}
