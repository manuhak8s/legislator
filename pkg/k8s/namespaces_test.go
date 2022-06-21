package k8s

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetNamespaceNames(t *testing.T) {
	namespaceList, err := generateCoreV1NamespaceList()
	assert.Nil(t, err)
	assert.NotNil(t, namespaceList)

	namespaceNames := GetNamespaceNames(&namespaceList)
	assert.NotNil(t, namespaceNames)
	assert.Equal(t, len(namespaceNames), len(namespaceList.Items))
}

func generateCoreV1NamespaceList() (corev1.NamespaceList, error){
	var namespaceList corev1.NamespaceList
	names := map[string]string{
		"1": "ns-1",
		"2": "ns-2",
		"3": "ns-3",
		"4": "ns-4",
	}

	for _, name := range names {
		namespace := corev1.Namespace {
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}
		namespaceList.Items = append(namespaceList.Items, namespace)
	}

	if len(namespaceList.Items) < 1 {
		return namespaceList, fmt.Errorf("generate namespace names error")
	}

	return namespaceList, nil
}