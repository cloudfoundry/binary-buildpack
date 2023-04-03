package integration_test

import (
	"path/filepath"

	"os"

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

	Describe("deploying a binary app", func() {
		BeforeEach(func() {
			SkipIfNotLinux()
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "default_app"))
			app.Stack = os.Getenv("CF_STACK")
		})

		Context("when specifying a buildpack", func() {
			BeforeEach(func() {
				app.Buildpacks = []string{"binary_buildpack"}
			})

			It("deploys successfully", func() {
				PushAppAndConfirm(app)

				Expect(app.GetBody("/")).To(ContainSubstring("Hello, world!"))
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
	})
})
