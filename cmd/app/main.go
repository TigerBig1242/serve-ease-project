package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/bigthamm/serve-ease/internal/config"
	// "github.com/bigthamm/serve-ease/internal/delivery/handler"
	router "github.com/bigthamm/serve-ease/internal/delivery/route"
	domain "github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/infrastructure/database"
	"github.com/bigthamm/serve-ease/internal/registry"
	// "github.com/bigthamm/serve-ease/internal/usecase"
)

func main() {
	databaseConfig := config.LoadConfig()

	db, err := database.ConnectDatabase(databaseConfig)

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Category{}, &domain.Menu{}, &domain.MenuOptions{}, &domain.DiningTable{},
		&domain.Order{}, &domain.OrderItems{}, &domain.OrderItemOptions{}, &domain.Payment{}, &domain.PaymentMethods{}, &domain.Roles{}, &domain.Users{})
	if err != nil {
		log.Fatalf("Database Migration failed: %v", err)
	}

	fmt.Println("Success QR Code saved as table_qr.png")

	// Create qr code at main.go file
	// qrUseCase := usecase.NewCreateQrCodeUseCase()
	// fmt.Println("--- System Initializing ---")
	// qrURL := qrUseCase.GenerateQrCode()
	// fmt.Println("QR Code has been generated at root folder:", qrURL)

	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}
		return c.SendString("Hello World at main.go")
	})

	request := registry.NewRegistry(db)

	router.SetupRoute(app, request.NewDiningTableHandler())

	// Call to use create qr code
	// qrHandler := handler.NewCreateQrCodeHandler(qrUseCase)
	// router.SetQrCodeRoute(app, qrHandler)

	router.SetQrCodeRoute(app, request.NewQrCodeHandler())
	router.SetMenuRoute(app, request.NewMenuHandler())
	router.SetCategoryRoute(app, request.NewCategoryHandler())

	app.Listen(":8080")
	fmt.Println("Hello serve ease")
}
