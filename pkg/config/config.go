package config

import (
	_"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the data model of a config file for legislator
type Config struct {
    //Default string `yaml:"default"`

    // Datafield for the connected sets that should be created
    ConnectedSets ConnectedSets `yaml:"connectedSets"`

}

// ReadConfig reads a yaml configuration of the given path
// A config instance has to be defined for reading the file
// This function returns a config data model with all fields from the config file
func (config *Config) ReadConfig(configPath string) (*Config, error) {

    yamlFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(yamlFile, config)
    if err != nil {
        return nil, err
    }

    return config, nil
}