package handler

import (
	"github.com/bigthamm/serve-ease/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	useCase usecase.CategoryUseCase
}

type ResponseCategory struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	Image        string `json:"image"`
	IsActive     bool   `json:"is_active"`
}

func NewCategoryHandler(useCase usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		useCase: useCase,
	}
}

func (handler *CategoryHandler) ListCategory(c fiber.Ctx) error {
	category, err := handler.useCase.GetCategory(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []ResponseCategory
	for _, listCategory := range category {
		responseCategory := ResponseCategory{
			CategoryID:   listCategory.CategoryID,
			CategoryName: listCategory.CategoryName,
			Image:        listCategory.Image,
			IsActive:     listCategory.IsActive,
		}
		response = append(response, responseCategory)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}
