package models

import (
	"api_pattern_go/api/global"
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

func (Teste) TableName() string {
	return global.TableTeste
}

type TesteFiltro struct {
	Nome string `json:"nome"`
}
