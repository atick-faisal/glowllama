package renderer

import (
"io"
)

// RawRenderer implements OutputRenderer with no rendering (passthrough)
type RawRenderer struct {
config Config
}

// NewRawRenderer creates a new raw (passthrough) renderer
func NewRawRenderer(config Config) *RawRenderer {
return &RawRenderer{
config: config,
}
}

// Render passes through input without modification
func (r *RawRenderer) Render(input string, w io.Writer) error {
_, err := w.Write([]byte(input))
return err
}

// ShouldRender always returns false for raw renderer
func (r *RawRenderer) ShouldRender() bool {
return false
}

// SupportsStreaming always returns true for raw renderer
func (r *RawRenderer) SupportsStreaming() bool {
return true
}
