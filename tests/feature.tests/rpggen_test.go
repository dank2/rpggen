package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/dank2/rpggen/internal/dnd/npc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	binaryName = "rpggen"
)

func TestNpcGeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rpggen Feature Suite")
}

var _ = Describe("Generate a npc", func() {
	Describe("When generating a NPC", func() {
		dir, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())

		cmd := exec.Command(path.Join(dir, binaryName))

		output, err := cmd.CombinedOutput()
		Expect(err).NotTo(HaveOccurred())

		It("Then a npc should be generated", func() {
			validJSON := json.Valid(output)
			Expect(validJSON).To(BeTrue())
		})

		It("Then name should be set", func() {
			var generatedNpc npc.Npc
			err := json.Unmarshal(output, &generatedNpc)
			Expect(err).NotTo(HaveOccurred())
			Expect(generatedNpc.Name).ToNot(BeEmpty())
		})

		// It("Then race should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Race).ToNot(BeEmpty())
		// })

		// It("Then gender should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Gender).ToNot(BeEmpty())
		// })

		// It("Then personality should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Personality).ToNot(BeEmpty())
		// })

		// It("Then appearance should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Appearance).ToNot(BeEmpty())
		// })

		// It("Then dream should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Dream).ToNot(BeEmpty())
		// })

		// It("Then flaw should be set", func() {
		// 	var generatedNpc npc.Npc
		// 	err := json.Unmarshal(output, &generatedNpc)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(generatedNpc.Flaw).ToNot(BeEmpty())
		// })
	})

})

func TestMain(m *testing.M) {
	err := os.Chdir("../..")
	if err != nil {
		fmt.Printf("Unable to change directory: %v\n", err)
		os.Exit(1)
	}

	make := exec.Command("make")
	err = make.Run()
	if err != nil {
		fmt.Printf("Unable to create binary: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
