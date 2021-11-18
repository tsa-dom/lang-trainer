package router_test

import (
	"encoding/json"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	login "github.com/tsa-dom/lang-trainer/app/routes/unauthorized"
	routes "github.com/tsa-dom/lang-trainer/app/routes/user"
	utils "github.com/tsa-dom/lang-trainer/app/testutils"
)

type Token struct {
	Token string `json:"token"`
}

var _ = Describe("User", func() {

	var token Token

	BeforeEach(func() {
		conn.InitTestDb()
		bodyReader := strings.NewReader(`{
			"username": "scrum",
			"password": "hash2"
		}`)
		response := utils.HttpRecorder(
			login.Login,
			bodyReader,
			nil,
			"",
		)
		json.Unmarshal(response.Body.Bytes(), &token)
	})

	Describe("User is logged in", func() {

		Context("User wants their details against the token", func() {

			It("user details are given", func() {
				bodyReader := strings.NewReader("{}")
				response := utils.HttpRecorder(
					routes.GetUser,
					bodyReader,
					router.AuthorizeUser(),
					token.Token,
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`{"user":{"id":2,"username":"scrum","privileges":"admin","token":""}}`))
			})

		})

	})

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User route suite")
}
