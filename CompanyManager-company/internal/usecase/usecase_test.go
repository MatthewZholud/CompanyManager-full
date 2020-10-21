package usecase

import (
	"encoding/json"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")

}

func newFixtureCompany1() *company.Company {
	return &company.Company{
		ID:        1,
		Name:      "Mega",
		Legalform: "OOO",
	}
}

type fakeCompany struct {
	ID        string
	Name      string
	Legalform string
}

var _ = Describe("Usecase", func() {
	Describe("StringToInt64", func() {
		It("Should return an error and zero if can't do Atoi conversion", func() {
			got, err := StringToInt64("c")
			Expect(err).To(HaveOccurred())
			Expect(got).To(Equal(int64(0)))
		})

		It("Should return an int64 value and nil error if Atoi conversion works", func() {
			got, err := StringToInt64("2")
			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(int64(2)))
		})
	})

	Describe("JsonToCompany", func() {
		It("Should return an error if not cant atoi", func() {
			comp := newFixtureCompany1()
			c, _ := json.Marshal(comp)
			got, err := JsonToCompany(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(comp))
		})

		It("Should return an int 64 variable if atoi works", func() {
			comp := fakeCompany{}
			c, _ := json.Marshal(comp)
			_, err := JsonToCompany(c)
			Expect(err).To(HaveOccurred())
		})
	})
})
