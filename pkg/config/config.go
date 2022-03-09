package config

import (
	_"fmt"
)

const constitutionConfigName = ".constitution.yaml"

type Config struct {

	// Default value for attaching network policies by one key-value-pair.
	// e.g. denyAll, allowAll, denyOtherNamespaces, ...
	// default string `yaml:"default"`

	// List of sets that contains pods or namespaces that are connected or not.
	// Each set represents an isolated bubble with its inner components.
	// ConnectedSets

	// Name of the set.
	// Name

	// Selects pods by their explicit name to a set.
	// PodSelector

	// Represents a pod by its name.
	// PodName

	// Regular expression of pods for selecting them by a naming convention.
	// e.g. etcd-*
	// PodExpression

	// Selects namespaces by their explicit name.
	// NamespaceSelector

	// Represents a namespace by its name.
	// NamespaceName

	// Regular expression of pods for selecting them by a naming convention.
	// NamespaceExpression


}