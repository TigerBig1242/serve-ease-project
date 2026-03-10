package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/bigthamm/serve-ease/internal/config"
	domain "github.com/bigthamm/serve-ease/internal/domain/entities"

	// "github.com/bigthamm/serve-ease/internal/domain/repository"
	"github.com/bigthamm/serve-ease/internal/infrastructure/database"
	"github.com/bigthamm/serve-ease/internal/usecase"

	"github.com/bigthamm/serve-ease/internal/infrastructure/repository"
	// "github.com/bigthamm/serve-ease/internal/usecase"
	"github.com/bigthamm/serve-ease/internal/delivery/handler"
	router "github.com/bigthamm/serve-ease/internal/delivery/route"
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
	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	// Registry Dining table
	diningTableRepo := repository.NewDiningTableInfra(db)
	diningTableUseCase := usecase.NewDiningTableUseCase(diningTableRepo)
	diningTableHandler := handler.NewDiningTableHandler(diningTableUseCase)
	router.SetupRoute(app, diningTableHandler)

	// Registry Category
	categoryRepo := repository.CategoryInfra(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	router.SetCategoryRoute(app, categoryHandler)

	// Registry Menu
	menuRepo := repository.MenuInfra(db)
	menuUseCase := usecase.NewMenuUseCase(menuRepo)
	menuHandler := handler.NewMenuHandler(menuUseCase)
	router.SetMenuRoute(app, menuHandler)

	app.Listen(":8080")
	fmt.Println("Hello serve ease")
}
