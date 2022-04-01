package config

import (
	_"errors"
	"os"
	"testing"
	_"log"

	"github.com/stretchr/testify/assert"
)

var c Config

func TestReadV2Config(t *testing.T) {
	configPath := "../../test_data/v2_data.yaml"
	_, err := os.Stat(configPath)
	assert.Nil(t, err)

	configData, err := c.ReadConfig(configPath)
	assert.NotNil(t, configData)
	assert.Nil(t, err)

	assert.NotNil(t, configData.ConnectedSets)
	assert.Equal(t, 2, len((configData.ConnectedSets)))
}