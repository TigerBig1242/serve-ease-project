package handler

import (
	"fmt"

	"github.com/bigthamm/serve-ease/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type CreateQrCodeHandler struct {
	createQrCode usecase.CreateQrCodeUseCase
}

func NewCreateQrCodeHandler(createQrCode usecase.CreateQrCodeUseCase /*menu usecase.MenuUseCase*/) *CreateQrCodeHandler {
	return &CreateQrCodeHandler{
		createQrCode: createQrCode,
	}
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

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Scan QR Code success",
		"token":   token,
	})
}
