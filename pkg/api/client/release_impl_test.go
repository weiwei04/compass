package client

import (
	"context"
	"time"

	"github.com/stretchr/testify/require"
	pb "github.com/weiwei04/compass/pkg/api/services/compass"
	//hapi_chart "k8s.io/helm/pkg/proto/hapi/chart"
	//hapi_release "k8s.io/helm/pkg/proto/hapi/release"
	"testing"

	"github.com/weiwei04/compass/pkg/services/compass"
)

func TestReleaseImpl(t *testing.T) {
	srv := compass.NewServer(compass.Config{
		TillerAddr:   "127.0.0.1:8910",
		RegistryAddr: "http://127.0.0.1:8911",
		RPCAddr:      "127.0.0.1:8912",
		RESTAddr:     "127.0.0.1:8913",
		Mock:         true,
	})

	go func() {
		// TODO: need to know it's started
		srv.Serve()
	}()

	time.Sleep(5 * time.Second)

	require := require.New(t)

	cli := NewReleaseClient("http://127.0.0.1:8913", &logger{})
	ctx := context.Background()

	// CreateRelease
	createReq := &CreateReleaseRequest{
		Chart:     "mynamespace/mychart:v0.1.0",
		Name:      "release0",
		Namespace: "namespace0",
	}
	createResp, err := cli.CreateRelease(ctx, createReq)
	require.NoError(err, "cli.CreateRelease")
	require.Equal(pb.NewFakeCreateReleaseResponse(
		(*pb.CreateReleaseRequest)(createReq)),
		(*pb.CreateReleaseResponse)(createResp))

	// DeleteRelease
	deleteReq := &DeleteReleaseRequest{
		Name:      "release0",
		Namespace: "namespace0",
	}
	deleteResp, err := cli.DeleteRelease(ctx, deleteReq)
	require.NoError(err, "cli.DeleteRelease")
	require.Equal(pb.NewFakeDeleteReleaseResponse(
		(*pb.DeleteReleaseRequest)(deleteReq)),
		(*pb.DeleteReleaseResponse)(deleteResp))

	//// UpdateRelease
	updateReq := &UpdateReleaseRequest{
		Name:      "release0",
		Namespace: "namespace0",
		//Values:    &hapi_chart.Config{},
	}
	updateResp, err := cli.UpdateRelease(ctx, updateReq)
	require.NoError(err, "cli.UpdateRelease")
	require.Equal(pb.NewFakeUpdateReleaseResponse(
		(*pb.UpdateReleaseRequest)(updateReq)),
		(*pb.UpdateReleaseResponse)(updateResp))

	//// UpgradeRelease
	upgradeReq := &UpgradeReleaseRequest{
		Chart: "myspace/mychart:v0.1.0",
		Name:  "release0",
		//Values:
		Namespace: "namespace0",
	}
	upgradeResp, err := cli.UpgradeRelease(ctx, upgradeReq)
	require.NoError(err)
	require.Equal(pb.NewFakeUpgradeReleaseResponse(
		(*pb.UpgradeReleaseRequest)(upgradeReq)),
		(*pb.UpgradeReleaseResponse)(upgradeResp))

	// ListReleases
	listReq := &ListReleasesRequest{
		Limit:     10,
		Offset:    "next",
		Namespace: "namespace0",
	}
	listResp, err := cli.ListReleases(ctx, listReq)
	require.NoError(err)
	require.Equal(pb.NewFakeListReleasesResponse(
		(*pb.ListReleasesRequest)(listReq)),
		(*pb.ListReleasesResponse)(listResp))

	//// GetReleaseContent
	getContentReq := &GetReleaseContentRequest{
		Name:      "release0",
		Version:   8,
		Namespace: "namespace0",
	}
	getContentResp, err := cli.GetReleaseContent(ctx, getContentReq)
	require.NoError(err)
	require.Equal(pb.NewFakeGetReleaseContentResponse(
		(*pb.GetReleaseContentRequest)(getContentReq)),
		(*pb.GetReleaseContentResponse)(getContentResp))

	// GetReleaseStatus
	getStatusReq := &GetReleaseStatusRequest{
		Name:      "release0",
		Version:   0,
		Namespace: "namespace0",
	}
	getStatusResp, err := cli.GetReleaseStatus(ctx, getStatusReq)
	require.NoError(err)
	require.Equal(pb.NewFakeGetReleaseStatusResponse(
		(*pb.GetReleaseStatusRequest)(getStatusReq)),
		(*pb.GetReleaseStatusResponse)(getStatusResp))

	// GetReleaseHistory
	historyReq := &GetReleaseHistoryRequest{
		Name:      "release0",
		Max:       10,
		Namespace: "namespace0",
	}
	historyResp, err := cli.GetReleaseHistory(ctx, historyReq)
	require.NoError(err)
	require.Equal(pb.NewFakeGetReleaseHistoryResponse(
		(*pb.GetReleaseHistoryRequest)(historyReq)),
		(*pb.GetReleaseHistoryResponse)(historyResp))

	// RollbackRelease
	rollbackReq := &RollbackReleaseRequest{
		Name:      "release0",
		Version:   1,
		Namespace: "namespace0",
	}
	rollbackResp, err := cli.RollbackRelease(ctx, rollbackReq)
	require.NoError(err)
	require.Equal(pb.NewFakeRollbackReleaseResponse(
		(*pb.RollbackReleaseRequest)(rollbackReq)),
		(*pb.RollbackReleaseResponse)(rollbackResp))
}
