package commands

import (
	"context"

	"github.com/aquasecurity/gitscan-kubernetes/pkg/k8s"
	"github.com/aquasecurity/gitscan-kubernetes/pkg/gitscank8s"
	"github.com/aquasecurity/gitscan/pkg/flag"
	"github.com/aquasecurity/gitscan/pkg/log"

	"golang.org/x/xerrors"
)

// clusterRun runs scan on kubernetes cluster
func clusterRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}

	artifacts, err := gitscank8s.New(cluster, log.Logger).ListArtifactAndNodeInfo(ctx)
	if err != nil {
		return xerrors.Errorf("get k8s artifacts error: %w", err)
	}

	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}
