package dnd

import (
	"bytes"
	"testing"

	"github.com/dank2/rpggen/rpggen/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		It("should not have generated an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have generated an npc", func() {
			Expect(out.String()).To(ContainSubstring("Generated NPC:"))
		})
	})
})
