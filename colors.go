package macchiato

import "github.com/fatih/color"

type cprintf func(format string, args ...interface{}) string

var (

	// Render with green for success
	sbf = color.New(color.Bold, color.FgGreen).SprintfFunc() // bold
	//sf  = color.New(color.FgGreen).SprintfFunc()             // normal

	// Render for verbosity
	bf  = color.New(color.Bold).SprintfFunc()                  // bold
	gf  = color.New(color.Faint).SprintfFunc()                 // light
	buf = color.New(color.Bold, color.Underline).SprintfFunc() // bold + underline

	// Render with a yellow/orange for warning
	wf = color.New(color.FgYellow).SprintfFunc() // normal

	// Render with blue for note
	bbf = color.New(color.Bold, color.FgHiBlue).SprintfFunc() // bold

	// Render with cyan for info
	cbf = color.New(color.Bold, color.FgHiCyan).SprintfFunc() // bold
	cf  = color.New(color.FgHiCyan).SprintfFunc()             // normal

	// Render with red for error
	rbf = color.New(color.Bold, color.FgRed).SprintfFunc() // bold
	rf  = color.New(color.FgRed).SprintfFunc()             // normal

)
