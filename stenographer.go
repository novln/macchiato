package macchiato

import (
	"sync"

	"github.com/fatih/color"
	"github.com/onsi/ginkgo/types"
)

// Stenographer implements github.com/onsi/ginkgo/reporters/stenographer.Stenographer interface.
// It handle the rendering of the Reporter in the user's console.
type Stenographer struct {
	lock    *sync.Mutex
	indexes map[int]string
}

// NewStenographer return a new Stenographer.
func NewStenographer(nc bool) *Stenographer {
	color.NoColor = nc
	return &Stenographer{
		lock:    &sync.Mutex{},
		indexes: make(map[int]string),
	}
}

// AnnounceSuite print the suite's description
func (s *Stenographer) AnnounceSuite(description string, randomSeed int64, randomizingAll bool, quiet bool) {
	renderTextWithoutSpace(buf(description))
}

// AnnounceAggregatedParallelRun will do nothing.
func (s *Stenographer) AnnounceAggregatedParallelRun(nodes int, quiet bool) {
	// Ignore rendering.
}

// AnnounceParallelRun will do nothing.
func (s *Stenographer) AnnounceParallelRun(node int, nodes int, succinct bool) {
	// Ignore rendering.
}

// AnnounceTotalNumberOfSpecs will do nothing.
func (s *Stenographer) AnnounceTotalNumberOfSpecs(total int, succinct bool) {
	// Ignore rendering.
}

// AnnounceNumberOfSpecs will do nothing.
func (s *Stenographer) AnnounceNumberOfSpecs(specsToRun int, total int, quiet bool) {
	// Ignore rendering.
}

// AnnounceSpecRunCompletion will print a summary with some statistics of specs states.
func (s *Stenographer) AnnounceSpecRunCompletion(summary *types.SuiteSummary, quiet bool) {

	if quiet && summary.SuiteSucceeded {
		return
	}

	status := func() string {
		if summary.SuiteSucceeded {
			return sbf("success")
		}
		if summary.NumberOfFailedSpecs == 1 {
			return rbf("failure")
		}
		return rbf("failures")
	}

	renderFinishedHeader(bf("Finished with"), status(), bf("in %.3f secs", summary.RunTime.Seconds()))

	renderFinishedStat(sbf, Icons.passed, summary.NumberOfPassedSpecs, "passed")
	renderFinishedStat(rbf, Icons.failed, summary.NumberOfFailedSpecs, "failed")
	renderFinishedStat(bbf, Icons.skipped, summary.NumberOfSkippedSpecs, "skipped")
	renderFinishedStat(cbf, Icons.pending, summary.NumberOfPendingSpecs, "pending")

	renderFinishedFooter()

}

// AnnounceSpecWillRun will do nothing.
func (s *Stenographer) AnnounceSpecWillRun(spec *types.SpecSummary) {
	// Ignore rendering.
}

// AnnounceBeforeSuiteFailure will print the failure that hapenned in BeforeSuite hook.
func (s *Stenographer) AnnounceBeforeSuiteFailure(summary *types.SetupSummary, quiet bool, fullTrace bool) {
	renderSuiteFailure("BeforeSuite", summary)
}

// AnnounceAfterSuiteFailure will print the failure that hapenned in AfterSuite hook.
func (s *Stenographer) AnnounceAfterSuiteFailure(summary *types.SetupSummary, quiet bool, fullTrace bool) {
	renderSuiteFailure("AfterSuite", summary)
}

// AnnounceCapturedOutput is a dead interface function, so it will do nothing.
func (s *Stenographer) AnnounceCapturedOutput(output string) {
	// Ignore rendering.
}

// AnnounceSuccesfulSpec will print the stack of the spec with its status.
func (s *Stenographer) AnnounceSuccesfulSpec(spec *types.SpecSummary) {
	s.renderLines(spec, false, func(space, text string) {
		renderLine(space, sbf(Icons.passed), gf(text))
	})
}

// AnnounceSuccesfulSlowSpec will print the stack of the spec with its status.
func (s *Stenographer) AnnounceSuccesfulSlowSpec(spec *types.SpecSummary, quiet bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLineWithContext(space, sbf(Icons.passed), gf(text), wf("(%.3f secs)", spec.RunTime.Seconds()))
	})
}

// AnnounceSuccesfulMeasurement will print the stack of the benchmark with its measurements.
func (s *Stenographer) AnnounceSuccesfulMeasurement(spec *types.SpecSummary, quiet bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLine(space, sbf(Icons.passed), gf(text))
		renderMeasurementContext(space, spec)
	})
}

// AnnouncePendingSpec will print the stack of the spec with its status.
func (s *Stenographer) AnnouncePendingSpec(spec *types.SpecSummary, verbose bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLine(space, cbf(Icons.pending), cf(text))
	})
}

// AnnounceSkippedSpec will do nothing.
func (s *Stenographer) AnnounceSkippedSpec(spec *types.SpecSummary, quiet bool, fullTrace bool) {
	// Ignore rendering.
}

// AnnounceSpecTimedOut will print the stack of the spec with its status.
func (s *Stenographer) AnnounceSpecTimedOut(spec *types.SpecSummary, quiet bool, fullTrace bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLineWithContext(space, rbf(Icons.failed), rf(text), rbf("(timeout)"))
	})
}

// AnnounceSpecPanicked will print the stack of the spec with its status.
func (s *Stenographer) AnnounceSpecPanicked(spec *types.SpecSummary, quiet bool, fullTrace bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLineWithContext(space, rbf(Icons.panicked), rf(text), rbf("(panic)"))
	})
}

// AnnounceSpecFailed will print the stack of the spec with its status.
func (s *Stenographer) AnnounceSpecFailed(spec *types.SpecSummary, quiet bool, fullTrace bool) {
	s.renderLines(spec, false, func(space, text string) {
		renderLine(space, rbf(Icons.failed), rf(text))
	})
}

// SummarizeFailures will print a summary of all failed specs with their context.
func (s *Stenographer) SummarizeFailures(summaries []*types.SpecSummary) {

	s.resetIndexes()
	failingSpecs := []*types.SpecSummary{}

	for _, summary := range summaries {
		if summary.HasFailureState() {
			failingSpecs = append(failingSpecs, summary)
		}
	}

	if len(failingSpecs) == 0 {
		return
	}

	renderFailedSpecHeader()

	for _, summary := range failingSpecs {
		if summary.HasFailureState() {
			s.renderLines(summary, false, func(space, text string) {
				renderFailedSpecContext(space, text, summary.Failure)
			})
		}
	}
}

func (s *Stenographer) resetIndexes() {
	s.indexes = make(map[int]string)
}

func (s *Stenographer) hasIndex(index int, text string) bool {
	l, ok := s.indexes[index]
	return ok && l == text
}

func (s *Stenographer) setIndex(index int, text string) {

	// Update indexes with new entry.
	s.indexes[index] = text

	// Delete trailing entries, if any...
	offset := 1
	l := len(s.indexes)

	for i := index; i < l; i++ {
		delete(s.indexes, (i + offset))
	}
}

func (s *Stenographer) renderLines(spec *types.SpecSummary, verbose bool, render func(space, text string)) {

	s.lock.Lock()
	defer s.lock.Unlock()

	startIndex := 1
	lastIndex := len(spec.ComponentTexts)
	if lastIndex == 1 {
		startIndex = 0
	}

	length := lastIndex - 1

	for i := startIndex; i < lastIndex; i++ {

		text := spec.ComponentTexts[i]
		location := spec.ComponentCodeLocations[i].String()

		if !s.hasIndex(i, text) {

			space := getSpace(i - startIndex)
			s.setIndex(i, text)

			if i == length {
				render(space, text)
			} else if i == startIndex {
				renderText(space, bf(text))
			} else {
				renderText(space, text)
			}

			if verbose {
				renderText(space, gf(location))
			}
		}
	}
}
