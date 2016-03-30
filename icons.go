package macchiato

import (
	"runtime"
)

var (
	// Icons is a default SpecsIcons
	Icons = newSpecsIcons()
)

// SpecsIcons contains unicode "icons" for specs status.
type SpecsIcons struct {
	passed  string
	failed  string
	pending string
	skipped string
}

func newSpecsIcons() SpecsIcons {

	if runtime.GOOS == "windows" {
		return SpecsIcons{
			passed:  " ",
			failed:  " ",
			pending: " ",
			skipped: " ",
		}
	}

	return SpecsIcons{
		passed:  `✔`,
		failed:  `✘`,
		pending: `❗`,
		skipped: `✱`,
	}
}
