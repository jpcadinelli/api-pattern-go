package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teste struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Nome string    `json:"nome"`
}

func (t *Teste) BeforeCreate(_ *gorm.DB) (err error) {
	t.Id = uuid.New()
	return err
}

type TesteFiltro struct {
	Nome string `json:"nome"`
}
