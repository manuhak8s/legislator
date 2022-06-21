package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

// GetK8SClient creates a kubernetes go-client based on the set kubeconfig env var.
// If no env var is set, GetK8SClient reads the current kubecontext from the .kube directory
// in the home directory. If no kubecontext can be read, an error will be triggered.
func GetK8sClient() (*kubernetes.Clientset, error) {
	var clientset *kubernetes.Clientset
	kubeconfigENV := os.Getenv("KUBECONFIG")
	if kubeconfigENV == "" {

		home, exists := os.LookupEnv("Home")
		if !exists {
			home, _ = os.UserHomeDir()
		}
	
		kubeConfigPath := filepath.Join(home, ".kube", "config")
	
		clientset, _ = initClientset(kubeConfigPath)
	} else {
		clientset, _ = initClientset(kubeconfigENV)
	}

	if clientset == nil {
		return nil, fmt.Errorf("unknown error while creating kubernetes client: please check your kubecontext definition or read the legislator instructions")
	}

	return clientset, nil
}

// initClientset accepts a string parameter as a path representation to a kubeconfig-file
// and creates clientset.
func initClientset(kubeconfigPath string) (*kubernetes.Clientset, error){
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// GetK8SDefaultClient initialized a default kubernetes client without any kubecontext or
// kubeconfig-file definitions. It can be used for testing purposes.
func GetK8SDefaultClient() (*kubernetes.Clientset, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}