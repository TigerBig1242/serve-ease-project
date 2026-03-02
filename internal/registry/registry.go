package registry

import (
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
	"github.com/bigthamm/serve-ease/internal/infrastructure/repository"
	"github.com/bigthamm/serve-ease/internal/usecase"
	"gorm.io/gorm"
)

type Registry struct {
	db            *gorm.DB
	TableHandler  *handler.DiningTableHandler
	QrCodeHandler *handler.CreateQrCodeHandler
	menuHandler   *handler.MenuHandler
}

func NewRegistry(db *gorm.DB) *Registry {
	// tableRepo := repository.CreateDiningTable(db)
	// tableUseCase := usecase.NewDiningTableUsecase(tableRepo)

	// qrCodeUseCase := usecase.NewCreateQrCodeUseCase()

	// return &Registry{
	// 	TableHandler:  handler.NewDiningTableHandler(tableUseCase),
	// 	QrCodeHandler: handler.NewCreateQrCodeHandler(qrCodeUseCase),
	// }
	return &Registry{db: db}
}

func (registry *Registry) NewDiningTableHandler() *handler.DiningTableHandler {
	tableRepo := repository.CreateDiningTable(registry.db)
	tableUseCase := usecase.NewDiningTableUsecase(tableRepo)
	return handler.NewDiningTableHandler(tableUseCase)
}

func (registry *Registry) NewQrCodeHandler() *handler.CreateQrCodeHandler {
	qrCodeUseCase := usecase.NewCreateQrCodeUseCase()
	return handler.NewCreateQrCodeHandler(qrCodeUseCase)
}

func (registry *Registry) NewMenuHandler() *handler.MenuHandler {
	menuRepo := repository.MenuInfra(registry.db)
	menuUseCase := usecase.NewMenuUseCase(menuRepo)
	return handler.NewMenuHandler(menuUseCase)
}
