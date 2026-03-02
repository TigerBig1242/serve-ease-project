package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
	"gorm.io/gorm"
)

type MenuInfraRepo struct {
	db *gorm.DB
}

func MenuInfra(db *gorm.DB) repository.MenuRepository {
	return &MenuInfraRepo{
		db: db,
	}
}

func (repo *MenuInfraRepo) GetMenu(ctx context.Context) ([]entities.Menu, error) {
	var Menus []entities.Menu
	err := repo.db.WithContext(ctx).Find(&Menus).Error
	if err != nil {
		return nil, err
	}
	return Menus, nil
}
