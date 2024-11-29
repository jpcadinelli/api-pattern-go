package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissaoRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Permissao, error)
	FindAll() ([]models.Permissao, error)
	Create(permissao *models.Permissao) error
	Update(permissao *models.Permissao) error
	Delete(id uuid.UUID) error
}

type permissaoRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissaoRepository(db *gorm.DB) PermissaoRepository {
	return &permissaoRepositoryImpl{db: db}
}

func (r *permissaoRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Permissao, error) {
	var permissao models.Permissao

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&permissao, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrPermissaoNaoEncontrada
	}

	return &permissao, nil
}

func (r *permissaoRepositoryImpl) FindAll() ([]models.Permissao, error) {
	var permissoes []models.Permissao

	tx := r.db.Find(&permissoes)
	if tx.Error != nil {
		return permissoes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return permissoes, erros.ErrPermissaoNaoEncontrada
	}

	return permissoes, nil
}

func (r *permissaoRepositoryImpl) Create(permissao *models.Permissao) error {
	return r.db.Create(permissao).Error
}

func (r *permissaoRepositoryImpl) Update(permissao *models.Permissao) error {
	return r.db.Save(permissao).Error
}

func (r *permissaoRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Permissao{}, "id = ?", id).Error
}
