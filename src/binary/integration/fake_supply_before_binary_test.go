package integration_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
			if !strings.HasPrefix(os.Getenv("CF_STACK"), "cflinuxfs") {
				Skip(fmt.Sprintf("Skipping because the current stack %s is not supported", os.Getenv("CF_STACK")))
			}

			if version, err := cutlass.ApiVersion(); err != nil || version == "2.65.0" {
				Skip("API version does not have multi-buildpack support")
			}

			app = cutlass.New(filepath.Join(bpDir, "fixtures", "fake_supply_binary_app"))
			app.Buildpacks = []string{
				"https://github.com/cloudfoundry/dotnet-core-buildpack#develop",
				"binary_buildpack",
			}
			app.Disk = "1G"
		})

		It("finds the supplied dependency in the runtime container", func() {
			PushAppAndConfirm(app)

			Expect(app.Stdout.String()).To(ContainSubstring("Supplying Dotnet Core"))
			Expect(app.GetBody("/")).To(MatchRegexp(`dotnet: \d+\.\d+\.\d+`))
		})
	})
})
