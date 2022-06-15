package k8s

import (
	"context"
	"fmt"

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


