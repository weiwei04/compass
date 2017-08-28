package compass

import (
	hapi_release "k8s.io/helm/pkg/proto/hapi/release"
)

func CompassReleaseInfo(in *hapi_release.Info) *Info {
	status := in.GetStatus()
	out := &Info{
		//FirstDeployed: in.GetFirstDeployed(),
		//LastDeployed:  in.GetLastDeployed(),
		//Deleted:       in.GetDeleted(),
		Status: &Status{
			Code:      int32(status.GetCode()),
			Resources: status.GetResources(),
			Notes:     status.GetNotes(),
		},
	}
	if in.GetFirstDeployed() != nil {
		out.FirstDeployed = in.FirstDeployed.Seconds
	}
	if in.GetLastDeployed() != nil {
		out.LastDeployed = in.LastDeployed.Seconds
	}
	if in.GetDeleted() != nil {
		out.Deleted = in.Deleted.Seconds
	}
	return out
}

func CompassRelease(in *hapi_release.Release) *Release {
	out := &Release{
		Name:      in.Name,
		Info:      CompassReleaseInfo(in.GetInfo()),
		Chart:     in.Chart,
		Config:    in.Config,
		Version:   in.Version,
		Namespace: in.Namespace,
	}
	// strip un marshalable google.prototbuf.Any files
	if out.Chart != nil {
		out.Chart.Files = nil
	}
	return out
}

func CompassReleaseSlice(in []*hapi_release.Release) []*Release {
	out := make([]*Release, len(in), len(in))
	for i := range in {
		out[i] = CompassRelease(in[i])
	}
	return out
}
