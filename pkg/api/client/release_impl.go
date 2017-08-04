package client

import (
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	tillerapi "k8s.io/helm/pkg/proto/hapi/services"
)

// all the ugly things

// NewCompassClient will return a default client impl
func NewCompassReleaseClient(addr string) Release {
	return &compassClient{addr: addr}
}

type compassClient struct {
	addr  string
	capi  compassapi.CompassServiceClient
	tapi  tillerapi.ReleaseServiceClient
	cconn *grpc.ClientConn
	tconn *grpc.ClientConn
}

func (c *compassClient) Connect() error {
	if c.cconn != nil {
		return nil
	}

	var err error

	defer func() {
		if err == nil {
			return
		}
		c.Shutdown()
	}()

	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	c.cconn, err = grpc.Dial(c.addr, opts...)
	if err != nil {
		return err
	}
	c.capi = compassapi.NewCompassServiceClient(c.cconn)

	c.tconn, err = grpc.Dial(c.addr, opts...)
	if err != nil {
		return err
	}
	c.tapi = tillerapi.NewReleaseServiceClient(c.tconn)

	return nil
}

func (c *compassClient) Shutdown() {
	if c.cconn != nil {
		c.cconn.Close()
	}
	if c.tconn != nil {
		c.tconn.Close()
	}
}

func (c *compassClient) CreateRelease(ctx context.Context, req *CreateReleaseRequest) (*CreateReleaseResponse, error) {
	resp, err := c.capi.CreateCompassRelease(ctx, (*compassapi.CreateCompassReleaseRequest)(req))
	return (*CreateReleaseResponse)(resp), err
}

func (c *compassClient) GetReleaseStatus(ctx context.Context, req *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error) {
	resp, err := c.tapi.GetReleaseStatus(ctx, (*tillerapi.GetReleaseStatusRequest)(req))
	return (*GetReleaseStatusResponse)(resp), err
}

func (c *compassClient) GetReleaseContent(ctx context.Context, req *GetReleaseContentRequest) (*GetReleaseContentResponse, error) {
	resp, err := c.tapi.GetReleaseContent(ctx, (*tillerapi.GetReleaseContentRequest)(req))
	return (*GetReleaseContentResponse)(resp), err
}

func (c *compassClient) ListReleases(ctx context.Context, req *ListReleasesRequest) (ListReleasesClient, error) {
	client, err := c.tapi.ListReleases(ctx, (*tillerapi.ListReleasesRequest)(req))
	return ListReleasesClient(client), err
}

func (c *compassClient) UpdateRelease(ctx context.Context, req *UpdateReleaseRequest) (*UpdateReleaseResponse, error) {
	resp, err := c.capi.UpdateCompassRelease(ctx, (*compassapi.UpdateCompassReleaseRequest)(req))
	return (*UpdateReleaseResponse)(resp), err
}

func (c *compassClient) UpgradeRelease(ctx context.Context, req *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error) {
	resp, err := c.capi.UpgradeCompassRelease(ctx, (*compassapi.UpgradeCompassReleaseRequest)(req))
	return (*UpgradeReleaseResponse)(resp), err
}

func (c *compassClient) DeleteRelease(ctx context.Context, req *DeleteReleaseRequest) (*DeleteReleaseResponse, error) {
	resp, err := c.tapi.UninstallRelease(ctx, (*tillerapi.UninstallReleaseRequest)(req))
	return (*DeleteReleaseResponse)(resp), err
}

func (c *compassClient) GetHistory(ctx context.Context, req *GetHistoryRequest) (*GetHistoryResponse, error) {
	resp, err := c.tapi.GetHistory(ctx, (*tillerapi.GetHistoryRequest)(req))
	return (*GetHistoryResponse)(resp), err
}

func (c *compassClient) RollbackRelease(ctx context.Context, req *RollbackReleaseRequest) (*RollbackReleaseResponse, error) {
	resp, err := c.tapi.RollbackRelease(ctx, (*tillerapi.RollbackReleaseRequest)(req))
	return (*RollbackReleaseResponse)(resp), err
}
