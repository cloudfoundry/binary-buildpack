package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF Binary Buildpack", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Describe("deploying an app with no start command", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "no_start_command"))
			app.Buildpacks = []string{"binary_buildpack"}
			app.StartCommand = "null"
			app.Memory = "512M"
		})

		Context("on a windows stack", func() {
			BeforeEach(func() {
				SkipIfNoWindowsStack()
				app.Stack = "windows2012R2"
				app.StartCommand = "null"
			})

			It("logs a warning message", func() {
				Expect(app.Push()).ToNot(Succeed())
				Expect(app.ConfirmBuildpack(buildpackVersion)).To(Succeed())
				Eventually(app.Stdout.String).Should(ContainSubstring("Error: no start command specified during staging or launch"))
			})
		})

		Context("on a linux stack", func() {
			BeforeEach(func() {
				app.Stack = "cflinuxfs2"
				app.StartCommand = "null"
			})

			It("logs a warning message", func() {
				Expect(app.Push()).ToNot(Succeed())
				Expect(app.ConfirmBuildpack(buildpackVersion)).To(Succeed())
				Eventually(app.Stdout.String).Should(ContainSubstring("Error: no start command specified during staging or launch"))
			})
		})
	})
})
