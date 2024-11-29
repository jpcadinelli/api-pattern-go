package repository

import (
	"api_pattern_go/api/global/erros"
	"api_pattern_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	FindById(id uuid.UUID) (*models.Usuario, error)
	FindAll() ([]models.Usuario, error)
	Create(usuario *models.Usuario) error
	Update(usuario *models.Usuario) error
	Delete(id uuid.UUID) error
}

type usuarioRepositoryImpl struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepositoryImpl{db: db}
}

func (r *usuarioRepositoryImpl) FindById(id uuid.UUID) (*models.Usuario, error) {
	var usuario models.Usuario

	tx := r.db.First(&usuario, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrUsuarioNaoEncontrado
	}

	return &usuario, nil
}

func (r *usuarioRepositoryImpl) FindAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario

	tx := r.db.Find(&usuarios)
	if tx.Error != nil {
		return usuarios, tx.Error
	}
	if tx.RowsAffected == 0 {
		return usuarios, erros.ErrUsuarioNaoEncontrado
	}

	return usuarios, nil
}

func (r *usuarioRepositoryImpl) Create(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *usuarioRepositoryImpl) Update(usuario *models.Usuario) error {
	return r.db.Save(usuario).Error
}

func (r *usuarioRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Usuario{}, "id = ?", id).Error
}
