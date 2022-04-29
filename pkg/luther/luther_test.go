package luther

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configPath = "../../test_data/configs/v2_data.yaml"
const namespace = "namespace-1"

func TestInitV1NetworkPolicies(t *testing.T) {

	networkPolicies, err := InitV1NetworkPolicies(configPath)
	assert.Nil(t, err)
	assert.Equal(t, len(networkPolicies), 2)

	for _, item := range networkPolicies {
		assert.Equal(t, reflect.TypeOf(item).Name(), "NetworkPolicy")
		assert.NotEmpty(t, item.TypeMeta)
		assert.NotEmpty(t, item.ObjectMeta)
		assert.NotEmpty(t, item.Spec)
		assert.Empty(t, item.Finalizers)
		assert.Empty(t, item.GenerateName)
	}
}

func TestFilterNetworkPolicyNames(t *testing.T) {

	setNames := []string{
		"set-1", 
		"set-2", 
		"database-set", 
		"monitoring",
	}

	networkPolicyNames := []string{
		"set-1-randomconvention", 
		"set-2-randomconvention", 
		"database-set-2345-7889-0943-9473", 
		"monitoring-set-2985-7840-7364-7789",
		"nw-policy-not-from-any-set",
	}

	filteredNetworkPolicyNames, err := FilterNetworkPolicyNames(setNames, networkPolicyNames, namespace)
	assert.Nil(t, err)
	assert.Equal(t, len(filteredNetworkPolicyNames), 4)
	assert.Equal(t, len(filteredNetworkPolicyNames), len(setNames))

	found := false
	for _, item := range filteredNetworkPolicyNames{
		if item == networkPolicyNames[4] {
			found = true
		}
	}
	assert.Equal(t, found, false)
}