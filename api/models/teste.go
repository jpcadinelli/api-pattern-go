package models

import "github.com/google/uuid"

type Teste struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Nome string    `json:"nome"`
}
