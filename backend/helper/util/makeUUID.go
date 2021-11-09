package util

import (
	"strings"

	"github.com/google/uuid"
)

func MakeUUID() string {
	return uuid.New().String()
}

func MakeUUIDNoDash() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
