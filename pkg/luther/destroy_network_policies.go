package luther

import (
	"context"
	"fmt"
	"strings"

	"github.com/manuhak8s/legislator/pkg/config"
	"github.com/manuhak8s/legislator/pkg/k8s"
	
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ExecuteDestruction is called by the destroy command an executes the
// process for removing network policies based on the given config file
// and additionally creates a kubernetes clientset.
func ExecuteDestruction(configPath string) {
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		fmt.Print(err)
	}

	err = DestroyConnectedSetNetworkPolicies(configPath, clientset)
	if err != nil {
		fmt.Print(err)
	}
}

// DestroyConnectedSetNetworkPolicies detects and destroys the target network policies 
// based on the given config file with a given clientset.
func DestroyConnectedSetNetworkPolicies(configPath string, clientset *kubernetes.Clientset) error {
	targetNetworkPolicyNames, namespaces, err := DetectConnectedSetTargets(configPath, clientset)
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

// DetectConnectedSetTargets detects the target network policies and namespaces 
// based on the given config file with the defined connected sets.
func DetectConnectedSetTargets(configPath string, clientset *kubernetes.Clientset) ([]string, []string, error) {
	var config config.Config
	configData, err := config.ReadConfig(configPath)
	if err != nil {
		return nil, nil, err
	}

	namespaces, err := k8s.GetNamespaces()
	if err != nil {
		return nil, nil, err
	}

	targetNamespaces, err := GetTargetNamespaceNames(configData.ConnectedSets, namespaces)
	if err != nil {
		return nil, nil, err
	}

	targetNetworkPolicyNames, err := GetTargetNetworkPolicyNames(targetNamespaces, clientset, configData)
	if err != nil {
		return nil, nil, err
	}

	return targetNetworkPolicyNames, k8s.GetNamespaceNames(namespaces), nil
}

// FilterNetworkPolicyNames filters network policies with the matching naming convention of a connected set
// inside a namespace and returns a list of the results.
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


// GetTargetNamespaceNames detects target namespaces by matching equal labels of a connected set targetNamespace field
// labels and namespace labels. It returns a string list of equalitites.
func GetTargetNamespaceNames(connectedSets config.ConnectedSets, namespaces *corev1.NamespaceList) ([]string, error){
	var targetNamespaces []string
	for _, set := range connectedSets {
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

	if len(targetNamespaces) < 1 {
		return nil, fmt.Errorf("error while filtering target namespace names")
	}

	return targetNamespaces, nil
}


// GetTargetNetworkPolicyNames detects target network policies from all namespaces by using a clientset
// and filters by facing them against the defined connected sets from the given config file.
func GetTargetNetworkPolicyNames(namespaceNames []string, clientset *kubernetes.Clientset, configData *config.Config) ([]string, error) {
	var targetNetworkPolicyNames []string
	var networkPolicyNames []string
	var connectedSetNames []string

	for _, set := range configData.ConnectedSets {
		connectedSetNames = append(connectedSetNames, set.Name)
	}
	
	for _, namespace := range namespaceNames {
		networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
	
		for _, networkPolicy := range networkPolicyList.Items {
			networkPolicyNames = append(networkPolicyNames, networkPolicy.Name)
		}
		
		targetNetworkPolicyNames, err = FilterNetworkPolicyNames(connectedSetNames, networkPolicyNames, namespace)
		if err != nil {
			return nil, err
		}
	}

	return targetNetworkPolicyNames, nil
}

// DestroyAllNetworkPolicies removes all network policies from a kubernetes cluster.
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