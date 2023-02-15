package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `villid:"required"` // ต้องไม่เป็นค่าว่าง
	Email      string
	EmployeeID string // รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว

}
