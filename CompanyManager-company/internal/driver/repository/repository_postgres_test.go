package repository

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/entity"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/usecase"
	"testing"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}


var _ = Describe("Repository", func() {

	var (
		//db sql.DB
		conn *postgresRepo
		c *entity.Company
	)


	BeforeEach(func() {
		PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"localhost", "5432", "postgres", "mypassword", "company_manager")
		db, _ := sql.Open("postgres", PsqlInfo)
		db.Ping()
		conn = NewPostgresRepository(db)
		c = &entity.Company{}
	})

	Context("Get", func() {

		It("Should return company with index 1", func() {
			comp, err := conn.Get(1)
			Expect(err).NotTo(HaveOccurred())
			c.ID = 1
			c.Name = "Company1"
			c.Legalform = "OOO"
			Expect(comp).To(Equal(c))
		})

		It("Should return an error because can't connect to DB", func() {
			db, err := sql.Open("postgres", "bad connection")
			conn = NewPostgresRepository(db)
			comp, err := conn.Get(1)
			Expect(err).To(HaveOccurred())
			c = nil
			Expect(comp).To(Equal(c))

		})
	})


	Context("Create and Delete", func() {

		It("Should crete company with index new index ", func() {
			c.Name = "Company4"
			c.Legalform = "OOO"
			comp, err := conn.Create(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(comp).To(Equal(comp))
			val, err := usecase.StringToInt64(comp)
			val1, err := conn.Delete(val)
			Expect(err).NotTo(HaveOccurred())
			Expect(val1).To(Equal("Successful delete"))
		})


		It("Should return an error because can't connect to DB", func() {
			db, err := sql.Open("postgres", "bad connection")
			conn = NewPostgresRepository(db)
			comp, err := conn.Create(c)
			Expect(err).To(HaveOccurred())
			c = nil
			Expect(comp).To(Equal(""))

		})

		It("Should return an error because can't connect to DB", func() {
			db, err := sql.Open("postgres", "bad connection")
			conn = NewPostgresRepository(db)
			comp, err := conn.Delete(6)
			Expect(err).To(HaveOccurred())
			c = nil
			Expect(comp).To(Equal(""))

		})
	})

	Context("Update", func() {

		It("Should update company with index 5", func() {
			c.ID = 4
			c.Name = "Company4"
			c.Legalform = "OOO"
			comp, err := conn.Update(c)
			Expect(err).NotTo(HaveOccurred())
			Expect(comp).To(Equal("Successful update"))
			c.ID = 4
			c.Name = "Company5"
			c.Legalform = "ZAO"
			conn.Update(c)
		})

		It("Should return an error because can't connect to DB", func() {
			db, err := sql.Open("postgres", "bad connection")
			conn = NewPostgresRepository(db)
			comp, err := conn.Update(c)
			Expect(err).To(HaveOccurred())
			c = nil
			Expect(comp).To(Equal(""))

		})
	})
})
