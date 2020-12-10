package dnd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/dank2/rpggen/internal/dnd/npc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

var (
	binaryName = "rpggen"
)

func TestNpcGeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate dnd npc")
}

var _ = Describe("Generate a npc", func() {
	wd, err := os.Getwd()
	Expect(err).ToNot(HaveOccurred())

	rootDir := filepath.Join(wd, "..", "..", "..")
	err = build(rootDir)
	Expect(err).ToNot(HaveOccurred())

	Context("When generating an NPC", func() {
		output, cmdErr := runRpgGen(rootDir, "dnd", "npc")

		generatedNpc := new(npc.Npc)
		marshalErr := yaml.Unmarshal([]byte(output), generatedNpc)

		It("should not output an error", func() {
			Expect(cmdErr).To(BeNil())
		})

		It("should give me some output", func() {
			Expect(output).ToNot(BeEmpty())
		})

		It("should generate a name", func() {
			Expect(marshalErr).To(BeNil())
			Expect(generatedNpc.Name).ToNot(BeEmpty())
		})
	})

})

func runRpgGen(wd string, args ...string) (string, error) {
	binary := filepath.Join(wd, "rpggen.exe")
	cmd := exec.Command(binary, args...)
	cmd.Dir = wd
	var output bytes.Buffer
	cmd.Stdout = &output
	var errOutput bytes.Buffer
	cmd.Stderr = &errOutput

	if err := cmd.Run(); err != nil {
		errStr := errOutput.String()
		return "", fmt.Errorf("error when running rpggen: %s\noutput: %s", err, errStr)
	}

	return output.String(), nil
}

func build(wd string) error {
	make := exec.Command("make", "build")
	make.Dir = wd
	err := make.Run()
	if err != nil {
		fmt.Printf("unable to 'make build' for tests: %v\n", err)
	}
	return nil
}
