package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	groups "github.com/tsa-dom/lang-trainer/app/models/groups"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

func CreateGroup() {

	Describe("Valid group details are given", func() {

		Context("Owner for a new group exists, then", func() {

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

		Context("Owner for a new group does not exits, then", func() {

			It("group is not created", func() {
				group := g.Group{OwnerId: 100, Name: "new group2", Description: "new group description2"}
				group, err := groups.CreateGroup(group)
				Expect(err.Error()).To(ContainSubstring("pq: insert or update on table \"groups\" violates foreign key constraint"))
				Expect(group).To(Equal(g.Group{Id: 0, OwnerId: 0, Name: "", Description: ""}))
			})

		})

	})

	Describe("Invalid group details are given", func() {

		Context("Group name is missing, then", func() {

			It("group is not created", func() {
				group := g.Group{OwnerId: 3, Description: "no name"}
				group, err := groups.CreateGroup(group)
				Expect(err.Error()).To(ContainSubstring("pq: new row for relation \"groups\" violates check constraint \"groups_name_check\""))
				Expect(group).To(Equal(g.Group{Id: 0, OwnerId: 0, Name: "", Description: ""}))
			})

		})

		Context("Group name is empty, then", func() {

			It("group is not created", func() {
				group := g.Group{OwnerId: 2, Name: "", Description: "this name is empty"}
				group, err := groups.CreateGroup(group)
				Expect(err.Error()).To(ContainSubstring("pq: new row for relation \"groups\" violates check constraint \"groups_name_check\""))
				Expect(group).To(Equal(g.Group{Id: 0, OwnerId: 0, Name: "", Description: ""}))
			})

		})

	})
}
