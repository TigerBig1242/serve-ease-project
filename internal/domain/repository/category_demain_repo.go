package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
)

type CategoryRepository interface {
	GetListCategory(ctx context.Context) ([]entities.Category, error)
}
