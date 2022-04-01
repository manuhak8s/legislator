package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var c Config

func TestReadConfig(t *testing.T) {
	configPath := "../../test_data/constitution.yaml"
	_, err := os.Stat(configPath)
	assert.Nil(t, err)

	configData, err := c.ReadConfig(configPath)
	assert.NotNil(t, configData)
	assert.Nil(t, err)

	assert.NotNil(t, configData.Default)
	assert.Equal(t, "denyAll", configData.Default)

	assert.NotNil(t, configData.ConnectedSets)
	assert.Equal(t, 3, len((configData.ConnectedSets)))

	for _, set := range configData.ConnectedSets {
		if set.PocSelector == nil {
			err = errors.New("no podname defined in pod selector")
			assert.Nil(t, err)
		}
	}

}