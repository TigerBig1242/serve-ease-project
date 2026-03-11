package route

import (
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
	"github.com/gofiber/fiber/v3"
)

func SetQrCodeRoute(app *fiber.App, scanQrCodeHandler *handler.CreateQrCodeHandler) {
	QrCode := app.Group("qr-code")
	{
		QrCode.Get("/admin/generate", func(c fiber.Ctx) error {
			url := scanQrCodeHandler.GenerateQrCode(c)
			return url
		})
		QrCode.Get("/scan-qr-code", scanQrCodeHandler.HandleScan)
	}

	// app.Get("/scan-qr-code", scanQrCodeHandler.HandleScan)
}
