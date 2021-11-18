package router_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	routes "github.com/tsa-dom/lang-trainer/app/routes/unauthorized"
	utils "github.com/tsa-dom/lang-trainer/app/testutils"
)

var _ = Describe("Visitor", func() {

	BeforeEach(func() {
		conn.InitTestDb()
	})

	Describe("Visitor tries to log in", func() {

		Context("Correct credentials are given, then", func() {

			It("visitor is not logged in", func() {
				bodyReader := strings.NewReader(`{
					"username": "scrum",
					"password": "hash2"	
				}`)
				response := utils.HttpRecorder(
					routes.Login,
					bodyReader,
					nil,
					"",
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`"username":"scrum"`))
				Expect(response.Result().Status).To(Equal("202 Accepted"))
			})

		})

		Context("Correct username and wrong password are given, then", func() {

			It("visitor is not logged in", func() {
				bodyReader := strings.NewReader(`{
					"username": "scrum",
					"password": "wrong"	
				}`)
				response := utils.HttpRecorder(
					routes.Login,
					bodyReader,
					nil,
					"",
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`{"error":"password and hash not match"}`))
				Expect(response.Result().Status).To(Equal("400 Bad Request"))
			})

		})

		Context("Incorrect username is given, then", func() {

			It("visitor is not logged in", func() {
				bodyReader := strings.NewReader(`{
					"username": "nothere",
					"password": "maybecorrect"	
				}`)
				response := utils.HttpRecorder(
					routes.Login,
					bodyReader,
					nil,
					"",
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`"error":"sql: no rows in result set"`))
				Expect(response.Result().Status).To(Equal("400 Bad Request"))
			})

		})

	})

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestUnauthorized(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unauthorized route suite")
}
