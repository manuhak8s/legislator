package k8s

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods creates a clientset based on the k8s packages
// Additionally this function reads all pods from the current kube context
// It returns a list of coreV1 pods
func GetPods(podName string) (*corev1.PodList, error){

	clientset, err := GetK8sClient()
	if err != nil {
		fmt.Println(err)
	}

	pods, err := clientset.CoreV1().Pods(podName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	return pods, nil
}

// GetPodNames accepts a list of coreV1 pods and returns all of their names
// The names are packed into a string slice
func GetPodNames(pods *corev1.PodList) []string {
	var podNameArr []string
	for _, pod := range pods.Items {
		podNameArr = append(podNameArr, pod.Name)
	}

	return podNameArr
}

// GetPod returns a coreV1 pod instance based on the required pod name
// This function compares alle pods with the given string value
// If no pod is found an empty pod instance and error is returned
func GetPod(podName string,pods *corev1.PodList) (corev1.Pod, error) {
	var emptyPodInstance corev1.Pod

	if len(pods.Items) < 1 {
		return emptyPodInstance, fmt.Errorf("no Pods found in current cluster context")
	}

	for _, pod := range pods.Items {
		if podName == pod.Name {
			return pod, nil
		}
	}

	return emptyPodInstance, fmt.Errorf("Pod not found")
}

// PodExists checks if a pod exists in the current kube context
// It accepts the string representation of a name and a coreV1 list of pods
// If an item matches a bool value true is returned, otherwise false
func PodExists(podName string, pods *corev1.PodList) bool {
	podNames := GetPodNames(pods)
	if len(podNames) < 1 {
		log.Println("no Pods found")
		return false
	}

	for _, name := range podNames {
		if podName == name {
			return true
		}
	}

	return false
}