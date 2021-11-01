package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app/testutils"
)

var _ = Describe("User", func() {

	BeforeEach(func() {
		testutils.InitTestDb()
	})

	Describe("Valid user details are given", func() {

		Context("The same username is not in db", func() {

			It("user is successfully created", func() {
				db := testutils.TestDbConnection()
				defer db.Close()
				user := User{Username: "Admin", PasswordHash: "thisishash", Priviledges: "admin"}

				err := createUser(db, user)
				Expect(err).To(BeNil())
			})

		})

		Context("The same username is in db", func() {

			It("user is not created", func() {
				db := testutils.TestDbConnection()
				defer db.Close()
				user := User{Username: "exists", PasswordHash: "thisishash", Priviledges: "admin"}

				err := createUser(db, user)
				Expect(err).NotTo(BeNil())
			})

		})

	})

	Describe("An iteger id is given", func() {

		Context("There is user with given id in db", func() {

			It("User details are returned", func() {
				db := testutils.TestDbConnection()
				defer db.Close()

				user, err := getUserByUsername(db, "desirable")
				Expect(user).To(Equal(User{Id: 4, Username: "desirable", PasswordHash: "hash4", Priviledges: "desire"}))
				Expect(err).To(BeNil())
			})

		})

		Context("There is not user with given id in db", func() {

			It("User details are not returned and error is given", func() {
				db := testutils.TestDbConnection()
				defer db.Close()

				user, err := getUserByUsername(db, "notfound")
				Expect(user).To(Equal(User{}))
				Expect(err).NotTo(BeNil())
			})

		})

	})

	AfterEach(func() {
		testutils.ClearDb()
	})

})

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}
