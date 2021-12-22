package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	groups "github.com/tsa-dom/lang-trainer/app/models/groups"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

func ModifyGroup() {

	Describe("Valid group details are given", func() {

		Context("A group is modified by its owner, then", func() {

			It("the groups is successfully modified", func() {
				group := g.Group{Id: 2, Name: "modification", Description: "this should work"}
				err := groups.ModifyGroup(2, &group)
				Expect(err.Error()).To(BeNil())
			})

		})

		Context("A group is not modified by its owner, then", func() {

			It("the groups is successfully modified", func() {
				group := g.Group{Id: 2, Name: "modification", Description: "this should work"}
				err := groups.ModifyGroup(2, &group)
				Expect(err.Error()).To(BeNil())
			})

		})

	})

}
