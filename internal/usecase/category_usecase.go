package usecase

import (
	"context"
	"fmt"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
)

type CategoryUseCase interface {
	GetCategory(ctx context.Context) ([]entities.Category, error)
}

type categoryUseCase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) *categoryUseCase {
	return &categoryUseCase{
		repo: repo,
	}
}

func (category *categoryUseCase) GetCategory(ctx context.Context) (listCategory []entities.Category, error error) {
	AllCategory, err := category.repo.GetListCategory(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch category: %w", err)
	}

	return AllCategory, nil
}
