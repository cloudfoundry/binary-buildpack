package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

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
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "fake_supply_binary_app"))
			app.Buildpacks = []string{
				"https://buildpacks.cloudfoundry.org/fixtures/new_supply_dotnet.zip",
				"binary_buildpack",
			}
			app.Disk = "1G"
		})

		It("finds the supplied dependency in the runtime container", func() {
			PushAppAndConfirm(app)

			Expect(app.Stdout.String()).To(ContainSubstring("SUPPLYING DOTNET"))
			Expect(app.GetBody("/")).To(ContainSubstring("dotnet: 1.0.1"))
		})
	})
})
