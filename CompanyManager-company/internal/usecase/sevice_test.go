package usecase

import (
	"encoding/json"
	"errors"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
	. "github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase/mock"

	"github.com/gojuno/minimock/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {

	var (
		mc      *minimock.Controller
		rep     *RepositoryMock
		useCase *UseCaseMock
		service *Service
		comp1   = &company.Company{}
		er      = errors.New("Mew error")
	)

	BeforeEach(func() {
		mc = minimock.NewController(GinkgoT())
		rep = NewRepositoryMock(mc)
		useCase = NewUseCaseMock(mc)
		service = NewService(rep)
	})

	Context("NewService", func() {
		It("Should't return not expected value", func() {
			Expect(service).NotTo(Equal(service.repo))
		})
	})

	Context("GetCompany", func() {
		It("Should return an error because can't convert  string to int", func() {
			useCase.GetCompanyMock.Expect([]byte("c")).Return(er)
			err := service.GetCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
		})

		It("Should return error if Get method returns error", func() {
			useCase.GetCompanyMock.Expect([]byte("6")).Return(er)
			rep.GetMock.Expect(int64(6)).Return(nil, er)

			err1 := service.GetCompany([]byte("6"))
			Expect(err1).To(HaveOccurred())

			comp, err2 := rep.Get(int64(6))
			Expect(err2).To(HaveOccurred())
			comp1 = nil
			Expect(comp).To(Equal(comp1))
		})

		It("Should return nil error", func() {
			useCase.GetCompanyMock.Expect([]byte("1")).Return(er)
			rep.GetMock.Expect(int64(1)).Return(comp1, nil)

			err1 := service.GetCompany([]byte("1"))

			comp, err2 := rep.Get(int64(1))
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal(comp1))

			Expect(err1).NotTo(HaveOccurred())
		})
	})

	Context("CreateCompany", func() {

		var (
			company = newFixtureCompany1()
			c, _    = json.Marshal(company)
		)

		It("Should return an error because can't convert from json to company", func() {
			useCase.CreateCompanyMock.Expect([]byte("c")).Return(er)
			err := service.CreateCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
		})

		It("Should return error if Create method returns error", func() {
			useCase.CreateCompanyMock.Expect(c).Return(er)
			rep.CreateMock.Expect(company).Return("", er)

			err1 := service.CreateCompany(c)

			comp, err2 := rep.Create(company)
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err1).To(HaveOccurred())
		})

		It("Should return nil error", func() {

			useCase.CreateCompanyMock.Expect(c).Return(nil)
			rep.CreateMock.Expect(company).Return("1", nil)

			err1 := service.CreateCompany(c)

			comp, err2 := rep.Create(company)
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal("1"))

			Expect(err1).NotTo(HaveOccurred())

		})

	})

	Context("UpdateCompany", func() {

		var (
			company = newFixtureCompany1()
			c, _    = json.Marshal(company)
		)

		It("Should return an error because can't convert from json to company", func() {
			useCase.UpdateCompanyMock.Expect([]byte("c")).Return(er)
			err := service.UpdateCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
		})

		It("Should return error if Update method returns error", func() {
			useCase.UpdateCompanyMock.Expect(c).Return(er)
			rep.UpdateMock.Expect(company).Return("", er)

			err1 := service.UpdateCompany(c)

			comp, err2 := rep.Update(company)
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err1).To(HaveOccurred())
		})

		It("Should return nil error", func() {

			useCase.UpdateCompanyMock.Expect(c).Return(nil)
			rep.UpdateMock.Expect(company).Return("1", nil)

			err1 := service.UpdateCompany(c)

			comp, err2 := rep.Update(company)
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal("1"))

			Expect(err1).NotTo(HaveOccurred())

		})

	})

	Context("DeleteCompany", func() {

		It("Should return an error because can't convert string to int", func() {
			useCase.DeleteCompanyMock.Expect([]byte("c")).Return(er)
			err := service.DeleteCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
		})

		It("Should return error if Delete method returns error", func() {
			useCase.DeleteCompanyMock.Expect([]byte("6")).Return(er)
			rep.DeleteMock.Expect(int64(6)).Return("", er)

			err1 := service.DeleteCompany([]byte("6"))
			Expect(err1).To(HaveOccurred())

			comp, err2 := rep.Delete(int64(6))
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))
		})

		It("Should return nil error", func() {
			useCase.DeleteCompanyMock.Expect([]byte("1")).Return(er)
			rep.DeleteMock.Expect(int64(1)).Return("", nil)

			err1 := service.DeleteCompany([]byte("1"))

			comp, err2 := rep.Delete(int64(1))
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err1).NotTo(HaveOccurred())
		})
	})

})



