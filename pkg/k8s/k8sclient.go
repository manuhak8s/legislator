package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

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
		return nil, fmt.Errorf("unknown error while creating kubernetes client: please try again")
	}

	return clientset, nil
}

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

func GetK8sDefaultClient() (*kubernetes.Clientset, error) {
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