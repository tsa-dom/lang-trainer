package models_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	users "github.com/tsa-dom/lang-trainer/app/models/users"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

var _ = Describe("User", func() {

	BeforeEach(func() {
		conn.InitTestDb()
	})

	Describe("Valid user details are given", func() {

		Context("The same username is not in db", func() {

			It("user is successfully created", func() {
				user := g.User{Username: "Admin", PasswordHash: "thisishash", Privileges: "admin"}
				createdUser, err := users.CreateUser(user)
				Expect(err).To(BeNil())
				Expect(createdUser).To(Equal(g.User{Id: createdUser.Id, PasswordHash: "thisishash", Username: "Admin", Privileges: "admin"}))
			})

		})

		Context("The same username is in db", func() {

			It("user is not created", func() {
				user := g.User{Username: "exists", PasswordHash: "thisishash", Privileges: "admin"}

				createdUser, err := users.CreateUser(user)
				Expect(err.Error()).To(ContainSubstring("pq: duplicate key value violates unique constraint \"users_username_key\""))
				Expect(createdUser).To(Equal(g.User{}))
			})

		})

	})

	Describe("An iteger id is given", func() {

		Context("There is user with given id in db", func() {

			It("User details are returned", func() {
				user, err := users.GetUserByUsername("desirable")
				Expect(user).To(Equal(g.User{Id: 4, Username: "desirable", PasswordHash: "hash4", Privileges: "desire"}))
				Expect(err).To(BeNil())
			})

		})

		Context("There is not user with given id in db", func() {

			It("User details are not returned and error is given", func() {
				user, err := users.GetUserByUsername("notfound")
				Expect(user).To(Equal(g.User{}))
				Expect(err.Error()).To(ContainSubstring("sql: no rows in result set"))
			})

		})

	})

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}
