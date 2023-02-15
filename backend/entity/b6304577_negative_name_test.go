package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	em := Employee{
		Name:       "", // ผิด
		Email:      "phanupongkanin@gamil.com",
		EmployeeID: "J12345678",
	}

	ok, err := govalidator.ValidateStruct(em)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err.Error()).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Name is not blank"))

}
