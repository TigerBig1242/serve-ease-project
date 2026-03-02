package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
)

type MenuRepository interface {
	GetMenu(ctx context.Context) ([]entities.Menu, error)
}
