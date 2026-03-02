package usecase

import (
	// "context"
	"fmt"
	"log"

	"github.com/bigthamm/serve-ease/internal/infrastructure/security"
	"github.com/bigthamm/serve-ease/internal/infrastructure/service"
)

type CreateQrCodeUseCase interface {
	GenerateQrCode() string
}

type qrCodeUseCase struct {
}

func NewCreateQrCodeUseCase() CreateQrCodeUseCase {
	return &qrCodeUseCase{}
}

func (usecase *qrCodeUseCase) GenerateQrCode() string {
	token, _ := security.GenerateRandomToken(16)

	baseURL := "https://1x2szl6m-8080.asse.devtunnels.ms"
	targetURL := fmt.Sprintf("%s/qr-code/scan-qr-code?token=%s", baseURL, token)

	fmt.Println("1 [Usecase] สร้าง URL สำเร็จ:", targetURL)

	errQR := service.CreateQRCode(targetURL, "qr_code_table.png")
	if errQR != nil {
		log.Println("Error generating QR: ", errQR)
	}

	return targetURL
}
