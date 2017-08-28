package compass

import (
	"fmt"

	hapi_chart "k8s.io/helm/pkg/proto/hapi/chart"
	hapi_release "k8s.io/helm/pkg/proto/hapi/release"
)

func NewFakeCreateReleaseResponse(req *CreateReleaseRequest) *CreateReleaseResponse {
	return &CreateReleaseResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
			Chart: &hapi_chart.Chart{
				Metadata: &hapi_chart.Metadata{
					Name: req.Chart,
				},
			},
			Config: req.Values,
		}),
	}
}

func NewFakeDeleteReleaseResponse(req *DeleteReleaseRequest) *DeleteReleaseResponse {
	return &DeleteReleaseResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
		}),
	}
}

func NewFakeUpdateReleaseResponse(req *UpdateReleaseRequest) *UpdateReleaseResponse {
	return &UpdateReleaseResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
			Config:    req.Values,
		}),
	}
}

func NewFakeUpgradeReleaseResponse(req *UpgradeReleaseRequest) *UpgradeReleaseResponse {
	return &UpgradeReleaseResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
			Chart: &hapi_chart.Chart{
				Metadata: &hapi_chart.Metadata{
					Name: req.Chart,
				},
			},
			Config: req.Values,
		}),
	}
}

func NewFakeListReleasesResponse(req *ListReleasesRequest) *ListReleasesResponse {
	releases := []*Release{}
	for i := int64(0); i < req.Limit; i++ {
		releases = append(releases, CompassRelease(&hapi_release.Release{
			Name:      fmt.Sprintf("release-%d", i),
			Namespace: req.Namespace,
			Chart: &hapi_chart.Chart{
				Metadata: &hapi_chart.Metadata{
					Name: "myspace/mychart:v0.1.0",
				},
			},
		}))
	}
	return &ListReleasesResponse{
		Count:    req.Limit,
		Next:     req.Offset,
		Total:    req.Limit,
		Releases: releases,
	}
}

func NewFakeGetReleaseStatusResponse(req *GetReleaseStatusRequest) *GetReleaseStatusResponse {
	return &GetReleaseStatusResponse{
		Name:      req.Name,
		Namespace: req.Namespace,
	}
}

func NewFakeGetReleaseContentResponse(req *GetReleaseContentRequest) *GetReleaseContentResponse {
	return &GetReleaseContentResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
			Version:   req.Version,
		}),
	}
}

func NewFakeGetReleaseHistoryResponse(req *GetReleaseHistoryRequest) *GetReleaseHistoryResponse {
	releases := []*Release{}
	for i := int32(0); i < req.Max; i++ {
		releases = append(releases, CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
		}))
	}
	return &GetReleaseHistoryResponse{
		Releases: releases,
	}
}

func NewFakeRollbackReleaseResponse(req *RollbackReleaseRequest) *RollbackReleaseResponse {
	return &RollbackReleaseResponse{
		Release: CompassRelease(&hapi_release.Release{
			Name:      req.Name,
			Namespace: req.Namespace,
		}),
	}
}

func NewFakeTestReleaseResponse(req *TestReleaseRequest) *TestReleaseResponse {
	return &TestReleaseResponse{
		Msg:    "ok",
		Status: hapi_release.TestRun_SUCCESS,
	}
}
