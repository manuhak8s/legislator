package cmd

import (
	"fmt"
	"context"
	"github.com/spf13/cobra"
	"github.com/manuhak8s/legislator/pkg/k8s"
	"github.com/manuhak8s/legislator/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var namespaceCmd = &cobra.Command{
	Use: "namespaces",
	Short: "Get namespaces from kubernetes cluster",
	Long: "Get namespaces from kubernetes cluster",
	Run: getNamespaceData,
}

func getNamespaceData(cmd *cobra.Command, args []string) {
	log.LogNamespaceReading()
	fmt.Println("")

	//clientset, err := k8s.GetK8sClient()
	clientset, err := k8s.GetK8sDefaultClient()
	if err != nil {
		fmt.Println(err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, ns := range namespaces.Items {
		fmt.Println(ns.Name)
	}
}