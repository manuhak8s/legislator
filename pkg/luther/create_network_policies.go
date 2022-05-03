package luther

import (
	"context"
	"fmt"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_"k8s.io/client-go/kubernetes"
)



func InitV1NetworkPolicies(configPath string) ([]v1.NetworkPolicy, error) {
	var config config.Config
	var networkPolicies []v1.NetworkPolicy
	var deploymentOpts []DeploymentOpts

	kubernetesDefaultLabel := "kubernetes.io/metadata.name"
	configData, _ := config.ReadConfig(configPath)
	namespaces, _ := k8s.GetNamespaces()

	for _, set := range configData.ConnectedSets {
		for _, ns := range namespaces.Items {
			for k1, v1 := range set.TargetNamespaces.MatchLabels {
				for k2, v2 := range ns.Labels {
					if k1 == k2 && v1 == v2 {
						deploymentOpt := DeploymentOpts{
							Namespace: ns.Name,
							NamespaceLabels: ns.Labels,
							ConnectedSet: set,
						}

						deploymentOpts = append(deploymentOpts, deploymentOpt)
					}
				}
			}
		}
	}

	for _, opt := range deploymentOpts {
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

	return networkPolicies, nil
}

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

func ExecuteLegislation(configPath string) {
	err := DeployV1NetworkPolicies(configPath)
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveLabel(labels map[string]string, label string) map[string]string {
	delete(labels,label)
	return labels
}