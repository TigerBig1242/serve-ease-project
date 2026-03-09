package handler

import (
	"fmt"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	"github.com/bigthamm/serve-ease/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type DiningTableHandler struct {
	usecase usecase.DingingTableUsecase
}

type CreateQrCodeHandler struct {
	createQrCode usecase.CreateQrCodeUseCase
	menu         usecase.MenuUseCase
}

type RequestBody struct {
	TableNumber int    `json:"table_number"`
	Status      string `json:"status"`
}

type CreateTableResponse struct {
	TableID     uint   `json:"table_id"`
	TableNumber int    `json:"table_number"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

const TimeFormat = "2/1/2006:15:04"

func NewDiningTableHandler(usecase usecase.DingingTableUsecase) *DiningTableHandler {
	return &DiningTableHandler{
		usecase: usecase,
	}
}

func NewCreateQrCodeHandler(createQrCode usecase.CreateQrCodeUseCase, menu usecase.MenuUseCase) *CreateQrCodeHandler {
	return &CreateQrCodeHandler{
		createQrCode: createQrCode,
		menu:         menu,
	}
}

func (handler *DiningTableHandler) Create(c fiber.Ctx) error {
	req := new(RequestBody)

	bodyErr := c.Bind().Body(&req)
	if bodyErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": bodyErr.Error(),
		})
	}

	responseTable := &entities.DiningTable{
		TableNumber: req.TableNumber,
		Status:      req.Status,
	}

	createErr := handler.usecase.CreateTable(c.Context(), responseTable)
	if createErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": createErr.Error(),
		})
	}

	response := CreateTableResponse{
		TableID:     responseTable.TableID,
		TableNumber: responseTable.TableNumber,
		Status:      responseTable.Status,
		// CreatedAt:   responseTable.CreatedAt.Format("2/1/2006:15:04"),
		CreatedAt: responseTable.CreatedAt.Format(TimeFormat),
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"data":   response,
		"error":  createErr,
	})
}

func (handler *DiningTableHandler) GetListTables(c fiber.Ctx) error {
	// responseTable := &entities.DiningTable{}
	getTables, err := handler.usecase.GetTables(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("Get tables :", getTables)
	var response []CreateTableResponse
	for _, tablesList := range getTables {
		responseTable := CreateTableResponse{
			TableID:     tablesList.TableID,
			TableNumber: tablesList.TableNumber,
			Status:      tablesList.Status,
			CreatedAt:   tablesList.CreatedAt.Format(TimeFormat),
		}
		response = append(response, responseTable)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (handler *DiningTableHandler) GetTableAvailable(c fiber.Ctx) error {
	availableTables, err := handler.usecase.TableAvailable(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []CreateTableResponse
	for _, tablesList := range availableTables {
		responseTable := CreateTableResponse{
			TableID:     tablesList.TableID,
			TableNumber: tablesList.TableNumber,
			Status:      tablesList.Status,
			CreatedAt:   tablesList.CreatedAt.Format("2/1/2006:15:04"),
		}
		response = append(response, responseTable)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (handler CreateQrCodeHandler) GenerateQrCode(c fiber.Ctx) error {
	url := handler.createQrCode.GenerateQrCode()

	return c.JSON(fiber.Map{
		"status":  "success",
		"url":     url,
		"message": "QR code generated successfully in your folder",
	})
}

func (handler CreateQrCodeHandler) HandleScan(c fiber.Ctx) error {
	token := c.Query("token")

	fmt.Println("--- 2 [Handler] มีคนสแกนเข้ามาแล้ว ---")
	fmt.Println("Token ที่ได้รับจากมือถือ :", token)

	menus, err := handler.menu.GetMenus(c.Context())
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

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Scan QR Code success",
		"data":    response,
		"token":   token,
	})
}
