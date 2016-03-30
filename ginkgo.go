package macchiato

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
)

// RunSpecs wrap ginkgo's RunSpecsWithCustomReporters by injecting its own Reporter.
func RunSpecs(t ginkgo.GinkgoTestingT, description string) bool {
	reporters := []ginkgo.Reporter{NewReporter()}
	return ginkgo.RunSpecsWithCustomReporters(t, description, reporters)
}

// NewReporter return a Macchiato reporter for Ginkgo.
func NewReporter() reporters.Reporter {
	stenographer := NewStenographer(config.DefaultReporterConfig.NoColor)
	return reporters.NewDefaultReporter(config.DefaultReporterConfig, stenographer)
}
