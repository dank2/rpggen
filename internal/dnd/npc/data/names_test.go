package data

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNpcNames(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rpggen Name specification Suite")
}

var _ = Describe("Name generation", func() {
	Describe("Given generating names", func() {
		Context("When given male as gender", func() {
			name := GetName("male")

			It("Then the name should be from male table", func() {
				Expect(name).NotTo(BeEmpty())
				Expect(norseNamesMale).Should(ContainElement(name))
			})
		})

		Context("When given female as gender", func() {
			name := GetName("female")

			It("Then the name should be from male table", func() {
				Expect(name).NotTo(BeEmpty())
				Expect(norseNamesFemale).Should(ContainElement(name))
			})
		})
	})
})
