package repository

import (
	"context"

	"github.com/nakaisoft/my-clinic-backend/internal/models"
)

// AdressRepository define os métodos do repositório.
type AdressRepository interface {
	Insert(ctx context.Context, patient *models.PersonalAddress) error
}