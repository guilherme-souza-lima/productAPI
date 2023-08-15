package utils

import (
	"github.com/gofrs/uuid"
	"time"
)

var GeneratorUUid = func() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

var TimeNow = func() time.Time {
	return time.Now().UTC()
}
