package spec

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v3"

	sp "github.com/aquasecurity/defsec/pkg/spec"
	"github.com/aquasecurity/gitscan/pkg/types"
)

type Severity string

// ComplianceSpec represent the compliance specification
type ComplianceSpec struct {
	Spec Spec `yaml:"spec"`
}

type Spec struct {
	ID               string    `yaml:"id"`
	Title            string    `yaml:"title"`
	Description      string    `yaml:"description"`
	Version          string    `yaml:"version"`
	RelatedResources []string  `yaml:"relatedResources"`
	Controls         []Control `yaml:"controls"`
}

// Control represent the cps controls data and mapping checks
type Control struct {
	ID            string        `yaml:"id"`
	Name          string        `yaml:"name"`
	Description   string        `yaml:"description,omitempty"`
	Checks        []SpecCheck   `yaml:"checks"`
	Severity      Severity      `yaml:"severity"`
	DefaultStatus ControlStatus `yaml:"defaultStatus,omitempty"`
}

// SpecCheck represent the scanner who perform the control check
type SpecCheck struct {
	ID string `yaml:"id"`
}

// ControlCheck provides the result of conducting a single audit step.
type ControlCheck struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description,omitempty"`
	PassTotal   int      `yaml:"passTotal"`
	FailTotal   int      `yaml:"failTotal"`
	Severity    Severity `yaml:"severity"`
}

type ControlStatus string

const (
	FailStatus ControlStatus = "FAIL"
	PassStatus ControlStatus = "PASS"
	WarnStatus ControlStatus = "WARN"
)

// Scanners reads spec control and determines the scanners by check ID prefix
func (cs *ComplianceSpec) Scanners() (types.Scanners, error) {
	scannerTypes := map[types.Scanner]struct{}{}
	for _, control := range cs.Spec.Controls {
		for _, check := range control.Checks {
			scannerType := scannerByCheckID(check.ID)
			if scannerType == types.UnknownScanner {
				return nil, xerrors.Errorf("unsupported check ID: %s", check.ID)
			}
			scannerTypes[scannerType] = struct{}{}
		}
	}
	return maps.Keys(scannerTypes), nil
}

// CheckIDs return list of compliance check IDs
func (cs *ComplianceSpec) CheckIDs() map[types.Scanner][]string {
	checkIDsMap := map[types.Scanner][]string{}
	for _, control := range cs.Spec.Controls {
		for _, check := range control.Checks {
			scannerType := scannerByCheckID(check.ID)
			checkIDsMap[scannerType] = append(checkIDsMap[scannerType], check.ID)
		}
	}
	return checkIDsMap
}

func scannerByCheckID(checkID string) types.Scanner {
	checkID = strings.ToLower(checkID)
	switch {
	case strings.HasPrefix(checkID, "cve-") || strings.HasPrefix(checkID, "dla-"):
		return types.VulnerabilityScanner
	case strings.HasPrefix(checkID, "avd-"):
		return types.MisconfigScanner
	default:
		return types.UnknownScanner
	}
}

// GetComplianceSpec accepct compliance flag name/path and return builtin or file system loaded spec
func GetComplianceSpec(specNameOrPath string) (ComplianceSpec, error) {
	var b []byte
	var err error
	if strings.HasPrefix(specNameOrPath, "@") {
		b, err = os.ReadFile(strings.TrimPrefix(specNameOrPath, "@"))
		if err != nil {
			return ComplianceSpec{}, fmt.Errorf("error retrieving compliance spec from path: %w", err)
		}
	} else {
		// TODO: GetSpecByName() should return []byte
		b = []byte(sp.NewSpecLoader().GetSpecByName(specNameOrPath))
	}

	var complianceSpec ComplianceSpec
	if err = yaml.Unmarshal(b, &complianceSpec); err != nil {
		return ComplianceSpec{}, xerrors.Errorf("spec yaml decode error: %w", err)
	}
	return complianceSpec, nil

}