package usecase

import (
	"encoding/json"
	"errors"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity/company"
	"github.com/gojuno/minimock/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"


	. "github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase/mock"

)

var _ = Describe("Service", func() {

	var (
		mc      *minimock.Controller
		rep     *RepositoryMock
		useCase *UseCaseMock
		service *Service
		comp1   = &company.Company{}
		er      = errors.New("new error")
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
		var resp []byte

		It("Should return an error because can't convert  string to int", func() {
			useCase.GetCompanyMock.Expect([]byte("c")).Return(nil, er)
			response, err := service.GetCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))
		})

		It("Should return error if Get method returns error", func() {
			useCase.GetCompanyMock.Expect([]byte("6")).Return(nil, er)
			rep.GetMock.Expect(int64(6)).Return(nil, er)

			response, err := service.GetCompany([]byte("6"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))

			comp, err2 := rep.Get(int64(6))
			Expect(err2).To(HaveOccurred())
			comp1 = nil
			Expect(comp).To(Equal(comp1))
		})

		It("Should return nil error", func() {

			useCase.GetCompanyMock.Expect([]byte("1")).Return(resp, nil)
			rep.GetMock.Expect(int64(1)).Return(comp1, nil)

			response, err := service.GetCompany([]byte("1"))

			comp, err2 := rep.Get(int64(1))
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal(comp1))

			resp = []byte("null")
			Expect(response).To(Equal(resp))
			Expect(err).NotTo(HaveOccurred())


		})
	})

	Context("CreateCompany", func() {

		var (
			company = newFixtureCompany1()
			c, _    = json.Marshal(company)
			resp []byte
		)

		It("Should return an error because can't convert from json to company", func() {
			useCase.CreateCompanyMock.Expect([]byte("c")).Return(nil, er)
			response, err := service.CreateCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))
		})

		It("Should return error if Create method returns error", func() {
			useCase.CreateCompanyMock.Expect(c).Return(nil, er)
			rep.CreateMock.Expect(company).Return("", er)

			response, err := service.CreateCompany(c)

			comp, err2 := rep.Create(company)
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))
		})

		It("Should return nil error", func() {

			useCase.CreateCompanyMock.Expect(c).Return(resp, nil)
			rep.CreateMock.Expect(company).Return("1", nil)

			response, err := service.CreateCompany(c)

			comp, err2 := rep.Create(company)
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal("1"))

			resp = []byte("1")
			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(resp))
		})
	})

	Context("UpdateCompany", func() {

		var (
			company = newFixtureCompany1()
			c, _    = json.Marshal(company)
			resp []byte
		)

		It("Should return an error because can't convert from json to company", func() {
			useCase.UpdateCompanyMock.Expect([]byte("c")).Return(nil, er)
			response, err := service.UpdateCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))
		})

		It("Should return error if Update method returns error", func() {
			useCase.UpdateCompanyMock.Expect(c).Return(nil, er)
			rep.UpdateMock.Expect(company).Return("", er)

			response, err := service.UpdateCompany(c)

			comp, err2 := rep.Update(company)
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))

		})

		It("Should return nil error", func() {

			useCase.UpdateCompanyMock.Expect(c).Return(resp, nil)
			rep.UpdateMock.Expect(company).Return("1", nil)

			response, err := service.UpdateCompany(c)

			comp, err2 := rep.Update(company)
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal("1"))

			Expect(err).NotTo(HaveOccurred())
			resp = []byte("1")
			Expect(response).To(Equal(resp))
		})
	})

	Context("DeleteCompany", func() {

		var resp []byte
		It("Should return an error because can't convert string to int", func() {
			useCase.DeleteCompanyMock.Expect([]byte("c")).Return(nil, er)
			response, err := service.DeleteCompany([]byte("c"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))
		})

		It("Should return error if Delete method returns error", func() {
			useCase.DeleteCompanyMock.Expect([]byte("6")).Return(nil, er)
			rep.DeleteMock.Expect(int64(6)).Return("", er)

			response, err := service.DeleteCompany([]byte("6"))
			Expect(err).To(HaveOccurred())
			Expect(response).To(Equal(resp))



			comp, err2 := rep.Delete(int64(6))
			Expect(err2).To(HaveOccurred())
			Expect(comp).To(Equal(""))
		})

		It("Should return nil error", func() {
			useCase.DeleteCompanyMock.Expect([]byte("1")).Return(resp, nil)
			rep.DeleteMock.Expect(int64(1)).Return("", nil)

			response, err := service.DeleteCompany([]byte("1"))

			comp, err2 := rep.Delete(int64(1))
			Expect(err2).NotTo(HaveOccurred())
			Expect(comp).To(Equal(""))

			Expect(err).NotTo(HaveOccurred())
			Expect(response).To(Equal(resp))

		})
	})

})



