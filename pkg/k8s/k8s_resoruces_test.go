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

func deleteTestNamespace(clientset *kubernetes.Clientset, namespaceName string) error {
	err := clientset.CoreV1().Namespaces().Delete(context.Background(), namespaceName,metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}