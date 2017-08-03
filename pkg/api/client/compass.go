package client

import (
	context "golang.org/x/net/context"
)

// runtime release manage api

type CompassClient interface {
	Connect() error
	Shutdown()
	CreateRelease(context.Context, *CreateReleaseRequest) (*CreateReleaseResponse, error)
	GetReleaseStatus(context.Context, *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error)
	GetReleaseContent(context.Context, *GetReleaseContentRequest) (*GetReleaseContentResponse, error)
	ListReleases(context.Context, *ListReleasesRequest) (ListReleasesClient, error)
	UpdateRelease(context.Context, *UpdateReleaseRequest) (*UpdateReleaseResponse, error)
	UpgradeRelease(context.Context, *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error)
	DeleteRelease(context.Context, *DeleteReleaseRequest) (*DeleteReleaseResponse, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	RollbackRelease(context.Context, *RollbackReleaseRequest) (*RollbackReleaseResponse, error)
}
