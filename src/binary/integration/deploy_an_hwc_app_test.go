package integration_test

import (
	"os"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF Binary Buildpack", func() {
	BeforeEach(SkipIfNotWindows)

	var app *cutlass.App

	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Describe("deploying a Windows HWC app", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "hwc_app"))
			app.Buildpacks = []string{cutlass.BuildpackNameForTest("binary", app.Stack)}
			app.Memory = "512M"
			app.Stack = os.Getenv("CF_STACK")
		})

		Context("without a command or Procfile", func() {
			BeforeEach(func() {
				app.StartCommand = "null"
			})

			It("logs a warning message", func() {
				Expect(app.Push()).ToNot(Succeed())
				Expect(app.ConfirmBuildpack(buildpackVersion)).To(Succeed())

				Eventually(app.Stdout.String).Should(ContainSubstring("Warning: We detected a Web.config in your app. This probably means that you want to use the hwc-buildpack. If you really want to use the binary-buildpack, you must specify a start command."))
			})
		})
	})
})
