package brats_test

import (
	"github.com/cloudfoundry/libbuildpack/bratshelper"
	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Binary buildpack", func() {
	var copyBrats (func(string) *cutlass.App)
	if CanRunForOneOfStacks("cflinuxfs2", "cflinuxfs3") {
		copyBrats = CopyBrats
		bratshelper.UnbuiltBuildpack("", copyBrats)
		bratshelper.DeployAppWithExecutableProfileScript("", copyBrats)
		bratshelper.DeployingAnAppWithAnUpdatedVersionOfTheSameBuildpack(copyBrats)
	} else {
		copyBrats = CopyBratsWindows
	}

	bratshelper.DeployAnAppWithSensitiveEnvironmentVariables(copyBrats)

})
