package users_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app/models"
	"github.com/tsa-dom/lang-trainer/app/models/users"
)

var _ = Describe("User", func() {

	var wg sync.WaitGroup

	BeforeEach(func() {
		models.InitDB("../../../schema.sql")
		db := models.GetDbConnection()
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

	Describe("Valid user details are given", func() {

		Context("The same username is not in db", func() {

			It("user is successfully created", func() {
				user := users.User{Username: "Admin", PasswordHash: "thisishash", Privileges: "admin"}
				createdUser, err := users.CreateUser(user)
				Expect(err).To(BeNil())
				Expect(createdUser).To(Equal(users.User{Id: 5, PasswordHash: "thisishash", Username: "Admin", Privileges: "admin"}))
			})

		})

		Context("The same username is in db", func() {

			It("user is not created", func() {
				user := users.User{Username: "exists", PasswordHash: "thisishash", Privileges: "admin"}

				createdUser, err := users.CreateUser(user)
				Expect(err.Error()).To(ContainSubstring("pq: duplicate key value violates unique constraint \"users_username_key\""))
				Expect(createdUser).To(Equal(users.User{}))
			})

		})

	})

	Describe("An iteger id is given", func() {

		Context("There is user with given id in db", func() {

			It("User details are returned", func() {
				user, err := users.GetUserByUsername("desirable")
				Expect(user).To(Equal(users.User{Id: 4, Username: "desirable", PasswordHash: "hash4", Privileges: "desire"}))
				Expect(err).To(BeNil())
			})

		})

		Context("There is not user with given id in db", func() {

			It("User details are not returned and error is given", func() {
				user, err := users.GetUserByUsername("notfound")
				Expect(user).To(Equal(users.User{}))
				Expect(err.Error()).To(ContainSubstring("sql: no rows in result set"))
			})

		})

	})

	AfterEach(func() {
		db := models.GetDbConnection()
		defer db.Close()
		clear := `
			DROP TABLE Users CASCADE;
			DROP TABLE Words CASCADE;
			DROP TABLE WordItems CASCADE;
			DROP TABLE Groups CASCADE;
			DROP TABLE GroupLinks CASCADE;
		`
		_, err := db.Exec(clear)
		if err != nil {
			log.Println(err)
		}
	})

})

func TestUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}
