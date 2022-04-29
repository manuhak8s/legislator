package k8s

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNamespaceFuncs(t *testing.T) {
	clientset, err := GetK8sClient()
	assert.Nil(t, err)
	
	err = initTestNamespace(clientset)
	assert.Nil(t, err)

	namespaces, err := GetNamespaces()
	assert.Nil(t, err)
	assert.Equal(t, true, len(namespaces.Items)>=1)

	namespaceNames := GetNamespaceNames(namespaces)
	assert.Equal(t, true, len(namespaceNames)>=1)
	assert.Equal(t, len(namespaces.Items), len(namespaceNames))

	testNamespace, err := GetNamespace("test-namespace", namespaces)
	assert.Nil(t, err)
	assert.Equal(t, "test-namespace", testNamespace.Name)

	exists := NamespaceExists("test-namespace", namespaces)
	assert.Equal(t, true, exists)

	err = deleteTestNamespace(clientset, "test-namespace")
	assert.Nil(t, err)
}

func TestPodFuncs(t *testing.T) {
	clientset, err := GetK8sClient()
	assert.Nil(t, err)
	
	err = initTestNamespace(clientset)
	assert.Nil(t, err)

	err = InitTestPods(clientset, "test-namespace")
	assert.Nil(t, err)

	pods, err := GetPods("test-namespace")
	assert.Nil(t, err)
	assert.Equal(t, true, len(pods.Items)>=1)
	assert.Equal(t, len(pods.Items), 2)

	podNames := GetPodNames(pods)
	assert.Equal(t, true, len(podNames)>=1)
	assert.Equal(t, len(pods.Items), len(podNames))

	testPod1, err := GetPod("test-pod-1", pods)
	assert.Nil(t, err)
	assert.Equal(t, "test-pod-1", testPod1.Name)

	testPod2, err := GetPod("test-pod-2", pods)
	assert.Nil(t, err)
	assert.Equal(t, "test-pod-2", testPod2.Name)

	exists := PodExists("test-pod-1", pods)
	assert.Equal(t, true, exists)

	exists = PodExists("test-pod-2", pods)
	assert.Equal(t, true, exists)

	err = deleteTestNamespace(clientset, "test-namespace")
	assert.Nil(t, err)

}

func initTestNamespace(clientset *kubernetes.Clientset) error {
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(context.Background(), namespace, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func InitTestPods(clientset *kubernetes.Clientset, nsName string) error{

	pod_1 := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod-1",
			Namespace: "test-namespace",
		},
		Spec: corev1.PodSpec {
			Containers: []corev1.Container{
				corev1.Container{
					Name: "nginx",
					Image: "nginx",
				},
			},
		},
	}

	pod_2 := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod-2",
			Namespace: "test-namespace",
		},
		Spec: corev1.PodSpec {
			Containers: []corev1.Container{
				corev1.Container{
					Name: "nginx",
					Image: "nginx",
				},
			},
		},
	}
	
	_, err := clientset.CoreV1().Pods(nsName).Create(context.Background(), pod_1, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	
	_, err = clientset.CoreV1().Pods(nsName).Create(context.Background(), pod_2, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func deleteTestNamespace(clientset *kubernetes.Clientset, namespaceName string) error {
	err := clientset.CoreV1().Namespaces().Delete(context.Background(), namespaceName,metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}