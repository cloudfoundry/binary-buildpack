package integration_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/switchblade"
	"github.com/sclevine/spec"

	. "github.com/cloudfoundry/switchblade/matchers"
	. "github.com/onsi/gomega"
)

func testFakeSupply(platform switchblade.Platform, fixtures string) func(*testing.T, spec.G, spec.S) {
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

		it("builds and runs the app when buildpack and fake supply are specified", func() {
			deployment, logs, err := platform.Deploy.
				WithBuildpacks(
					"fake_supply",
					"binary_buildpack",
				).
				Execute(name, filepath.Join(fixtures, "fake_supply"))
			Expect(err).NotTo(HaveOccurred())

			Expect(logs).To(ContainSubstring("Running Fake Supply Buildpack"))
			Eventually(deployment).Should(Serve(MatchRegexp(`Dummy App running on localhost:\d+`)))
		})
	}
}
