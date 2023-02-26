package commands

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/gitscan-kubernetes/pkg/k8s"
	"github.com/aquasecurity/gitscan-kubernetes/pkg/gitscank8s"
	"github.com/aquasecurity/gitscan/pkg/flag"
	"github.com/aquasecurity/gitscan/pkg/log"
)

// namespaceRun runs scan on kubernetes cluster
func namespaceRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}

	gitscank8s := gitscank8s.New(cluster, log.Logger).Namespace(getNamespace(opts, cluster.GetCurrentNamespace()))

	artifacts, err := gitscank8s.ListArtifacts(ctx)
	if err != nil {
		return xerrors.Errorf("get k8s artifacts error: %w", err)
	}

	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}

func getNamespace(opts flag.Options, currentNamespace string) string {
	if len(opts.K8sOptions.Namespace) > 0 {
		return opts.K8sOptions.Namespace
	}

	return currentNamespace
}
