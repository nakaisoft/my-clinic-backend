package repository

import (
	"context"

	"github.com/nakaisoft/my-clinic-backend/internal/models"
	"github.com/nakaisoft/my-clinic-backend/internal/repository"
	"gorm.io/gorm"
)

// patientRepositoryImpl implementa PatientRepository com GORM.
type patientRepositoryImpl struct {
	db *gorm.DB
}

// NewPatientRepository cria um novo repositório de pacientes.
func NewPatientRepository(db *gorm.DB) repository.PatientRepository {
	return &patientRepositoryImpl{db: db}
}

// Insert adiciona um novo paciente, incluindo endereços.
func (r *patientRepositoryImpl) Insert(ctx context.Context, patient *models.PatientModel) error {
	return r.db.WithContext(ctx).Create(patient).Error
}

// GetByID busca um paciente pelo ID, incluindo os endereços.
func (r *patientRepositoryImpl) GetByID(ctx context.Context, id string) (*models.PatientModel, error) {
	var patient models.PatientModel
	err := r.db.WithContext(ctx).Preload("Addresses").First(&patient, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// GetAll retorna todos os pacientes do banco de dados, incluindo endereços.
func (r *patientRepositoryImpl) GetAll(ctx context.Context) ([]models.PatientModel, error) {
	var patients []models.PatientModel
	err := r.db.WithContext(ctx).Preload("Addresses").Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// Update atualiza os dados de um paciente.
func (r *patientRepositoryImpl) Update(ctx context.Context, patient *models.PatientModel) error {
	return r.db.WithContext(ctx).Save(patient).Error
}

// Delete remove um paciente pelo ID e seus endereços (por causa do `OnDelete:CASCADE`).
func (r *patientRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.PatientModel{}, "id = ?", id).Error
}
