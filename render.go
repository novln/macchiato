package macchiato

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/onsi/ginkgo/types"
)

func getSpace(level int) string {
	return strings.Repeat(" ", (level+1)*2)
}

func renderTextWithoutSpace(text string) {
	renderText("", text)
}

func renderNewLine() {
	renderText("", "")
}

func renderText(space, text string) {
	fmt.Printf("%s%s\n", space, text)
}

func renderLine(space, icon, text string) {
	fmt.Printf("%s%s %s\n", space, icon, text)
}

func renderLineWithContext(space, icon, text, context string) {
	fmt.Printf("%s%s %s %s\n", space, icon, text, context)
}

func renderFinishedHeader(intro, status, runtime string) {
	fmt.Println()
	fmt.Printf("%s %s %s\n", intro, status, runtime)
	fmt.Println()
}

func renderFinishedStat(color cprintf, icon string, number int, status string) {
	if number > 0 {

		plural := "s"
		if number > 1 {
			plural = ""
		}

		l := color(" %s %d test%s %s", icon, number, plural, status)
		fmt.Println(l)
	}
}

func renderFinishedFooter() {
	renderNewLine()
}

func renderFailedSpecHeader() {
	renderNewLine()
	renderTextWithoutSpace(buf("Summary:"))
	renderNewLine()
}

func renderFailedSpecContext(space, text string, failure types.SpecFailure) {
	renderLine(space, rbf(Icons.failed), rf(text))
	renderText(space, gf(failure.Location.String()))
	renderNewLine()
	renderTextWithoutSpace(rf(failure.Message))
	if failure.ForwardedPanic != "" {
		renderTextWithoutSpace(rf(failure.ForwardedPanic))
	}
	renderNewLine()
}

func renderPanickedSpecContext(space, text string, failure types.SpecFailure) {
	renderLine(space, rbf(Icons.panicked), rf(text))
	renderNewLine()
	scanner := bufio.NewScanner(strings.NewReader(failure.Location.FullStackTrace))
	for scanner.Scan() {
		renderText(space, strings.Replace(scanner.Text(), "\t", "  ", 1))
	}
	renderNewLine()
	renderTextWithoutSpace(rf(failure.Message))
	if failure.ForwardedPanic != "" {
		renderTextWithoutSpace(rf(failure.ForwardedPanic))
	}
	renderNewLine()
}

func sortMeasurementKeys(measurements map[string]*types.SpecMeasurement) []string {
	orderedKeys := make([]string, len(measurements))
	for key, measurement := range measurements {
		orderedKeys[measurement.Order] = key
	}
	return orderedKeys
}

func renderMeasurementContext(space string, spec *types.SpecSummary) {

	if len(spec.Measurements) == 0 {
		return
	}

	for _, key := range sortMeasurementKeys(spec.Measurements) {

		m := spec.Measurements[key]
		r := fmt.Sprintf("  %s: %s%s < %s%s (± %s%s) < %s%s", m.Name,
			fmt.Sprintf("%.3f", m.Smallest), m.Units,
			fmt.Sprintf("%.3f", m.Average), m.Units,
			fmt.Sprintf("%.3f", m.StdDeviation), m.Units,
			fmt.Sprintf("%.3f", m.Largest), m.Units,
		)

		renderText(space, gf(r))
	}
}

func renderSuiteFailure(name string, summary *types.SetupSummary) {
	renderNewLine()
	renderTextWithoutSpace(buf(name))
	renderNewLine()
	renderFailedSpecContext(getSpace(0), fmt.Sprintf("An error has occurred with %s.", name), summary.Failure)
}
