package models

import (
	"api_pattern_go/api/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permissao struct {
	Id        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
}

func (u *Permissao) BeforeCreate(_ *gorm.DB) (err error) {
	u.Id = uuid.New()
	return err
}

func (Permissao) TableName() string {
	return global.TablePermissao
}

type PermissaoUsuario struct {
	Id          uuid.UUID `json:"id"`
	IdPermissao uuid.UUID `json:"idPermissao" gorm:"column:id_permissao"`
	IdUsuario   uuid.UUID `json:"idUsuario" gorm:"column:id_usuario"`
}

func (u *PermissaoUsuario) BeforeCreate(_ *gorm.DB) (err error) {
	u.Id = uuid.New()
	return err
}

func (PermissaoUsuario) TableName() string {
	return global.TablePermissaoUsuario
}
