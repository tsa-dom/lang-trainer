package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

// Modifying word items contains a bug. There should be possibility to remove word items.

// This method needs more testing
func ModifyWord() {
	words := models.Words{}
	wordItems := []models.WordItem{
		{Name: "addition1", Description: "addition desc"},
		{Id: 3, Name: "modification1", Description: "modification desc"},
	}
	expectedItems := []models.WordItem{
		{Id: 10, Name: "addition1", Description: "addition desc"},
		{Id: 3, Name: "modification1", Description: "modification desc"},
	}

	Describe("A word to be modified is given", func() {

		Context("The word is modified by valid user, then", func() {

			It("the word is successfully modified", func() {
				word := models.Word{Id: 2, Name: "new name", Description: "new description", Items: wordItems}
				err := words.Modify(2, &word)
				Expect(err).To(BeNil())
				Expect(word).To(Equal(models.Word{Id: 2, OwnerId: 2, Name: "new name", Description: "new description", Items: expectedItems}))
			})

		})

		/* Context("A word item is going to be removed from a the word, then", func() {

			It("the word is successfully modified", func() {
				word := models.Word{Id: 2, Name: "new name", Description: "new description", Items{}}
			})

		}) */

	})
}
