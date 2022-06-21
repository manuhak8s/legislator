package luther

import (
	"github.com/manuhak8s/legislator/pkg/config"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var c config.Config
var testConfigPath = "../../test_data/configs/valid-config.yaml"

func TestInitLutherOpts(t *testing.T) {
	configData, err := c.ReadConfig(testConfigPath)
	assert.NotNil(t, configData)
	assert.Nil(t, err)

	namespaceList, err := generateCoreV1NamespaceList()
	assert.Nil(t, err)
	assert.NotNil(t, namespaceList)

	lutherOpts, err := InitLutherOpts(c.ConnectedSets, &namespaceList) 
	assert.Nil(t, err)
	assert.NotNil(t, lutherOpts)
	assert.Equal(t, len(lutherOpts), len(namespaceList.Items))

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
				Labels: map[string]string{
					"app": "nodejs",
				},
			},
		}
		namespaceList.Items = append(namespaceList.Items, namespace)
	}

	if len(namespaceList.Items) < 1 {
		return namespaceList, fmt.Errorf("generate namespace names error")
	}

	return namespaceList, nil
}