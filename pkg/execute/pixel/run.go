// Package pixel opens a window for days that visualise a twod.Map. It is kept
// apart from pkg/execute so that other days do not link OpenGL.
package pixel

import (
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

// Run behaves like execute.Run, rendering maps to a window unless CI is set.
func Run(run func(string) (any, any), tests execute.TestCases, puzzle string, verbose bool) {
	if os.Getenv("CI") == "true" {
		execute.Run(run, tests, puzzle, verbose)
		return
	}
	twod.RenderingEnabled = true
	twod.Renderer = render
	pixelgl.Run(func() {
		execute.Run(run, tests, puzzle, verbose)
	})
}
