package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

func RemoveGroups() {
	groups := models.Groups{}

	Describe("Group ids are given", func() {

		Context("Groups to be removed are owned by correct user, then", func() {

			It("groups are successfully removed", func() {
				ids := []int{4, 6}
				err := groups.RemoveGroups(1, &ids)
				Expect(err).To(BeNil())
				Expect(ids).To(Equal([]int{4, 6}))
			})

		})

		Context("Groups are owned by multiple users, then", func() {

			It("groups are not removed", func() {
				ids := []int{3, 6}
				err := groups.RemoveGroups(1, &ids)
				Expect(err.Error()).To(ContainSubstring("ids not match"))
				Expect(ids).To(Equal([]int{}))
			})

		})

		Context("Some of groups does not exist, then", func() {

			It("groups are not removed", func() {
				ids := []int{4, 7}
				err := groups.RemoveGroups(1, &ids)
				Expect(err.Error()).To(ContainSubstring("ids not match"))
				Expect(ids).To(Equal([]int{}))
			})

		})

	})

}
