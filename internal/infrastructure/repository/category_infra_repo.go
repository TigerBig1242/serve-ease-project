package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
	"gorm.io/gorm"
)

type CategoryInfraRepo struct {
	db *gorm.DB
}

func CategoryInfra(db *gorm.DB) repository.CategoryRepository {
	return &CategoryInfraRepo{
		db: db,
	}
}

func (repo *CategoryInfraRepo) GetListCategory(ctx context.Context) ([]entities.Category, error) {
	var Category []entities.Category
	err := repo.db.WithContext(ctx).Find(&Category).Error
	if err != nil {
		return nil, err
	}

	return Category, nil
}
