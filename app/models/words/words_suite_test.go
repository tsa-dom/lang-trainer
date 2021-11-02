package words_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app/models"
	"github.com/tsa-dom/lang-trainer/app/models/words"
)

var _ = Describe("Word", func() {

	BeforeEach(func() {
		models.InitDB("../../../schema.sql")
		db := models.GetDbConnection()
		defer db.Close()

		path := filepath.Join("../../../testdata.sql")
		c, ioErr := ioutil.ReadFile(path)
		if ioErr != nil {
			log.Fatal("Error loading testdata.sql file")
		}
		sql := string(c)
		db.Exec(sql)
	})

	Describe("Valid group details are given", func() {

		Context("Owner for a new group exists", func() {

			It("group is successfully created", func() {
				group := words.Group{OwnerId: 3, Name: "new group", Description: "new group description"}
				group, err := words.CreateGroup(group)
				Expect(err).To(BeNil())
				Expect(group).To(Equal(words.Group{
					Id:          1,
					OwnerId:     3,
					Name:        "new group",
					Description: "new group description",
				}))
			})

		})

		Context("Owner for a new group does not exits", func() {

			It("group is not created", func() {
				group := words.Group{OwnerId: 100, Name: "new group2", Description: "new group description2"}
				group, err := words.CreateGroup(group)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"groups\" violates foreign key constraint"))
				Expect(group).To(Equal(words.Group{Id: 0, OwnerId: 0, Name: "", Description: ""}))
			})

		})

	})

	Describe("Word items are given", func() {

		items := []words.WordItem{
			{Name: "item1", Description: "this is item1"},
			{Name: "item2", Description: "this is item2"},
			{Name: "item3", Description: "this is item3"},
		}

		Context("Items are added to existing word", func() {

			It("word items are succesfully added", func() {
				err := words.AddItemsToWord(3, items)
				Expect(err).To(BeNil())
			})

		})

		Context("Items are added to nonexisting word", func() {

			It("word items are not added", func() {
				err := words.AddItemsToWord(100, items)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"worditems\" violates foreign key constraint"))
			})

		})

	})

	Describe("Word id is given", func() {

		Context("Word is fetched from database with valid id", func() {

			It("correct word is fetched", func() {
				expected := words.Word{
					Id:          2,
					OwnerId:     2,
					Name:        "This is word2",
					Description: "Awesome text2",
					Items: []words.WordItem{
						{
							Id:          1,
							Name:        "Item1",
							Description: "item desc1",
						},
						{
							Id:          2,
							Name:        "Item2",
							Description: "item desc2",
						},
						{
							Id:          3,
							Name:        "Item3",
							Description: "item desc3",
						},
					},
				}
				word, err := words.GetWordById(2)
				Expect(err).To(BeNil())
				Expect(word).To(Equal(expected))
			})

		})

		Context("Word is fetched from database with invalid id", func() {

			It("word is not returned", func() {
				word, err := words.GetWordById(100)
				Expect(err.Error()).To(ContainSubstring("sql: no rows in result set"))
				Expect(word).To(Equal(words.Word{Id: 0, OwnerId: 0, Name: "", Description: "", Items: nil}))
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

func TestWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Word Suite")
}
