package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id             uuid.UUID `json:"id"`
	PrimeiroNome   string    `json:"primeiroNome" validate:"required"`
	UltimoNome     string    `json:"ultimoNome" validate:"required"`
	CPF            string    `json:"cpf" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	Password       string    `json:"password" validate:"required"`
	DataNascimento time.Time `json:"dataNascimento" validate:"required"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Id = uuid.New()
	u.CreatedAt = time.Now()
	return err
}
