package luther

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
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

func GenerateNetworkPolicyName(setName string) string{
	uid := uuid.NewUUID()

	return setName + "-" + string(validateUID(uid))
}