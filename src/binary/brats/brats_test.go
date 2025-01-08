package brats_test

import (
	"github.com/cloudfoundry/libbuildpack/bratshelper"
	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Binary buildpack", func() {
	var copyBrats (func(string) *cutlass.App)
	if CanRunForOneOfStacks("cflinuxfs3", "cflinuxfs4") {
		copyBrats = CopyBrats
		bratshelper.UnbuiltBuildpack("", copyBrats)
		bratshelper.DeployAppWithExecutableProfileScript("", copyBrats)
	} else {
		copyBrats = CopyBratsWindows
	}

})
