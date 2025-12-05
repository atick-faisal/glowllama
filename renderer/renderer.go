package renderer

import (
"io"
)

// OutputRenderer defines the interface for rendering model output
type OutputRenderer interface {
// Render processes the input string and writes the rendered output to the writer
Render(input string, w io.Writer) error

// ShouldRender returns whether rendering is enabled
ShouldRender() bool

// SupportsStreaming returns whether the renderer can handle streaming output
SupportsStreaming() bool
}

// Config holds configuration for the renderer
type Config struct {
Style          string // Renderer style (dark, light, auto)
EnableColor    bool   // Whether to enable color output
EnableRendering bool  // Whether rendering is enabled
BufferStream   bool   // Whether to buffer streaming output
Width          int    // Terminal width
}

// DefaultConfig returns a default renderer configuration
func DefaultConfig() Config {
return Config{
Style:          "dark",
EnableColor:    true,
EnableRendering: true,
BufferStream:   true,
Width:          80,
}
}
