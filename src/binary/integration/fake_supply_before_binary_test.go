package integration_test

import (
	"os"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"
	"github.com/cloudfoundry/libbuildpack/packager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("running supply buildpacks before the binary buildpack", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Context("the app is pushed once", func() {
		BeforeEach(func() {
			if version, err := cutlass.ApiVersion(); err != nil || version == "2.65.0" {
				Skip("API version does not have multi-buildpack support")
			}

			var err error
			var file string
			var extensionBuildpackDir = filepath.Join(bpDir, "fixtures", "fake_supply_buildpack")
			var version = "0.0.0"
			var cached = true
			var stack = os.Getenv("CF_STACK")
			var buildpackName = "fake_supply"

			file, err = packager.Package(extensionBuildpackDir, packager.CacheDir, version, stack, cached)
			Expect(err).NotTo(HaveOccurred())

			err = cutlass.CreateOrUpdateBuildpack(buildpackName, file, stack)
			Expect(err).NotTo(HaveOccurred())

			app = cutlass.New(filepath.Join(bpDir, "fixtures", "fake_supply_app"))
			app.Buildpacks = []string{
				"fake_supply_buildpack",
				"binary_buildpack",
			}
			app.Disk = "1G"
			app.Memory = "1G"
		})

		It("finds the supplied dependency in the runtime container", func() {
			PushAppAndConfirm(app)

			Expect(app.Stdout.String()).To(ContainSubstring("Running Fake Supply Buildpack"))
			Expect(app.GetBody("/")).To(MatchRegexp(`Dummy App running on localhost:\d+`))
		})
	})
})
