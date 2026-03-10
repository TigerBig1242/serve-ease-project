package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
)

type DiningTableRepository interface {
	Create(ctx context.Context, DiningTable *entities.DiningTable) error
	GetAllTable(ctx context.Context) ([]entities.DiningTable, error)
}
