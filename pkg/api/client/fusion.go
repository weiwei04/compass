package client

import (
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	tillerapi "k8s.io/"
)

type CompassClient interface {
	CreateRelease()
	ListReleases(ns string, limit int64, offset int64, status int64 /*TODO int64 -> ?*/)
	UpdateRelease()
	UpgradeRelease()
	DeleteRelease()
	GetReleaseHistory()
	RollbackRelease()
	GetReleaseStatus()
	GetReleaseContent()
}
