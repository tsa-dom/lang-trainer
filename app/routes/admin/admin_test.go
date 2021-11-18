package routes_test

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	routes "github.com/tsa-dom/lang-trainer/app/routes/admin"
)

func httpRecorder(handler func(c *gin.Context), body io.Reader, token string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.Use(router.AuthorizeAdmin())
	engine.POST("/", handler)

	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", body)
	ctx.Request.Header.Set("Authorization", "Bearer "+token)
	engine.HandleContext(ctx)
	return w
}

var _ = Describe("Admin", func() {

	var wg sync.WaitGroup

	BeforeEach(func() {
		conn.InitTestDb()
		db := conn.GetDbConnection()
		defer db.Close()
		wg.Add(5)
		defer wg.Done()

		path := filepath.Join("../../../testdata.sql")
		c, ioErr := ioutil.ReadFile(path)
		if ioErr != nil {
			log.Fatal("Error loading testdata.sql file")
		}
		sql := string(c)
		db.Exec(sql)
	})

	Describe("Creating an user and valid authorization header is given.", func() {

		Context("Valid user details are given, then", func() {

			It("new user is created", func() {
				bodyReader := strings.NewReader(`{
					"username": "Test",
					"password": "Secret",
					"privileges": "admin"
				}`)
				response := httpRecorder(routes.SignNewUser, bodyReader, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicHJpdmlsZWdlcyI6ImFkbWluIiwidGltZSI6IjIwMjEtMTEtMTdUMTc6MTg6MDMuOTY0NzE4M1oiLCJ1c2VybmFtZSI6IkFkbWluIn0.N1dgkSCzyKebKzywIlMGHZ1qanK-Lu7IJQ4R55PfL4E")
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
				response := httpRecorder(routes.SignNewUser, bodyReader, "fyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicHJpdmlsZWdlcyI6ImFkbWluIiwidGltZSI6IjIwMjEtMTEtMTdUMTc6MTg6MDMuOTY0NzE4M1oiLCJ1c2VybmFtZSI6IkFkbWluIn0.N1dgkSCzyKebKzywIlMGHZ1qanK-Lu7IJQ4R55PfL4E")
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
