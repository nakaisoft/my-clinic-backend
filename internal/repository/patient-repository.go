package repository

import (
	"context"

	"github.com/nakaisoft/my-clinic-backend/internal/models"

)

// PatientRepository define os métodos do repositório.
type PatientRepository interface {
	Insert(ctx context.Context, patient *models.PatientModel) error
	GetByID(ctx context.Context, id string) (*models.PatientModel, error)
	GetAll(ctx context.Context) ([]models.PatientModel, error)
	Update(ctx context.Context, patient *models.PatientModel) error
	Delete(ctx context.Context, id string) error
}
