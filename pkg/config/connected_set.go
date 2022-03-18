package config

type ConnectedSets []struct {
	// Name of the connected set represented by a string value
	Name string `yaml:"name"`

	// Selector for selecting kubernetes components to a set
	// e.g. pods or namespaces 
	PocSelector PodSelector `yaml:"podSelector"`
}