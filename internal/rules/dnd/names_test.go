package dnd

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Name generation", func() {
	Describe("Given generating names", func() {
		Context("When given male as gender", func() {
			name := GetName("male")
			name = name
			It("Then the name should be from male table", func() {
				Fail("Test not implemented")
			})
		})
	})
})
