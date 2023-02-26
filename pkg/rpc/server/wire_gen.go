// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/aquasecurity/gitscan/pkg/detector/ospkg"
	"github.com/aquasecurity/gitscan/pkg/fanal/applier"
	"github.com/aquasecurity/gitscan/pkg/fanal/cache"
	"github.com/aquasecurity/gitscan/pkg/scanner/local"
	"github.com/aquasecurity/gitscan/pkg/vulnerability"
)

// Injectors from inject.go:

func initializeScanServer(localArtifactCache cache.LocalArtifactCache) *ScanServer {
	applierApplier := applier.NewApplier(localArtifactCache)
	detector := ospkg.Detector{}
	config := db.Config{}
	client := vulnerability.NewClient(config)
	scanner := local.NewScanner(applierApplier, detector, client)
	scanServer := NewScanServer(scanner)
	return scanServer
}
