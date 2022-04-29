package config

// ConnectedSets represents a listed struct of sets
type ConnectedSets []struct {
	// Name of the connected set represented by a string value
	Name string `yaml:"name"`

	// Selector for selecting kubernetes pods and assigns them to a set
	PodSelector PodSelector `yaml:"podSelector"`
}