package routes_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	routes "github.com/tsa-dom/lang-trainer/app/routes/admin"
	utils "github.com/tsa-dom/lang-trainer/app/testutils"
)

var _ = Describe("Admin", func() {

	BeforeEach(func() {
		conn.InitTestDb()
	})

	Describe("Creating an user and valid authorization header is given.", func() {

		Context("Valid user details are given, then", func() {

			It("new user is created", func() {
				bodyReader := strings.NewReader(`{
					"username": "Test",
					"password": "Secret",
					"privileges": "admin"
				}`)
				response := utils.HttpRecorder(
					routes.SignNewUser,
					bodyReader,
					router.AuthorizeAdmin(),
					"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicHJpdmlsZWdlcyI6ImFkbWluIiwidGltZSI6IjIwMjEtMTEtMTdUMTc6MTg6MDMuOTY0NzE4M1oiLCJ1c2VybmFtZSI6IkFkbWluIn0.N1dgkSCzyKebKzywIlMGHZ1qanK-Lu7IJQ4R55PfL4E")
				body := response.Body
				Expect(body).To(ContainSubstring(`"username":"Test"`))
				Expect(body).NotTo(ContainSubstring(`assword`))
				Expect(response.Result().Status).To(Equal("202 Accepted"))
			})

		})

	})

	Describe("Creating an user and invalid authorization header is given.", func() {

		Context("Valid user details are given, then", func() {

			It("new user is not created", func() {
				bodyReader := strings.NewReader(`{
					"username": "Test",
					"password": "Secret",
					"privileges": "admin"
				}`)
				response := utils.HttpRecorder(
					routes.SignNewUser,
					bodyReader,
					router.AuthorizeAdmin(),
					"fyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicHJpdmlsZWdlcyI6ImFkbWluIiwidGltZSI6IjIwMjEtMTEtMTdUMTc6MTg6MDMuOTY0NzE4M1oiLCJ1c2VybmFtZSI6IkFkbWluIn0.N1dgkSCzyKebKzywIlMGHZ1qanK-Lu7IJQ4R55PfL4E")
				body := response.Body
				Expect(body).To(ContainSubstring(`invalid Authorization token`))
				Expect(response.Result().Status).To(Equal("403 Forbidden"))
			})

		})

	})

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Admin route suite")
}
