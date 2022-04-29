package luther

import (
	"context"
	"fmt"
	"strings"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DetectConnectedSetNetworkPolicies(configPath string, namespace string) ([]string, error) {
	var config config.Config
	var connectedSetNames []string
	var networkPolicyNames []string

	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return nil, err
	}

	connectedSets := configData.ConnectedSets

	for _, set := range connectedSets {
		connectedSetNames = append(connectedSetNames, set.Name)
	}

	if len(connectedSetNames) < 1 {
		return nil, fmt.Errorf("no connected set names detected at config: %s", configPath)
	}

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return nil, err
	}

	networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(networkPolicyList.Items) < 1 {
		return nil, fmt.Errorf("no network policies found hat namespace: %s", namespace)
	}

	for _, networkPolicy := range networkPolicyList.Items {
		networkPolicyNames = append(networkPolicyNames, networkPolicy.Name)
	}
	
	targetNetworkPolicyNames, err := FilterNetworkPolicyNames(connectedSetNames, networkPolicyNames, namespace)
	if err != nil {
		return nil, err
	}

	return targetNetworkPolicyNames, nil
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

func DestroyConnectedSetNetworkPolicies(namespace string, configPath string) error {
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		return err
	}

	targetNetworkPolicyNames, err := DetectConnectedSetNetworkPolicies(configPath, namespace)
	if err != nil {
		return err
	}

	for _, networkPolicy := range targetNetworkPolicyNames {
		err = clientset.NetworkingV1().NetworkPolicies(namespace).Delete(context.Background(), networkPolicy, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteDestruction() {
	namespace := "namespace-1"
	configPath := "/Users/manuelhaugg/legislator/test_data/configs/v2_data.yaml"
	err := DestroyConnectedSetNetworkPolicies(namespace, configPath)
	if err != nil {
		fmt.Print(err)
	}
}