package brats_test

import (
	"github.com/cloudfoundry/libbuildpack/bratshelper"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Binary buildpack", func() {
	bratshelper.UnbuiltBuildpack("", CopyBrats)
	bratshelper.DeployingAnAppWithAnUpdatedVersionOfTheSameBuildpack(CopyBrats)
	bratshelper.DeployAppWithExecutableProfileScript("", CopyBrats)
	bratshelper.DeployAnAppWithSensitiveEnvironmentVariables(CopyBrats)
})
