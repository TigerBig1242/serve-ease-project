package usecase

import (
	"context"
	"fmt"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
)

type MenuUseCase interface {
	GetMenus(ctx context.Context) ([]entities.Menu, error)
}

type menuUseCase struct {
	repo repository.MenuRepository
}

func NewMenuUseCase(repo repository.MenuRepository) *menuUseCase {
	return &menuUseCase{
		repo: repo,
	}
}

func (menu *menuUseCase) GetMenus(ctx context.Context) (menus []entities.Menu, error error) {
	AllMenu, err := menu.repo.GetMenu(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch menu: %w", err)
	}
	return AllMenu, nil
}
