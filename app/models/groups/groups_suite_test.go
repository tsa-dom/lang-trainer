package models_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	groups "github.com/tsa-dom/lang-trainer/app/models/groups"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

var _ = Describe("Group", func() {

	var wg sync.WaitGroup

	BeforeEach(func() {
		conn.InitDB("../../../schema.sql")
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

	Describe("Valid group details are given", func() {

		Context("Owner for a new group exists", func() {

			It("group is successfully created", func() {
				group := g.Group{OwnerId: 3, Name: "new group", Description: "new group description"}
				group, err := groups.CreateGroup(group)
				Expect(err).To(BeNil())
				Expect(group).To(Equal(g.Group{
					Id:          3,
					OwnerId:     3,
					Name:        "new group",
					Description: "new group description",
				}))
			})

		})

		Context("Owner for a new group does not exits", func() {

			It("group is not created", func() {
				group := g.Group{OwnerId: 100, Name: "new group2", Description: "new group description2"}
				group, err := groups.CreateGroup(group)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"groups\" violates foreign key constraint"))
				Expect(group).To(Equal(g.Group{Id: 0, OwnerId: 0, Name: "", Description: ""}))
			})

		})

	})

	Describe("Word items are given", func() {

		items := []g.WordItem{
			{Name: "item1", Description: "this is item1"},
			{Name: "item2", Description: "this is item2"},
			{Name: "item3", Description: "this is item3"},
		}

		Context("Items are added to existing word", func() {

			It("word items are succesfully added", func() {
				wordItems, err := groups.AddItemsToWord(3, items)
				Expect(err).To(BeNil())
				log.Println(wordItems)
			})

		})

		Context("Items are added to nonexisting word", func() {

			It("word items are not added", func() {
				wordItems, err := groups.AddItemsToWord(100, items)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"worditems\" violates foreign key constraint"))
				Expect(wordItems).To(Equal([]g.WordItem{}))
			})

		})

	})

	Describe("Word id is given", func() {

		Context("Word is fetched from database with valid id", func() {

			It("correct word is fetched", func() {
				expected := g.Word{
					Id:          2,
					OwnerId:     2,
					Name:        "This is word2",
					Description: "Awesome text2",
					Items: []g.WordItem{
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
				word, err := groups.GetWordById(2)
				Expect(err).To(BeNil())
				Expect(word).To(Equal(expected))
			})

		})

		Context("Word is fetched from database with invalid id", func() {

			It("word is not returned", func() {
				word, err := groups.GetWordById(100)
				Expect(err.Error()).To(ContainSubstring("sql: no rows in result set"))
				Expect(word).To(Equal(g.Word{Id: 0, OwnerId: 0, Name: "", Description: "", Items: nil}))
			})

		})

		Describe("Group id is given", func() {

			Context("Words are fetched from group which contains words", func() {

				It("correct words are returned", func() {
					expected := []g.Word{{
						Id:          2,
						OwnerId:     2,
						Name:        "This is word2",
						Description: "Awesome text2",
						GroupId:     1,
						Items: []g.WordItem{
							{
								Id:          3,
								Name:        "Item3",
								Description: "item desc3",
							},
							{
								Id:          2,
								Name:        "Item2",
								Description: "item desc2",
							},
							{
								Id:          1,
								Name:        "Item1",
								Description: "item desc1",
							},
						},
					},
						{
							Id:          1,
							OwnerId:     2,
							Name:        "This is word",
							Description: "Awesome text",
							GroupId:     1,
							Items: []g.WordItem{
								{
									Id:          5,
									Name:        "Item5",
									Description: "item desc5",
								},
								{
									Id:          4,
									Name:        "Item4",
									Description: "item desc4",
								},
							},
						}}
					sort.SliceStable(expected, func(i, j int) bool {
						return expected[i].Id < expected[j].Id
					})
					words, err := groups.GetWordsInGroup(1)
					sort.SliceStable(words, func(i, j int) bool {
						return words[i].Id < words[j].Id
					})
					Expect(err).To(BeNil())
					Expect(words).To(Equal(expected))
				})

			})
		})

	})

	AfterEach(func() {
		db := conn.GetDbConnection()
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
			log.Panic(err)
		}
	})

})

func TestWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Suite")
}
