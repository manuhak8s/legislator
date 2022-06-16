package luther

import (
	"context"
	"fmt"
	_"encoding/json"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	"github.com/manuhak8s/legislator/pkg/logger"
	
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ExecuteLegislation is called by the apply command an executes the
// process for creating connected sets based on the given config path
// and creates a kubernetes clientset.
func ExecuteLegislation(configPath string) {
	logger.TriggerOutput("loading", "executing legislation of config file: " + configPath)
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		logger.TriggerOutput("fail", err.Error())
	}

	err = DeployV1NetworkPolicies(configPath, clientset)
	if err != nil {
		logger.TriggerOutput("fail", err.Error())
	}
	
	if err == nil {
		logger.TriggerOutput("success", "connected sets based on config file successfully deployed as network policies")
	}
}

// DeployV1NetworkPolicies deploys the network policies based on their initialization
// with the networking v1 interface by using the kubernetes clientset.
func DeployV1NetworkPolicies(configPath string, clientset *kubernetes.Clientset) error {
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

	networkPolicies, err := CreateV1NetworkPolicies(deploymentOpts, configPath)
	if err != nil {
		return nil, err
	}

	return networkPolicies, nil
}

// CreateV1NetworkPolicies creates network policies based on the v1 networking package
// and given deployment information.
// It returns a list of v1.NetworkPolicy.
func CreateV1NetworkPolicies(lutherOpts []LutherOpts, configPath string) ([]v1.NetworkPolicy, error) {
	logger.TriggerOutput("loading", "... creating network policies ...")
	var networkPolicies []v1.NetworkPolicy
	kubernetesDefaultLabel := "kubernetes.io/metadata.name"

	for _, opt := range lutherOpts {
		opt.NamespaceLabels = RemoveLabel(opt.NamespaceLabels, kubernetesDefaultLabel)
		networkPolicyPeers := initV1NetworkPolicyPeers(opt)

		v1NetworkPolicy := &v1.NetworkPolicy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "NetworkPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: GenerateNetworkPolicyName(configPath, opt.ConnectedSet.Name),
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
				Ingress: []v1.NetworkPolicyIngressRule{},
			},
		}

		networkPolicyIngressRule := v1.NetworkPolicyIngressRule{
			From: []v1.NetworkPolicyPeer{},
		}
		networkPolicyIngressRule.From = append(networkPolicyIngressRule.From, networkPolicyPeers...)

		v1NetworkPolicy.Spec.Ingress = append(v1NetworkPolicy.Spec.Ingress, networkPolicyIngressRule)

		networkPolicies = append(networkPolicies, *v1NetworkPolicy)
	}

	if len(networkPolicies) < 1 {
		return nil, fmt.Errorf("error while creating network policy instances")
	}

	return networkPolicies, nil
}

func initV1NetworkPolicyPeers(option LutherOpts) []v1.NetworkPolicyPeer {
	var networkPolicyPeers []v1.NetworkPolicyPeer
	
	defaultPodSelector := metav1.LabelSelector{
		MatchLabels: option.ConnectedSet.PodSelector.MatchLabels,
	}

	defaultNetworkPolicyPeer := v1.NetworkPolicyPeer {
		PodSelector: &defaultPodSelector,
	}
	networkPolicyPeers = append(networkPolicyPeers, defaultNetworkPolicyPeer)


	if len(option.ConnectedSet.TargetNamespaces.MatchLabels) > 1 {
		for key, value := range option.ConnectedSet.TargetNamespaces.MatchLabels {
			for nsKey, nsValue := range option.NamespaceLabels{
				if key != nsKey && value != nsValue {
					podSelector := metav1.LabelSelector{
						MatchLabels: option.ConnectedSet.PodSelector.MatchLabels,
					}	
					namespaceSelector := metav1.LabelSelector {
						MatchLabels: map[string]string{
							key:value,					
						},
					}
				
					networkPolicyPeer := v1.NetworkPolicyPeer {
						PodSelector: &podSelector,
						NamespaceSelector: &namespaceSelector,
					}
				
					networkPolicyPeers = append(networkPolicyPeers, networkPolicyPeer)	
				}
			}		
		}
	}

	return networkPolicyPeers
}