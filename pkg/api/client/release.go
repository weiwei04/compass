package client

import (
	api "github.com/weiwei04/compass/pkg/api/services/compass"
	"golang.org/x/net/context"
)

// all the ugly things

type CreateReleaseRequest api.CreateReleaseRequest
type CreateReleaseResponse api.CreateReleaseResponse

type UpdateReleaseRequest api.UpdateReleaseRequest
type UpdateReleaseResponse api.UpdateReleaseResponse

type UpgradeReleaseRequest api.UpgradeReleaseRequest
type UpgradeReleaseResponse api.UpgradeReleaseResponse

type ListReleasesRequest api.ListReleasesRequest
type ListReleasesResponse api.ListReleasesResponse

type GetReleaseStatusRequest api.GetReleaseStatusRequest
type GetReleaseStatusResponse api.GetReleaseStatusResponse

type GetReleaseContentRequest api.GetReleaseContentRequest
type GetReleaseContentResponse api.GetReleaseContentResponse

type DeleteReleaseRequest api.DeleteReleaseRequest
type DeleteReleaseResponse api.DeleteReleaseResponse

type RollbackReleaseRequest api.RollbackReleaseRequest
type RollbackReleaseResponse api.RollbackReleaseResponse

type GetReleaseHistoryRequest api.GetReleaseHistoryRequest
type GetReleaseHistoryResponse api.GetReleaseHistoryResponse

// runtime release manage api

type Release interface {
	CreateRelease(ctx context.Context, in *CreateReleaseRequest) (*CreateReleaseResponse, error)
	DeleteRelease(ctx context.Context, in *DeleteReleaseRequest) (*DeleteReleaseResponse, error)
	UpdateRelease(ctx context.Context, in *UpdateReleaseRequest) (*UpdateReleaseResponse, error)
	UpgradeRelease(ctx context.Context, in *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error)
	// call the api like this
	// get /v1/namespaces/NAMESPACE/releases?limit=10&offset=denis
	ListReleases(ctx context.Context, in *ListReleasesRequest) (*ListReleasesResponse, error)
	GetReleaseStatus(ctx context.Context, in *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error)
	GetReleaseContent(ctx context.Context, in *GetReleaseContentRequest) (*GetReleaseContentResponse, error)
	GetReleaseHistory(ctx context.Context, in *GetReleaseHistoryRequest) (*GetReleaseHistoryResponse, error)
	RollbackRelease(ctx context.Context, in *RollbackReleaseRequest) (*RollbackReleaseResponse, error)
	//RunReleaseTest(ctx context.Context, in *TestReleaseRequest) (*api.TestReleaseResponse, error)
}

//type Release interface {
//	CreateRelease(context.Context, *CreateReleaseRequest) (*CreateReleaseResponse, error)
//	GetReleaseStatus(context.Context, *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error)
//	GetReleaseContent(context.Context, *GetReleaseContentRequest) (*GetReleaseContentResponse, error)
//	ListReleases(context.Context, *ListReleasesRequest) (ListReleasesClient, error)
//	UpdateRelease(context.Context, *UpdateReleaseRequest) (*UpdateReleaseResponse, error)
//	UpgradeRelease(context.Context, *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error)
//	DeleteRelease(context.Context, *DeleteReleaseRequest) (*DeleteReleaseResponse, error)
//	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
//	RollbackRelease(context.Context, *RollbackReleaseRequest) (*RollbackReleaseResponse, error)
//}
