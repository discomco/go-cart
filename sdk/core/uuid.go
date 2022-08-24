package core

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

func NewUuid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error [%+v] in cleanUid", err)
	}
	return uid.String(), err
}

func NullId() string {
	return uuid.Nil.String()
}

func CleanNullId() string {
	return strings.Replace(uuid.Nil.String(), "-", "", -1)
}

func CleanUuid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error [%+v] in cleanUid:", err)
	}
	return strings.Replace(uid.String(), "-", "", -1), err
}
