package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var c Config
var validTestConfigPath = "../../test_data/configs/valid-config.yaml"
var unvalidTestConfigPath = "../../test_data/configs/unvalid-config.yaml"

func TestReadConfig(t *testing.T) {
	validConfigData, err := c.ReadConfig(validTestConfigPath)
	assert.NotNil(t, validConfigData)
	assert.Nil(t, err)

	unvalidConfigData, err := c.ReadConfig(unvalidTestConfigPath)
	assert.Nil(t, unvalidConfigData)
	assert.NotNil(t, err)
}

func TestValidateLegislatorConfig(t *testing.T) {
	validConfigData, err := c.ReadConfig(validTestConfigPath)
	assert.NotNil(t, validConfigData)
	assert.Nil(t, err)

	err = validateLegislatorConfig(validConfigData, validTestConfigPath)
	assert.Nil(t, err)

	unvalidConfigData, err := c.ReadConfig(unvalidTestConfigPath)
	assert.Nil(t, unvalidConfigData)
	assert.NotNil(t, err)
}

func TestCheckForConnectedSetObligationFields(t *testing.T) {
	podSelector := PodSelector{
		MatchLabels: map[string]string{
			"foo": "bar",
		},
	}

	targetNamespaces := TargetNamespaces{
		MatchLabels: map[string]string{
			"foo": "bar",
		},
	}

	validConnectedSet := ConnectedSet{
		Name: "test-name",
		PodSelector: podSelector,
		TargetNamespaces: targetNamespaces,
	}
	err := checkForConnectedSetObligationFields(validConnectedSet, "/some/path")
	assert.Nil(t, err)

	unvalidConnectedSetWithoutName := ConnectedSet{
		PodSelector: podSelector,
		TargetNamespaces: targetNamespaces,
	}
	err = checkForConnectedSetObligationFields(unvalidConnectedSetWithoutName, "/some/path")
	assert.NotNil(t, err)

	unvalidConnectedSetWithoutTargetNamespaces := ConnectedSet{
		Name: "test-name",
		PodSelector: podSelector,
	}
	err = checkForConnectedSetObligationFields(unvalidConnectedSetWithoutTargetNamespaces, "/some/path")
	assert.NotNil(t, err)
}

func TestCheckForDuplicates(t *testing.T) {
	stringSliceWithoutDuplicates := []string{"one", "two", "three"}
	stringSliceWithDuplicates := []string{"one", "two", "three", "one"}

	err := checkForDuplicates(stringSliceWithoutDuplicates)
	assert.Nil(t, err)

	err = checkForDuplicates(stringSliceWithDuplicates)
	assert.NotNil(t, err)
}