package registry

import (
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
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
	return handler.NewCreateQrCodeHandler(qrCodeUseCase /*menuUseCase*/)
}
