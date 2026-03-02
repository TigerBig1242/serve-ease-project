package service

import "github.com/skip2/go-qrcode"

func CreateQRCode(data string, fileName string) error {
	err := qrcode.WriteFile(data, qrcode.Medium, 256, fileName)
	if err != nil {
		return err
	}

	return nil
}
