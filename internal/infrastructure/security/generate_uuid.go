package security

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUUID() (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.UUID{}, err
	}

	fmt.Printf("UUID V7 (String): %s\n", id.String())
	fmt.Printf("UUID V7 (Bytes): %v\n", id)

	return id, err
}
