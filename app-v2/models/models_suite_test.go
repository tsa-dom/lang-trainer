package models_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	conn "github.com/tsa-dom/lang-trainer/app-v2/db"
	. "github.com/tsa-dom/lang-trainer/app-v2/models/tests"
)

var _ = Describe("Models", func() {

	BeforeEach(func() {
		conn.InitTestDb()
	})

	CreateGroup()
	MiscGroups()
	CreateWord()
	RemoveGroups()
	ModifyWord()
	RemoveWords()

	AfterEach(func() {
		conn.ClearTestDb()
	})

})

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}
