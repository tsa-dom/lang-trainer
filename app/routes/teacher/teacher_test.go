package routes_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	routes "github.com/tsa-dom/lang-trainer/app/routes/teacher"
	utils "github.com/tsa-dom/lang-trainer/app/testutils"
)

var teacherToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwicHJpdmlsZWdlcyI6InRlYWNoZXIiLCJ0aW1lIjoiMjAyMS0xMS0xOFQxMzo0NToyOC45MTY1NDM3WiIsInVzZXJuYW1lIjoiVGVhY2hlciJ9.2Y5h62bW9Y3ulC4TvYjxm8MpJ_fYyVmTpsxu5XB2pZE"
var studentToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywicHJpdmlsZWdlcyI6InN0dWRlbnQiLCJ0aW1lIjoiMjAyMS0xMS0xOFQxMzo1NzozOS41OTI2OTM0WiIsInVzZXJuYW1lIjoiU3R1ZGVudCJ9.gan47F6rvzmf5wXHSId_d73Z553GTJ6AJ9J-ZW085QI"

var _ = Describe("Teacher", func() {

	BeforeEach(func() {
		conn.InitTestDb()
	})

	Describe("Valid details are given.", func() {

		Context("Teacher is logged in, then", func() {

			It("a group is successfully added", func() {
				bodyReader := strings.NewReader(`{
					"name": "New group",
					"description": "This is awesome"
				}`)
				response := utils.HttpRecorder(
					routes.AddGroup,
					bodyReader,
					router.AuthorizeTeacher(),
					teacherToken,
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`{"group":{"id":3,"ownerId":2,"name":"New group","description":"This is awesome"}`))
				Expect(response.Result().Status).To(Equal("202 Accepted"))
			})

			It("a word is successfully added to existing group", func() {
				bodyReader := strings.NewReader(`{
					"name": "A new word",
					"description": "This is awesome word",
					"groupId": 2,
					"items": [{
						"name": "Item1",
						"description": "Item1desc"
					}, {
						"name": "Item2",
						"description": "Item2desc"
					}]
				}`)
				response := utils.HttpRecorder(
					routes.AddWordToGroup,
					bodyReader,
					router.AuthorizeTeacher(),
					teacherToken,
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`{"word":{"id":5,"ownerId":2,"name":"A new word","description":"This is awesome word","groupId":2,"items":[{"id":10,"name":"Item1","description":"Item1desc"},{"id":11,"name":"Item2","description":"Item2desc"}]}}`))
				Expect(response.Result().Status).To(Equal("202 Accepted"))
			})

		})

		Context("Student is logged in, then", func() {

			It("a group is not created", func() {
				bodyReader := strings.NewReader(`{
					"name": "New group",
					"description": "This is awesome"
				}`)
				response := utils.HttpRecorder(
					routes.AddGroup,
					bodyReader,
					router.AuthorizeTeacher(),
					studentToken,
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`are you teacher?`))
				Expect(response.Result().Status).To(Equal("403 Forbidden"))
			})

		})

	})

	Describe("Invalid details are given", func() {

		Context("Teacher is logged in", func() {

			It("a word is not assigned to a group if group does not exist", func() {
				bodyReader := strings.NewReader(`{
					"name": "A new word",
					"description": "This is awesome word",
					"groupId": 100,
					"items": [{
						"name": "Item1",
						"description": "Item1desc"
					}, {
						"name": "Item2",
						"description": "Item2desc"
					}]
				}`)
				response := utils.HttpRecorder(
					routes.AddWordToGroup,
					bodyReader,
					router.AuthorizeTeacher(),
					teacherToken,
				)
				body := response.Body
				Expect(body).To(ContainSubstring(`insert or update on table \"grouplinks\" violates foreign key constraint`))
				Expect(response.Result().Status).To(Equal("500 Internal Server Error"))
			})

		})

	})

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestTeacher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Teacher route suite")
}
