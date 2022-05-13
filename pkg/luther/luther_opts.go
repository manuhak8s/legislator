package luther

import (
	"fmt"
	"reflect"

	"github.com/manuhak8s/legislator/pkg/config"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
)

type LutherOpts struct {
	Namespace string
	NamespaceLabels map[string]string
	ConnectedSet config.ConnectedSet
	NetworkPolicies []v1.NetworkPolicy
}

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

	err := ValidateLutherOpts(lutherOpts)
	if err != nil {
		return nil, err
	}

	return lutherOpts, nil
}

func ValidateLutherOpts(opts []LutherOpts) error{
	if len(opts) < 1 {
		return fmt.Errorf("no items initialized in lutherOpts: empty slice")
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