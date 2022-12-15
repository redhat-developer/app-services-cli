package clustercmdutil

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"
)

// annotations are mapped to k8s labels, check that it's not used to set any reserved domain labels
var reservedDomains = []string{"kubernetes.io/", "k8s.io/", "openshift.io/"}

// BuildAnnotationsMap accepts a comma-separated list of string annotations in `key=value` format
// and returns the annotations as a map
func BuildAnnotationsMap(configs []string) (map[string]string, error) {

	annotations := make(map[string]string)

	for _, configPair := range configs {
		parts := strings.Split(configPair, "=")

		keyErrs := validation.IsQualifiedName(parts[0])
		if len(keyErrs) != 0 {
			return nil, fmt.Errorf("invalid annotation key %s: %s", parts[0], strings.Join(keyErrs, "; "))
		}

		for _, d := range reservedDomains {
			if strings.Contains(parts[0], d) {
				return nil, fmt.Errorf("cannot use reserved annotation %s from domain %s", parts[0], d)
			}
		}

		if len(parts) == 2 {
			errs := validation.IsValidLabelValue(parts[1])
			if len(errs) != 0 {
				return nil, fmt.Errorf("invalid annotation value %s: %s", parts[1], strings.Join(errs, "; "))
			}

			annotations[parts[0]] = parts[1]
		} else {
			annotations[parts[0]] = ""
		}

	}

	return annotations, nil

}
