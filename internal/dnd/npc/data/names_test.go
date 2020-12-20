package data

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNpcNames(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rpggen Name specification Suite")
}

var _ = Describe("Name generation", func() {
	Context("When getting a name for a human male", func() {
		name := GetName("Male", "Human")

		It("Then the name should be from male table", func() {
			Expect(name).NotTo(BeEmpty())
			Expect(name).Should(BeElementOf(NamesNorseMale))
		})
	})

	Context("When getting a name for a human female", func() {
		name := GetName("Female", "Human")

		It("Then the name should be from female table", func() {
			Expect(name).NotTo(BeEmpty())
			Expect(name).Should(BeElementOf(NamesNorseFemale))
		})
	})

	Context("When getting a name for a half-orc male", func() {
		name := GetName("Male", "Half-orc")

		It("Then the name should be from half-orc male table", func() {
			Expect(name).NotTo(BeEmpty())
			Expect(name).Should(BeElementOf(NamesHalforcMale))
		})
	})

	Context("When getting a name for a half-orc female", func() {
		name := GetName("Female", "Half-orc")

		It("Then the name should be from half-orc female table", func() {
			Expect(name).NotTo(BeEmpty())
			Expect(name).Should(BeElementOf(NamesHalforcFemale))
		})
	})

	Context("When getting a name for a dragonborn male", func() {
		name := GetName("Male", "Dragonborn")

		It("Then the name should be from dragonborn male table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesDragonbornMale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesDragonbornClan))
		})
	})

	Context("When getting a name for a dragonborn female", func() {
		name := GetName("Female", "Dragonborn")

		It("Then the name should be from dragonborn female table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesDragonbornFemale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesDragonbornClan))
		})
	})

	Context("When getting a name for a dwarf male", func() {
		name := GetName("Male", "Dwarf")

		It("Then the name should be from dwarf male table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesDwarfMale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesDwarfClan))
		})
	})

	Context("When getting a name for a dwarf female", func() {
		name := GetName("Female", "Dwarf")

		It("Then the name should be from dwarf female table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesDwarfFemale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesDwarfClan))
		})
	})

	Context("When getting a name for a halfling male", func() {
		name := GetName("Male", "Halfling")

		It("Then the name should be from halfling male and family table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndFamily := strings.Split(name, " ")
			Expect(nameAndFamily[0]).Should(BeElementOf(NamesHalflingMale))
			Expect(nameAndFamily[1]).Should(BeElementOf(NamesHalflingFamily))
		})
	})

	Context("When getting a name for a halfling female", func() {
		name := GetName("Female", "Halfling")

		It("Then the name should be from halfling female and family table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndFamily := strings.Split(name, " ")
			Expect(nameAndFamily[0]).Should(BeElementOf(NamesHalflingFemale))
			Expect(nameAndFamily[1]).Should(BeElementOf(NamesHalflingFamily))
		})
	})

	Context("When getting a name for a gnome male", func() {
		name := GetName("Male", "Gnome")

		It("Then the name should be from gnome male table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesGnomeMale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesGnomeClan))
		})
	})

	Context("When getting a name for a gnome female", func() {
		name := GetName("Female", "Gnome")

		It("Then the name should be from gnome female table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " of clan ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesGnomeFemale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesGnomeClan))
		})
	})

	Context("When getting a name for a elf male", func() {
		name := GetName("Male", "Elf")

		It("Then the name should be from elf male table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesElfMale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesElfFamily))
		})
	})

	Context("When getting a name for a elf female", func() {
		name := GetName("Female", "Elf")

		It("Then the name should be from elf female table", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndClan := strings.Split(name, " ")
			Expect(nameAndClan[0]).Should(BeElementOf(NamesElfFemale))
			Expect(nameAndClan[1]).Should(BeElementOf(NamesElfFamily))
		})
	})

	Context("When getting a name for a half-elf male", func() {
		name := GetName("Male", "Elf")

		It("Then the name should be from either the elf male table plus clan or human male", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndFamily := strings.Split(name, " ")
			humanAndElfNamesMale := append(NamesNorseMale, NamesElfMale...)
			Expect(nameAndFamily[0]).Should(BeElementOf(humanAndElfNamesMale))
			if len(nameAndFamily) > 1 {
				Expect(nameAndFamily[1]).Should(BeElementOf(NamesElfFamily))
			}
		})
	})

	Context("When getting a name for a half-elf female", func() {
		name := GetName("Female", "Elf")

		It("Then the name should be from either the elf female table plus clan or human female", func() {
			Expect(name).NotTo(BeEmpty())
			nameAndFamily := strings.Split(name, " ")
			humanAndElfNamesFemale := append(NamesNorseFemale, NamesElfFemale...)
			Expect(nameAndFamily[0]).Should(BeElementOf(humanAndElfNamesFemale))
			if len(nameAndFamily) > 1 {
				Expect(nameAndFamily[1]).Should(BeElementOf(NamesElfFamily))
			}
		})
	})
})
