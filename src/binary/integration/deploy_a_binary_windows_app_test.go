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

	Describe("deploying a Windows batch script", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "windows_app"))
			app.Stack = os.Getenv("CF_STACK")
		})

		Context("when specifying a buildpack", func() {
			BeforeEach(func() {
				app.Buildpacks = []string{"binary_buildpack"}
			})

			It("deploys successfully", func() {
				PushAppAndConfirm(app)
				Eventually(app.Stdout.String).Should(ContainSubstring("Hello, world!"))
			})
		})

		Context("without specifying a buildpack", func() {
			BeforeEach(func() {
				app.Buildpacks = []string{}
			})

			It("fails to stage", func() {
				Expect(app.Push()).ToNot(Succeed())

				Eventually(app.Stdout.String).Should(ContainSubstring("None of the buildpacks detected a compatible application"))
			})
		})

		Context("without a command or Procfile", func() {
			BeforeEach(func() {
				app.Buildpacks = []string{"binary_buildpack"}
				app.Memory = "512M"
				app.StartCommand = "null"
			})

			It("logs an error message", func() {
				app.Push()
				Expect(app.ConfirmBuildpack(buildpackVersion)).To(Succeed())

				Eventually(app.Stdout.String).Should(ContainSubstring("Error: no start command specified during staging or launch"))
			})
		})
	})
})
