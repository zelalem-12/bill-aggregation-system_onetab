package service

import "github.com/google/uuid"

func ToString(ID uuid.UUID) string {
	return ID.String()
}

func ToUUID(ID string) (uuid.UUID, error) {
	return uuid.Parse(ID)
}
