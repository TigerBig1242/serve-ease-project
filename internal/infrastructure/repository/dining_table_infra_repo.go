package repository

import (
	"context"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
	"gorm.io/gorm"
)

type diningTableInfraRepo struct {
	db *gorm.DB
}

func CreateDiningTable(db *gorm.DB) repository.DiningTableRepository {
	if db == nil {
		panic("Data is empty, Require data to create")
	}

	return &diningTableInfraRepo{db: db}
}

func (repo *diningTableInfraRepo) Create(ctx context.Context, diningTable *entities.DiningTable) error {
	return repo.db.WithContext(ctx).Create(diningTable).Error
}

func (repo *diningTableInfraRepo) GetAllTable(ctx context.Context) ([]entities.DiningTable, error) {
	var diningTables []entities.DiningTable
	err := repo.db.WithContext(ctx).Find(&diningTables).Error
	if err != nil {
		return nil, err
	}
	return diningTables, nil
}
