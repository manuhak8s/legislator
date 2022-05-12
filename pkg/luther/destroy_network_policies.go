package luther

import (
	"context"
	"fmt"
	"strings"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DetectConnectedSetNetworkPolicies(configPath string) ([]string, []string, error) {
	var config config.Config
	var connectedSetNames []string
	var networkPolicyNames []string
	var targetNetworkPolicyNames []string
	var targetNamespaces []string

	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return nil, nil, err
	}

	connectedSets := configData.ConnectedSets

	for _, set := range connectedSets {
		connectedSetNames = append(connectedSetNames, set.Name)
	}

	if len(connectedSetNames) < 1 {
		return nil, nil, fmt.Errorf("no connected set names detected at config: %s", configPath)
	}

	namespaces, err := k8s.GetNamespaces()
	if err != nil {
		return nil, nil, err
	}

	for _, set := range configData.ConnectedSets {
		for _, ns := range namespaces.Items {
			for k1, v1 := range set.TargetNamespaces.MatchLabels {
				for k2, v2 := range ns.Labels {
					if k1 == k2 && v1 == v2 {
						targetNamespaces = append(targetNamespaces, ns.Name)
					}
				}
			}
		}
	}

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return nil, nil, err
	}

	for _, namespace := range targetNamespaces {
		networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return nil, nil, err
		}
	
		for _, networkPolicy := range networkPolicyList.Items {
			networkPolicyNames = append(networkPolicyNames, networkPolicy.Name)
		}
		
		targetNetworkPolicyNames, err = FilterNetworkPolicyNames(connectedSetNames, networkPolicyNames, namespace)
		if err != nil {
			return nil, nil, err
		}
	}

	return targetNetworkPolicyNames, k8s.GetNamespaceNames(namespaces), nil
}

func FilterNetworkPolicyNames(setNames []string, networkPolicyNames []string, namespace string,) ([]string, error){
	var filteredNetworkPolicyNames []string

	for _, networkPolicyName := range networkPolicyNames {
		for _, setName := range setNames {
			if strings.HasPrefix(networkPolicyName, setName) {
				filteredNetworkPolicyNames = append(filteredNetworkPolicyNames, networkPolicyName)
			}
		}
	}

	if len(filteredNetworkPolicyNames) < 1 {
		return nil, fmt.Errorf("no networkpolicies with matching connected sets found in current namespace: %s", namespace)
	}

	return filteredNetworkPolicyNames, nil
}

func DestroyConnectedSetNetworkPolicies(configPath string) error {
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	targetNetworkPolicyNames, namespaces, err := DetectConnectedSetNetworkPolicies(configPath)
	if err != nil {
		return err
	}

	for _, networkPolicy := range targetNetworkPolicyNames {
		for _, namespace := range namespaces{
			clientset.NetworkingV1().NetworkPolicies(namespace).Delete(context.Background(), networkPolicy, metav1.DeleteOptions{})
		}

	}

	return nil
}

func ExecuteDestruction(configPath string) {
	err := DestroyConnectedSetNetworkPolicies(configPath)
	if err != nil {
		fmt.Print(err)
	}
}

func DestroyAllNetworkPolicies() error {
	var networkPolicies *v1.NetworkPolicyList

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	namespaces, err := k8s.GetNamespaces()
	if err != nil {
		return err
	}

	for _, ns := range namespaces.Items {
		networkPolicies, err = clientset.NetworkingV1().NetworkPolicies(ns.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return err
		}

		for _, policy := range networkPolicies.Items {
			err = clientset.NetworkingV1().NetworkPolicies(ns.Name).Delete(context.Background(), policy.Name, metav1.DeleteOptions{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}