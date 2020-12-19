package dnd

import (
	"bytes"
	"testing"

	"github.com/dank2/rpggen/internal/dnd/npc"
	"github.com/dank2/rpggen/internal/dnd/npc/data"
	"github.com/dank2/rpggen/rpggen/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestNpcGeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate dnd npc")
}

var _ = Describe("Dnd NPC spec tests", func() {
	Context("when generating a default dnd npc", func() {

		var out bytes.Buffer
		var errOut bytes.Buffer
		command := cmd.NewRootCommand(&out, &errOut)
		command.SetArgs([]string{"dnd", "npc"})
		err := command.Execute()

		npc := new(npc.Npc)
		yaml.Unmarshal([]byte(out.String()), npc)

		It("should not have generated an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have a gender", func() {
			Expect(npc.Gender).To(BeElementOf([]string{"male", "female"}))
		})

		It("should have a race", func() {
			Expect(npc.Race).To(BeElementOf(getValues(data.Races)))
		})

		It("should have an appearance", func() {
			appearances := getValues(data.Appearances)
			Expect(npc.Appearance).To(BeElementOf(appearances))
		})

		It("should have a personality", func() {
			mannerism := getValues(data.Mannerism)
			Expect(npc.Mannerism).To(BeElementOf(mannerism))
		})

		It("should have a bond", func() {
			Expect(npc.Bond).ToNot(BeEmpty())
		})

		It("should have a flaw", func() {
			Expect(npc.Flaw).To(BeElementOf(data.Flaws))
		})
	})
})

func getValues(m map[int]string) []string {
	i := 0
	values := make([]string, len(m))
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}
