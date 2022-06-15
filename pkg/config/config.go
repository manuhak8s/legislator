package config

import (
	"fmt"
	"io/ioutil"

    "github.com/manuhak8s/legislator/pkg/logger"
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

    logger.TriggerOutput("loading", "... validating config file ... ")
    err = ValidateLegislatorConfig(config, configPath)
    if err != nil {
        return nil, err
    }
    logger.TriggerOutput("success", "config validation successfully finished :)")
    return config, nil
}

func ValidateLegislatorConfig(config *Config, configPath string) error {
    if config.ConnectedSets == nil {
        return fmt.Errorf("error: no connectedSets field defined at config: %s - please read the legislator instructions for further information", configPath)
    }
    
    var setNames []string
    for _, set := range config.ConnectedSets {
        setNames = append(setNames, set.Name)
    }

    err := checkForDuplicates(setNames)
    if err != nil {
        return err
    }

    for _, set := range config.ConnectedSets {
        err = checkForConnectedSetObligationFields(set, configPath)
        if err != nil {
            return err
        }
    }

    return nil
}

func checkForDuplicates(stringSlice []string) error {
    dupcount := 0
    for i := 0; i < len(stringSlice); i++ {
        for j := i + 1; j < len(stringSlice); j++ {
            if stringSlice[i] == stringSlice[j] {
                dupcount++
                break
            }
        }
    }
    if dupcount != 0 {
        return fmt.Errorf("duplicate connected set name found: duplicates are not allowed as set names, for further information read the legislator instructions")
    }
    return nil
}

func checkForConnectedSetObligationFields(connectedSet ConnectedSet, configPath string) error {
    if connectedSet.Name == "" {
        return fmt.Errorf("obligation field error: no name defined for connected set at config: %s - a name has to be set for a connected set, for further information read the legislator instructions", configPath)
    }

    if connectedSet.TargetNamespaces.MatchLabels == nil && connectedSet.TargetNamespaces.MatchExpressions == nil {
        return fmt.Errorf("obligation field error: no selector defined for target namespaces at config: %s - a selector has to be defined for a connected set, for further information read the legislator instructions", configPath)
    }

    /*if connectedSet.PodSelector.MatchLabels == nil && connectedSet.PodSelector.MatchExpressions == nil {
        return fmt.Errorf("obligation field error: no selector defined target pods at config: %s - a selector has to be defined for a connected set, for further information read the legislator instructions", configPath)
    }*/

    return nil
}