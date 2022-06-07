package luther

import (
	"os"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"path/filepath"
)

var usedIds []types.UID

func validateUID(uid types.UID) types.UID{
	for _, id := range usedIds {
		if id == uid {
			return uuid.NewUUID()
		}
	}
	
	return uid
}

func GenerateNetworkPolicyName(configPath string, setName string) string{
	uid := uuid.NewUUID()
	usedIds = append(usedIds, uid)
	return fileNameWithoutExtSliceNotation(configPath) + "-" + setName + "-" + string(validateUID(uid))
}

func fileNameWithoutExtSliceNotation(path string) string {
	fileStat, _ := os.Stat(path)
	fileName := fileStat.Name()
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func RemoveLabel(labels map[string]string, label string) map[string]string {
	delete(labels,label)
	return labels
}