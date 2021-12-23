package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

func CreateWord() {
	words := models.Words{}
	wordItems := []models.WordItem{
		{Name: "item1", Description: "item1desc"},
		{Name: "item2", Description: "item2desc"},
		{Name: "item3", Description: "item3desc"},
	}
	wordItemsWithoutName := []models.WordItem{
		{Name: "item1", Description: "item1desc"},
		{Name: "", Description: "item2desc"},
		{Name: "item3", Description: "item3desc"},
	}
	expectedItems := []models.WordItem{
		{Id: 10, Name: "item1", Description: "item1desc"},
		{Id: 11, Name: "item2", Description: "item2desc"},
		{Id: 12, Name: "item3", Description: "item3desc"},
	}
	notCreatedItems := []models.WordItem{
		{Id: 0, Name: "item1", Description: "item1desc"},
		{Id: 0, Name: "item2", Description: "item2desc"},
		{Id: 0, Name: "item3", Description: "item3desc"},
	}
	notCreatedItems2 := []models.WordItem{
		{Id: 0, Name: "item1", Description: "item1desc"},
		{Id: 0, Name: "", Description: "item2desc"},
		{Id: 0, Name: "item3", Description: "item3desc"},
	}

	Describe("Valid word details are given", func() {

		Context("An owner and a group for a new word exists, then", func() {

			It("the word is successfully created", func() {
				word := models.Word{Name: "Word1", Description: "This is a new word", Items: wordItems}
				err := words.Create(3, &word)
				Expect(err).To(BeNil())
				Expect(word).To(Equal(models.Word{Id: 5, Name: "Word1", OwnerId: 3, Description: "This is a new word", Items: expectedItems}))
			})

		})

		Context("An owner for a new word does not exist, then", func() {

			It("the word is not created", func() {
				word := models.Word{Name: "Word2", Description: "This cannot be created", Items: wordItems}
				err := words.Create(100, &word)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"words\" violates foreign key constraint \"words_owner_id_fkey\""))
				Expect(word).To(Equal(models.Word{Name: "Word2", Description: "This cannot be created", Items: notCreatedItems}))
			})

		})

	})

	Describe("Invalid word details are given", func() {

		Context("A word name is empty, then", func() {

			It("the word is not created", func() {
				word := models.Word{Name: "", Description: "This is empty", Items: wordItems}
				err := words.Create(3, &word)
				Expect(err.Error()).To(ContainSubstring("pq: new row for relation \"words\" violates check constraint \"words_word_check\""))
				Expect(word).To(Equal(models.Word{Description: "This is empty", Items: notCreatedItems}))
			})

		})

		Context("A word contains a word item which name is empty, then", func() {

			It("then word is not created", func() {
				word := models.Word{Name: "Lost", Description: "Item lost", Items: wordItemsWithoutName}
				err := words.Create(3, &word)
				Expect(err.Error()).To(ContainSubstring("pq: new row for relation \"worditems\" violates check constraint \"worditems_word_check\""))
				Expect(word).To(Equal(models.Word{Id: 5, Name: "Lost", Description: "Item lost", Items: notCreatedItems2}))
			})

		})

	})
}
