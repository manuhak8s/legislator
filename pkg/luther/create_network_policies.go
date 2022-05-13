package luther

import (
	"context"
	"fmt"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExecuteLegislation is called by the apply command an executes the
// process for creating connected sets based on the given config path
func ExecuteLegislation(configPath string) {
	err := DeployV1NetworkPolicies(configPath)
	if err != nil {
		fmt.Println(err)
	}
}

// DeployV1NetworkPolicies creates a kubernetes clientset and deploys the 
// network policies based on their initialization with the networking v1 interface
func DeployV1NetworkPolicies(configPath string) error {
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	networkPolicies, err := InitV1NetworkPolicies(configPath)
	if err != nil {
		return err
	}

	for _, policy := range networkPolicies {
		_, err = clientset.NetworkingV1().NetworkPolicies(policy.Namespace).Create(context.Background(), &policy, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// InitV1NetworkPolicies initialized network policies based on the given config. 
// It creates the instance of the config and transfers the deployment information
// to the creation process of the network policies.
func InitV1NetworkPolicies(configPath string) ([]v1.NetworkPolicy, error) {
	var config config.Config

	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return nil, err
	}

	namespaces, err := k8s.GetNamespaces()
	if err != nil {
		return nil, err
	}
	
	deploymentOpts, err := InitLutherOpts(configData.ConnectedSets, namespaces)
	if err != nil {
		return nil, err
	}

	networkPolicies, err := CreateV1NetworkPolicies(deploymentOpts)
	if err != nil {
		return nil, err
	}

	return networkPolicies, nil
}

// CreateV1NetworkPolicies creates network policies based on the v1 networking package
// and given deployment information.
// It returns a list of v1.NetworkPolicy.
func CreateV1NetworkPolicies(lutherOpts []LutherOpts) ([]v1.NetworkPolicy, error) {
	var networkPolicies []v1.NetworkPolicy
	kubernetesDefaultLabel := "kubernetes.io/metadata.name"

	for _, opt := range lutherOpts {
		opt.NamespaceLabels = RemoveLabel(opt.NamespaceLabels, kubernetesDefaultLabel)
		
		podSelector := metav1.LabelSelector{
			MatchLabels: opt.ConnectedSet.PodSelector.MatchLabels,
		}

		v1NetworkPolicy := &v1.NetworkPolicy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "NetworkPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: GenerateNetworkPolicyName(opt.ConnectedSet.Name),
				Namespace: opt.Namespace,
				Labels: opt.NamespaceLabels,
			},
			Spec: v1.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{
					MatchLabels: opt.ConnectedSet.PodSelector.MatchLabels,
				},
				PolicyTypes: []v1.PolicyType{
					"Ingress",
				},
				Ingress: []v1.NetworkPolicyIngressRule{
					v1.NetworkPolicyIngressRule{
						From: []v1.NetworkPolicyPeer{
							v1.NetworkPolicyPeer{
								PodSelector: &podSelector,
							},
						},
					},
				},
			},
		}

		networkPolicies = append(networkPolicies, *v1NetworkPolicy)
	}

	if len(networkPolicies) < 1 {
		return nil, fmt.Errorf("error while creating network policy instances")
	}

	return networkPolicies, nil
}