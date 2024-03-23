package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() (uuid.UUID, error) {
	var emptyUUID uuid.UUID
	uuid, err := uuid.NewRandom()
	if err != nil {
		return emptyUUID, err
	}
	return uuid, nil
}
