package cmd

import (
"bytes"
"io"
"os"

"github.com/glowllama/glowllama/renderer"
"golang.org/x/term"
)

var globalRenderer renderer.OutputRenderer

// initRenderer initializes the global renderer based on options
func initRenderer(opts runOptions) error {
config := renderer.DefaultConfig()

// Check if output is a terminal
isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

// Disable rendering if not a terminal or raw mode is requested
if !isTerminal || opts.RawOutput {
globalRenderer = renderer.NewRawRenderer(config)
return nil
}

// Configure renderer based on flags
config.EnableRendering = !opts.RawOutput
config.EnableColor = !opts.NoStyle
config.Style = opts.RendererStyle

// Get terminal width
if width, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
config.Width = width
}

// Create Glamour renderer
r, err := renderer.NewGlamourRenderer(config)
if err != nil {
// Fallback to raw renderer on error
globalRenderer = renderer.NewRawRenderer(config)
return err
}

globalRenderer = r
return nil
}

// renderOutput renders markdown output if rendering is enabled
func renderOutput(content string) string {
if globalRenderer == nil || !globalRenderer.ShouldRender() {
return content
}

var buf bytes.Buffer
err := globalRenderer.Render(content, &buf)
if err != nil {
// Fallback to raw output on error
return content
}

return buf.String()
}

// writeRenderedOutput writes rendered output to the writer
func writeRenderedOutput(content string, w io.Writer) error {
if globalRenderer == nil {
_, err := w.Write([]byte(content))
return err
}

return globalRenderer.Render(content, w)
}
