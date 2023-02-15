package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name is not blank"` // ต้องไม่เป็นค่าว่าง
	Email      string `valid:"email"`
	EmployeeID string // รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว

}
