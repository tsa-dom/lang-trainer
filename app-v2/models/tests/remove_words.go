package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

// These methods needs testing
func RemoveWords() {
	words := models.Words{}

	Describe("Word ids are given", func() {

		Context("Words to be removed are owned by correct user, then", func() {

			It("words are successfully removed", func() {
				ids := []int{4, 5, 7}
				err := words.RemoveWords(4, &ids)
				Expect(err).To(BeNil())
				Expect(ids).To(Equal([]int{4, 5, 7}))
			})

		})

		Context("Words are owned by multiple users, then", func() {

			It("words are not removed", func() {
				ids := []int{3, 5, 7}
				err := words.RemoveWords(4, &ids)
				Expect(err.Error()).To(ContainSubstring("ids not match"))
				Expect(ids).To(Equal([]int{}))
			})

		})

		Context("Some of words does not exist, then", func() {

			It("words are not removed", func() {
				ids := []int{4, 5, 100}
				err := words.RemoveWords(4, &ids)
				Expect(err.Error()).To(ContainSubstring("ids not match"))
				Expect(ids).To(Equal([]int{}))
			})

		})

	})
}
