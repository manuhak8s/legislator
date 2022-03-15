package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const constitutionConfigName = ".constitution.yaml"

type Config struct {
    Default string `yaml:"default"`
    ConnectedSets ConnectedSets `yaml:"connectedSets"`
}

func (config *Config) ReadConfig() (*Config, error) {

    yamlFile, err := ioutil.ReadFile("/Users/manuelhaugg/legislator/test.constitution.yaml")
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(yamlFile, config)
    if err != nil {
        return nil, err
    }

    return config, nil
}

func (config *Config) GetDefaultPolicy() (string, error) {
    defaultPolicy := config.Default
    if defaultPolicy == "" {
        return "", fmt.Errorf("no default policy found in constitution config")
    }
    
    return defaultPolicy, nil
}

func (config *Config) GetConnectedSets() (ConnectedSets, error) {
    connectedSets := config.ConnectedSets
    
    return connectedSets, nil
}

