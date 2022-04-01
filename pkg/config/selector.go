package config

// PodSelector represents the required fields for selecting pods
type PodSelector struct {
    // Map of key-value-pairs that shall translated to a network policy
    MatchLabels map[string]string `yaml:"matchLabels"`

    // String slice that should allow to select multiple pods by a defined naming convention
    MatchExpressions []string `yaml:"matchExpressions"`
}