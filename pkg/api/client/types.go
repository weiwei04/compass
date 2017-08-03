package client

import (
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	tillerapi "k8s.io/helm/pkg/proto/hapi/services"
)

// all the ugly things

type CreateReleaseRequest compassapi.CreateCompassReleaseRequest
type CreateReleaseResponse compassapi.CreateCompassReleaseResponse

type ListReleasesRequest tillerapi.ListReleasesRequest
type ListReleasesClient tillerapi.ReleaseService_ListReleasesClient

type GetReleaseStatusRequest tillerapi.GetReleaseStatusRequest
type GetReleaseStatusResponse tillerapi.GetReleaseStatusResponse

type GetReleaseContentRequest tillerapi.GetReleaseContentRequest
type GetReleaseContentResponse tillerapi.GetReleaseContentResponse

type UpdateReleaseRequest tillerapi.UpdateReleaseRequest
type UpdateReleaseResponse tillerapi.UpdateReleaseResponse

type UpgradeReleaseRequest tillerapi.UpdateReleaseRequest
type UpgradeReleaseResponse tillerapi.UpdateReleaseResponse

type DeleteReleaseRequest tillerapi.UninstallReleaseRequest
type DeleteReleaseResponse tillerapi.UninstallReleaseResponse

type RollbackReleaseRequest tillerapi.RollbackReleaseRequest
type RollbackReleaseResponse tillerapi.RollbackReleaseResponse

type GetHistoryRequest tillerapi.GetHistoryRequest
type GetHistoryResponse tillerapi.GetHistoryResponse

//type CompassClient interface {
//	CreateRelease()
//	ListReleases()
//	UpdateRelease()
//	UpgradeRelease()
//	DeleteRelease()
//	GetReleaseHistory()
//	RollbackRelease()
//	GetReleaseStatus()
//	GetReleaseContent()
//}
