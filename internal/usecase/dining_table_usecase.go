package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/domain/repository"
)

type DingingTableUseCase interface {
	CreateTable(ctx context.Context, table *entities.DiningTable) error
	GetTables(ctx context.Context) ([]entities.DiningTable, error)
	TableAvailable(ctx context.Context) ([]entities.DiningTable, error)
}

type diningTableUseCase struct {
	repo repository.DiningTableRepository
}

func NewDiningTableUseCase(repo repository.DiningTableRepository) *diningTableUseCase {
	return &diningTableUseCase{
		repo: repo,
	}
}

func (table *diningTableUseCase) CreateTable(ctx context.Context, diningTable *entities.DiningTable) error {
	if diningTable == nil {
		return errors.New("Require data to create table")
	}

	if diningTable.TableNumber == "0" {
		return errors.New("Number table must be not 0")
	}

	if diningTable.Status == "" {
		return errors.New("Require status is not empty")
	}

	createErr := table.repo.Create(ctx, diningTable)
	if createErr != nil {
		return fmt.Errorf("failed to save dining table to database: %w", createErr)
	}

	return nil
}

// Get all table
func (table *diningTableUseCase) GetTables(ctx context.Context) (tables []entities.DiningTable, error error) {
	DiningTables, err := table.repo.GetAllTable(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch dining tables: %w", err)
	}

	return DiningTables, nil
}

// Get table that has available status
func (table *diningTableUseCase) TableAvailable(ctx context.Context) (tables []entities.DiningTable, error error) {
	tableAvailable, err := table.repo.GetAllTable(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch table available")
	}

	var availableTables []entities.DiningTable
	for _, available := range tableAvailable {
		if available.Status != "booking" {
			fmt.Printf("Table ID %d is available\n", available.TableID)
			availableTables = append(availableTables, available)
		}
	}
	return availableTables, nil
}
