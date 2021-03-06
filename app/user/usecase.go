package user

import (
	"context"

	"github.com/iamaul/evonix-backend-api/app/models"
)

type Usecase interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByName(ctx context.Context, name string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Store(context.Context, *models.User) error
}
