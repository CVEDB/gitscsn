package commands

import (
	"context"
	"strings"

	"github.com/aquasecurity/gitscan/pkg/flag"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/gitscan-kubernetes/pkg/artifacts"
	"github.com/aquasecurity/gitscan-kubernetes/pkg/k8s"
	"github.com/aquasecurity/gitscan-kubernetes/pkg/gitscank8s"
	"github.com/aquasecurity/gitscan/pkg/log"
)

// resourceRun runs scan on kubernetes cluster
func resourceRun(ctx context.Context, args []string, opts flag.Options, cluster k8s.Cluster) error {
	kind, name, err := extractKindAndName(args)
	if err != nil {
		return err
	}

	gitscank8s := gitscank8s.New(cluster, log.Logger).Namespace(getNamespace(opts, cluster.GetCurrentNamespace()))
	runner := newRunner(opts, cluster.GetCurrentContext())

	if len(name) == 0 { // pods or configmaps etc
		if err = validateReportArguments(opts); err != nil {
			return err
		}

		targets, err := gitscank8s.Resources(kind).ListArtifacts(ctx)
		if err != nil {
			return err
		}

		return runner.run(ctx, targets)
	}

	// pod/NAME or pod NAME etc
	artifact, err := gitscank8s.GetArtifact(ctx, kind, name)
	if err != nil {
		return err
	}

	return runner.run(ctx, []*artifacts.Artifact{artifact})
}

func extractKindAndName(args []string) (string, string, error) {
	switch len(args) {
	case 1:
		s := strings.Split(args[0], "/")
		if len(s) != 2 {
			return args[0], "", nil
		}

		return s[0], s[1], nil
	case 2:
		return args[0], args[1], nil
	}

	return "", "", xerrors.Errorf("can't parse arguments %v. Please run `gitscan k8s` for usage.", args)
}
