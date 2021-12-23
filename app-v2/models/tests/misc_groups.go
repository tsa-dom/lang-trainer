package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsa-dom/lang-trainer/app-v2/models"
)

func MiscGroups() {
	groups := models.Groups{}
	returnedGroups := []models.Group{
		{Id: 1, Name: "Group", Description: "This is awesome", OwnerId: 2},
		{Id: 2, Name: "Propably empty", Description: "This should be empty", OwnerId: 2},
	}

	Describe("The owner for groups is given", func() {

		Context("Groups are fetched from the database, then", func() {

			It("groups are successfully returned", func() {
				allGroups := []models.Group{}
				err := groups.GetAll(2, &allGroups)
				Expect(err).To(BeNil())
				Expect(allGroups).To(Equal(returnedGroups))
			})

		})

	})

}
