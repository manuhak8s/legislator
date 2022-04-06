package k8s

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaces creates a clientset based on the k8s packages
// Additionally this function reads all namespaces from the current kube context
// It returns a list of coreV1 namespaces
func GetNamespaces() (*corev1.NamespaceList, error){

	clientset, err := GetK8sClient()
	if err != nil {
		fmt.Println(err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	return namespaces, nil
}

// GetNamespaceNames accepts a list of coreV1 namespaces and returns all of their names
// The names are packed into a string slice
func GetNamespaceNames(namespaces *corev1.NamespaceList) []string {
	var nsNameArr []string
	for _, ns := range namespaces.Items {
		nsNameArr = append(nsNameArr, ns.Name)
	}

	return nsNameArr
}

// GetNamespace returns a coreV1 namespace instance based on the required namespace name
// This function compares alle namespaces with the given string value
// If no namespaces is found an empty namespace instance and error is returned
func GetNamespace(nsName string,namespaces *corev1.NamespaceList) (corev1.Namespace, error) {
	var emptyNamespaceInstance corev1.Namespace

	if len(namespaces.Items) < 1 {
		return emptyNamespaceInstance, fmt.Errorf("no namespaces found in current cluster context")
	}

	for _, ns := range namespaces.Items {
		if nsName == ns.Name {
			return ns, nil
		}
	}

	return emptyNamespaceInstance, fmt.Errorf("namespace not found")
}

// NamespaceExists checks if a namespace exists in the current kube context
// It accepts the string representation of a name and a coreV1 list of namespaces
// If an item matches a bool value true is returned, otherwise false
func NamespaceExists(nsName string, namespaces *corev1.NamespaceList) bool {
	nsNames := GetNamespaceNames(namespaces)
	if len(nsNames) < 1 {
		log.Println("no namespaces found")
		return false
	}

	for _, name := range nsNames {
		if nsName == name {
			return true
		}
	}

	return false
}