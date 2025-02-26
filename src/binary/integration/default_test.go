package integration_test

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/switchblade"
	"github.com/sclevine/spec"

	. "github.com/cloudfoundry/switchblade/matchers"
	. "github.com/onsi/gomega"
)

func testDefault(platform switchblade.Platform, fixtures string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect     = NewWithT(t).Expect
			Eventually = NewWithT(t).Eventually

			name string
		)

		it.Before(func() {
			var err error
			name, err = switchblade.RandomName()
			Expect(err).NotTo(HaveOccurred())
		})

		it.After(func() {
			Expect(platform.Delete.Execute(name)).To(Succeed())
		})

		it("builds and runs the app when buildpack is specified", func() {
			deployment, _, err := platform.Deploy.
				WithBuildpacks("binary_buildpack").
				Execute(name, filepath.Join(fixtures, "default"))
			Expect(err).NotTo(HaveOccurred())

			Eventually(deployment).Should(Serve(ContainSubstring("Hello, world!")))
		})

		context("the buildpack is not specified", func() {
			it("it fails to run", func() {
				_, logs, err := platform.Deploy.
					Execute(name, filepath.Join(fixtures, "default"))
				Expect(err).To(HaveOccurred())

				Expect(logs).To(ContainSubstring("None of the buildpacks detected a compatible application"))
			})
		})

		context("there is no start command given", func() {
			it("it fails to run", func() {
				deployment, _, err := platform.Deploy.
					WithBuildpacks("binary_buildpack").
					Execute(name, filepath.Join(fixtures, "no_start_command"))
				Expect(err).NotTo(HaveOccurred())

				cmd := exec.Command("docker", "container", "logs", deployment.Name)

				output, err := cmd.CombinedOutput()
				Expect(err).NotTo(HaveOccurred())

				Expect(output).To(ContainSubstring("Error: no start command specified during staging or launch"), string(output))
			})
		})
	}
}
