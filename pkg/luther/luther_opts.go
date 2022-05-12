package luther

import (
	"github.com/manuhak8s/legislator/pkg/config"
	v1 "k8s.io/api/networking/v1"
)

type LutherOpts struct {
	Namespace string
	NamespaceLabels map[string]string
	ConnectedSet config.ConnectedSet
	NetworkPolicies []v1.NetworkPolicy
}