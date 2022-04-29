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

	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return nil, err
	}

	connectedSets := configData.ConnectedSets

	for _, set := range connectedSets {
		podSelector := metav1.LabelSelector{
			MatchLabels: set.PodSelector.MatchLabels,
		}

		v1NetworkPolicy := &v1.NetworkPolicy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind: "NetworkPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: generateNetworkPolicyName(set.Name),
				Namespace: "namespace-1",
				Labels: set.PodSelector.MatchLabels,
			},
			Spec: v1.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{
					MatchLabels: set.PodSelector.MatchLabels,
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

func DeployV1NetworkPolicies(namespace string, configPath string) error {
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	networkPolicies, err := InitV1NetworkPolicies(configPath)
	if err != nil {
		return err
	}

	for _, policy := range networkPolicies {
		_, err = clientset.NetworkingV1().NetworkPolicies(namespace).Create(context.Background(), &policy, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteLegislation() {
	namespace := "namespace-1"
	configPath := "/Users/manuelhaugg/legislator/test_data/configs/v2_data.yaml"
	err := DeployV1NetworkPolicies(namespace, configPath)
	if err != nil {
		fmt.Println(err)
	}
}

/*func DeployNetworkPolicies(configPath string) error{
	var config config.Config

	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return err
	}

	networkPolicies, err := GetV1NetworkPolicies(configData.ConnectedSets)
	if err != nil {
		return err
	}

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	for _, policy := range networkPolicies{
		CreateNetworkPolicy(clientset, &policy)
	}

	return nil
}

func GetV1NetworkPolicies(sets config.ConnectedSets) ([]v1.NetworkPolicy, error){
	var nwPolicies []v1.NetworkPolicy

	for _, set := range sets {
		podSelector := metav1.LabelSelector{
			MatchLabels: set.PodSelector.MatchLabels,
		}

		nwPolicy := &v1.NetworkPolicy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "NetworkPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: generateNetworkPolicyName(set.Name),
				Namespace: "namespace-1",
				Labels:    set.PodSelector.MatchLabels,
	
			},
			Spec: v1.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{
					MatchLabels: set.PodSelector.MatchLabels,
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

		nwPolicies = append(nwPolicies, *nwPolicy)
	}

	if len(nwPolicies)<1 {
		return nil, fmt.Errorf("error while creating v1 network policies: no translation occured")
	}

	return nwPolicies, nil
}

func CreateNetworkPolicy(clientSet *kubernetes.Clientset, nwPolicy *v1.NetworkPolicy) error {
	_, err := clientSet.NetworkingV1().NetworkPolicies("namespace-1").Create(context.Background(), nwPolicy, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	
	return nil
}*/