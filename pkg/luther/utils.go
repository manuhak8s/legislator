package luther

import (
	"os"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"path/filepath"
)

// usedIds represents a slice with UIDs and saves them during an executed runtime.
var usedIds []types.UID

// validateUID checks if the given uid is unique and returns it.
// Otherwise a new one will be generated, appended to the usedIds slice and returned.
func validateUID(uid types.UID) types.UID{
	for _, id := range usedIds {
		if id == uid {
			newID := uuid.NewUUID()
			usedIds = append(usedIds, newID)
			return newID
		}
	}
	
	return uid
}

// GenerateNetworkPolicaName generates a unique string representation based on
// the name of the config file, the name of the associated set and an uid.
func generateNetworkPolicyName(configPath string, setName string) string{
	uid := uuid.NewUUID()
	usedIds = append(usedIds, uid)
	return fileNameWithoutExtSliceNotation(configPath) + "-" + setName + "-" + string(validateUID(uid))
}

// fileNameWithoutExtSliceNotation returns the filename of a given path without 
// its extension.
func fileNameWithoutExtSliceNotation(path string) string {
	fileStat, _ := os.Stat(path)
	fileName := fileStat.Name()
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

// RemoveLabel removes a key-value-pair of a given map by checking the existance
// of the key with the given label. After that it returns the map.
func removeLabel(labels map[string]string, label string) map[string]string {
	delete(labels,label)
	return labels
}