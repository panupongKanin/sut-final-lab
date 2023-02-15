package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAlltrue(t *testing.T) {
	g := NewGomegaWithT(t)

	em := Employee{
		Name:       "panupong",
		Email:      "phanupongkanin@gamil.com",
		EmployeeID: "J12345678",
	}

	ok, _ := govalidator.ValidateStruct(em)

	g.Expect(ok).To(BeTrue())

}
