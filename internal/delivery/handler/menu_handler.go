package handler

import (
	"github.com/bigthamm/serve-ease/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type MenuHandler struct {
	usecase usecase.MenuUseCase
}

type ResponseMenu struct {
	MenuID      uint   `json:"menu_id"`
	CategoryID  uint   `json:"category_id"`
	MenuName    string `json:"menu_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func NewMenuHandler(usecase usecase.MenuUseCase) *MenuHandler {
	return &MenuHandler{
		usecase: usecase,
	}
}

func (handler *MenuHandler) GetListMenu(c fiber.Ctx) error {
	menus, err := handler.usecase.GetMenus(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []ResponseMenu
	for _, listMenu := range menus {
		responseMenu := ResponseMenu{
			MenuID:      listMenu.MenuId,
			CategoryID:  listMenu.CategoryID,
			MenuName:    listMenu.MenuName,
			Description: listMenu.Description,
			Price:       listMenu.Price,
		}
		response = append(response, responseMenu)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}
