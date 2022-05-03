package config

// TargetNamespaces represents the labels of namespaces where network policies
// become deployed.
type TargetNamespaces struct {
    // Map of key-value-pairs that shall be a deployment destination
    MatchLabels map[string]string `yaml:"matchLabels"`

    // String slice that should allow to select multiple namespaces by a defined naming convention
    MatchExpressions []string `yaml:"matchExpressions"`
}