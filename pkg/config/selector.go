package config

type PodSelector []struct {
	// list of elements that should be selected by the selector
	PodName string `yaml:"podName"`
	PodExpression string `yaml:"podExpression"`
}