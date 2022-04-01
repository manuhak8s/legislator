package cmd

import (
	"context"
	"fmt"

	"github.com/manuhak8s/legislator/pkg/k8s"
	"github.com/manuhak8s/legislator/pkg/log"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var listPolicyCmd = &cobra.Command{
	Use: "list_policies",
	Short: "list network policies from current namespace",
	Long: "list network policies from current namespace",
	Run: listPolicies,
}

func listPolicies(cmd *cobra.Command, args []string) {
	log.LogNetworkPolicyReading()

	//clientset, err := k8s.GetK8sClient()
	clientset, err := k8s.GetK8sClient()
	if err != nil {
		fmt.Println(err)
	}

	networkPolicies, err := clientset.NetworkingV1().NetworkPolicies("test-namespace-1").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, ns := range networkPolicies.Items {
		fmt.Println(ns.Name)
	}
}

var deletePolicyCmd = &cobra.Command{
	Use: "delete_policy",
	Short: "remove network policy from current namespace",
	Long: "remove network policy from current namespace",
	Run: deletePolicy,
}

func deletePolicy(cmd *cobra.Command, args []string) {
	log.LogNetworkPolicyRemoving()

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		fmt.Println(err)
	}

	err = clientset.NetworkingV1().NetworkPolicies("test-namespace-1").Delete(context.Background(), "default-deny-all", metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("- network policy deleted -")
	}
}

var createPolicyCmd = &cobra.Command{
	Use: "create_policy",
	Short: "create network policy into current namespace",
	Long: "create network policy into current namespace",
	Run: createPolicy,
}

func createPolicy(cmd *cobra.Command, args []string) {
	log.LogNetworkPolicyCreating()

	clientset, err := k8s.GetK8sClient()
	if err != nil {
		fmt.Println(err)
	}
	var policy *v1.NetworkPolicy
	policy.APIVersion = "networking.k8s.io/v1"
	policy.Kind = "NetworkPolicy"
	policy.ObjectMeta.Name = "default-deny-all"
	policy.ObjectMeta.Namespace = "test-namespace-1"
	//policy.Spec.PodSelector = 
	policy.Spec.PolicyTypes = []v1.PolicyType{"Ingress", "Egress"}
	networkPolicy, err := clientset.NetworkingV1().NetworkPolicies("test-namespace-1").Create(context.Background(), policy, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}

	log.LogNetworkPolicyApllying()
	//err = clientset.NetworkingV1().NetworkPolicies().
	fmt.Println(networkPolicy.String())


}