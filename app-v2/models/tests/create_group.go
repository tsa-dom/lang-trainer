package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

func CreateGroup() {
	groups := models.Groups{}

	Describe("Valid group details are given", func() {

		Context("Owner for a new group exists, then", func() {

			It("group is successfully created", func() {
				group := models.Group{Name: "New Group", Description: "This is a new group"}
				err := groups.Create(3, &group)
				Expect(err).To(BeNil())
				Expect(group).To(Equal(models.Group{
					Id:          7,
					OwnerId:     3,
					Name:        "New Group",
					Description: "This is a new group",
				}))
			})

		})

		Context("Owner for a new group does not exist, then", func() {

			It("group is not created", func() {
				group := models.Group{Name: "New Group 2", Description: "this is a new group"}
				err := groups.Create(100, &group)
				Expect(err.Error()).To(ContainSubstring("insert or update on table \"groups\" violates foreign key constraint \"groups_owner_id_fkey\""))
				Expect(group).To(Equal(models.Group{}))
			})

		})

	})

	Describe("Invalid group details are given", func() {

		Context("Group name is empty, then", func() {

			It("group is not created", func() {
				group := models.Group{Name: "", Description: "this name is empty"}
				err := groups.Create(2, &group)
				Expect(err.Error()).To(ContainSubstring("pq: new row for relation \"groups\" violates check constraint \"groups_name_check\""))
				Expect(group).To(Equal(models.Group{}))
			})

		})

	})
}
