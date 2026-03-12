package registry

import (
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
	"github.com/bigthamm/serve-ease/internal/infrastructure/repository"
	"github.com/bigthamm/serve-ease/internal/usecase"
	"gorm.io/gorm"
)

type Registry struct {
	db            *gorm.DB
	QrCodeHandler *handler.CreateQrCodeHandler
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{db: db}
}

func (registry *Registry) NewQrCodeHandler() *handler.CreateQrCodeHandler {
	qrCodeUseCase := usecase.NewCreateQrCodeUseCase()
	menuUseCase := registry.getMenuUseCase()
	return handler.NewCreateQrCodeHandler(qrCodeUseCase, menuUseCase)
}

func (registry *Registry) NewDiningTableHandler() *handler.DiningTableHandler {
	tableRepo := repository.NewDiningTableInfra(registry.db)
	tableUseCase := usecase.NewDiningTableUseCase(tableRepo)
	return handler.NewDiningTableHandler(tableUseCase)
}

func (registry *Registry) NewMenuHandler() *handler.MenuHandler {
	menuRepo := repository.MenuInfra(registry.db)
	menuUseCase := usecase.NewMenuUseCase(menuRepo)
	return handler.NewMenuHandler(menuUseCase)
}

func (registry *Registry) getMenuUseCase() usecase.MenuUseCase {
	repo := repository.MenuInfra(registry.db)
	return usecase.NewMenuUseCase(repo)
}

func (registry *Registry) NewCategoryHandler() *handler.CategoryHandler {
	categoryRepo := repository.CategoryInfra(registry.db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	return handler.NewCategoryHandler(categoryUseCase)
}
