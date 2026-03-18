package route

import (
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupTableRoute(app *fiber.App, createTableHandler *handler.DiningTableHandler) {
	createTable := app.Group("/dining-table")
	{
		createTable.Post("/create-table", createTableHandler.Create)
		createTable.Get("/get-list-tables", createTableHandler.GetListTables)
		createTable.Get("/available", createTableHandler.GetTableAvailable)
	}
}

func SetMenuRoute(app *fiber.App, menusHandler *handler.MenuHandler) {
	app.Get("/get-list-menus", menusHandler.GetListMenu)
}

func SetCategoryRoute(app *fiber.App, categoryHandler *handler.CategoryHandler) {
	app.Get("/get-category", categoryHandler.ListCategory)
}

func SetQrCodeRoute(app *fiber.App, scanQrCodeHandler *handler.CreateQrCodeHandler) {
	QrCode := app.Group("qr-code")
	{
		QrCode.Get("/admin/generate", func(c fiber.Ctx) error {
			url := scanQrCodeHandler.GenerateQrCode(c)
			return url
		})
		QrCode.Get("/scan-qr-code", scanQrCodeHandler.HandleScan)
	}
}
