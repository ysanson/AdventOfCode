package twod

import "image/color"

var (
	RenderingEnabled = false
	RenderingMap     = make(map[interface{}]color.Color)
	// Renderer draws a map. It is set by pkg/execute/pixel, which keeps the
	// OpenGL dependency out of days that never render.
	Renderer func(Map)
)

func (m Map) Render() {
	if RenderingEnabled && Renderer != nil {
		Renderer(m)
	}
}
