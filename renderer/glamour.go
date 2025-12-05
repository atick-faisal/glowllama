package renderer

import (
"io"
"os"

"github.com/charmbracelet/glamour"
"golang.org/x/term"
)

// GlamourRenderer implements OutputRenderer using Glamour
type GlamourRenderer struct {
config Config
renderer *glamour.TermRenderer
}

// NewGlamourRenderer creates a new Glamour-based renderer
func NewGlamourRenderer(config Config) (*GlamourRenderer, error) {
// Determine terminal width
width := config.Width
if width <= 0 {
if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
width = w
} else {
width = 80
}
}

// Create glamour renderer with style
var r *glamour.TermRenderer
var err error

if config.Style != "" {
r, err = glamour.NewTermRenderer(
glamour.WithStylePath(config.Style),
glamour.WithWordWrap(width),
)
} else {
r, err = glamour.NewTermRenderer(
glamour.WithAutoStyle(),
glamour.WithWordWrap(width),
)
}

if err != nil {
return nil, err
}

return &GlamourRenderer{
config:   config,
renderer: r,
}, nil
}

// Render processes markdown input and writes rendered output
func (g *GlamourRenderer) Render(input string, w io.Writer) error {
if !g.config.EnableRendering {
_, err := w.Write([]byte(input))
return err
}

rendered, err := g.renderer.Render(input)
if err != nil {
// Fallback to raw output on error
_, writeErr := w.Write([]byte(input))
if writeErr != nil {
return writeErr
}
return err
}

_, err = w.Write([]byte(rendered))
return err
}

// ShouldRender returns whether rendering is enabled
func (g *GlamourRenderer) ShouldRender() bool {
return g.config.EnableRendering
}

// SupportsStreaming returns whether streaming is supported
func (g *GlamourRenderer) SupportsStreaming() bool {
// For now, we buffer by default for best rendering
return !g.config.BufferStream
}
