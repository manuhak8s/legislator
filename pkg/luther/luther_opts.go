package luther

import (
	"fmt"
	"reflect"

	"github.com/manuhak8s/legislator/pkg/config"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
)

// The LutherOpts struct contains fields for a more efficient creation of
// k8s network policies based of the networking v1 package. A single instance
// the namespacename and its labels with the ConnectedSets and their network polices.
type LutherOpts struct {
	Namespace string
	NamespaceLabels map[string]string
	ConnectedSet config.ConnectedSet
	NetworkPolicies []v1.NetworkPolicy
}

// InitLutherOpts initialized the LutherOpts instances based on the given config and the
// accepts additionally a corev1 namespace list to pair them. If the labeling equality is given, 
// a LutherOpts struct will be initialized and appended to a slice which will be returned.
func InitLutherOpts(connectedSets config.ConnectedSets, namespaces *corev1.NamespaceList) ([]LutherOpts, error){
	var lutherOpts []LutherOpts

	for _, set := range connectedSets {
		for _, ns := range namespaces.Items {
			for k1, v1 := range set.TargetNamespaces.MatchLabels {
				for k2, v2 := range ns.Labels {
					if k1 == k2 && v1 == v2 {
						lutherOpt := LutherOpts{
							Namespace: ns.Name,
							NamespaceLabels: ns.Labels,
							ConnectedSet: set,
						}

						lutherOpts = append(lutherOpts, lutherOpt)
					}
				}
			}
		}
	}

	err := validateLutherOpts(lutherOpts)
	if err != nil {
		return nil, err
	}

	return lutherOpts, nil
}

// validateLutherOpts contains a validation process applies it to a given slice of LutherOpts:
// - minimal length has to be one option
// - target namespace with labels has to be set
// - connected set definition has to be complete without empty fields
// This functions ensures a correct further course after the initialization of the LutherOpts.
func validateLutherOpts(opts []LutherOpts) error{
	if len(opts) < 1 {
		return fmt.Errorf("no resources found at current kube context with matching labeling: cannot create network policies based on defined connected sets, for further information read the legislator instructions")
	}

	for _, opt := range opts {
		if opt.Namespace == "" || opt.NamespaceLabels == nil{
			return fmt.Errorf("no namespace data defined in lutherOpts")
		} 
		
		reflectedValue := reflect.ValueOf(opt.ConnectedSet)
		structValues := make([]interface{}, reflectedValue.NumField())

		for i := 0; i < reflectedValue.NumField(); i++ {
			structValues[i] = reflectedValue.Field(i).Interface()
		}

		for _, value := range structValues {
			if value == nil {
				return fmt.Errorf("empty field found in lutherOpts.ConnectedSet: %v", value)
			}
		}
		
	}
	return nil
}