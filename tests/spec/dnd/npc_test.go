package dnd

import (
	"bytes"
	"strings"
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

		It("should have a race", func() {
			Expect(npc.Race).To(BeElementOf(data.Races))
		})

		It("should have an appearance", func() {
			split := strings.Split(npc.Appearance, ", ")
			Expect(len(split)).To(BeNumerically(">=", 2))
			Expect(split[0]).To(BeElementOf([]string{"Male", "Female"}))
			Expect(strings.Join(split[1:], ", ")).To(BeElementOf(data.Appearances))
		})

		It("should have a personality", func() {
			Expect(npc.Mannerism).To(BeElementOf(data.Mannerism))
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
