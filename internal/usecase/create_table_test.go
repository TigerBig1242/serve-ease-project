package usecase_test

import (
	"context"
	// "errors"
	"testing"

	"github.com/bigthamm/serve-ease/internal/domain/entities"
	// "github.com/bigthamm/serve-ease/internal/usecase"
)

type MockCreateTableRepo struct {
	// เราสามารถกำหนดได้ว่าอยากให้ฟังก์ชัน Create คืนค่า Error อะไรออกมา
	mockCreateErr error
}

func (m *MockCreateTableRepo) Create(ctx context.Context, incident *entities.DiningTable) error {
	return m.mockCreateErr
}

func TestCreateReport(t *testing.T) {
	// t.Run("ควรบันทึกสำเร็จเมื่อข้อมูลถูกต้อง", func(t *testing.T) {
	// 	// Arrange: เตรียมของ
	// 	mockRepo := &MockCreateTableRepo{mockCreateErr: nil} // ให้จำลองว่า DB ทำงานปกติ
	// 	uc := usecase.NewDiningTableUsecase(mockRepo)
	// 	input := &entities.DiningTable{TableNumber: 1}

	// 	// Act: ลงมือรัน
	// 	err := uc.CreateTable(context.Background(), input)

	// 	// Assert: ตรวจสอบผล
	// 	if err != nil {
	// 		t.Errorf("ต้องไม่มี error แต่ได้: %v", err)
	// 	}
	// })

	// t.Run("ควรส่ง Error กลับมาถ้า Database มีปัญหา", func(t *testing.T) {
	// 	// Arrange: ให้จำลองว่า DB ระเบิด
	// 	mockRepo := &MockCreateTableRepo{mockCreateErr: errors.New("database connection failed")}
	// 	uc := usecase.NewDiningTableUsecase(mockRepo)
	// 	input := &entities.DiningTable{TableNumber: 2}

	// 	// Act
	// 	err := uc.CreateTable(context.Background(), input)

	// 	// Assert
	// 	if err == nil {
	// 		t.Error("ต้องได้ error แต่กลับได้ nil")
	// 	}
	// })
}
