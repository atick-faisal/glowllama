## Product Requirements Document (PRD): Glowllama

A fork of Ollama with integrated terminal-rendered Markdown output via Glamour.

---

## 1. Overview

Glowllama is a user-facing CLI application derived from Ollama that enhances the terminal experience by rendering model output using Glamour. It preserves full compatibility with the Ollama ecosystem (models, server protocol, config files, system architecture) while providing improved readability and developer ergonomics.

The project also includes a full rebrand of the CLI, repository, binaries, build artifacts, and documentation from Ollama to Glowllama.

---

## 2. Goals and Non-Goals

### Goals

1. Fully rendered Markdown output in the terminal using Glamour as the renderer.
2. A drop-in CLI compatible with existing Ollama commands unless explicitly deprecated.
3. Complete rebranding (binary names, commands, flags, build outputs).
4. Updated build pipeline (Makefile, Go build configs, Dockerfiles, GitHub Actions, release artifacts, versioning).
5. Clean internal refactor around the output pipeline to formalize the rendering interface.

### Non-Goals

1. Changing the server API or model execution behavior.
2. Modifying the model download, quantization, or storage formats.
3. Changing licensing terms beyond complying with MIT.
4. Introducing UI/UX features beyond terminal-based Markdown rendering.

---

## 3. Functional Requirements

### 3.1 Rendering Pipeline

* The CLI must render model output using Glamour with a default style (e.g. Dark or Auto).
* Streaming mode must remain functional:

  * Option A: buffer until completion and render once (preferred).
  * Option B: incremental rendering (more complex; optional v2).
* Flags to control rendering:

  * `--raw` to output un-rendered text (compat mode).
  * `--renderer=<style>` to choose Glamour styles.
  * `--no-style` to disable styling but preserve formatting.
* Output should retain color and formatting using ANSI codes.

### 3.2 Compatibility

* Existing Ollama commands should map:

  * `ollama run` -> `glowllama run`
  * `ollama pull` -> `glowllama pull`
  * `ollama list` -> `glowllama list`

### 3.3 Configuration

* Global config file:

  * `~/.config/glowllama/config.yaml`
* Config options:

  * `renderer: glamour-dark`
  * `streaming_mode: buffered | passthrough`
  * `color: auto | force | off`

### 3.4 User-Facing Branding

* Replace executable name: `ollama` -> `glowllama`
* Update all CLI help menus, examples, and printed text.
* Update README, docs, and installation instructions.
* Update logos and ASCII art (if present).
* Update repository metadata, tags, and release versioning.

---

## 4. Technical Requirements

### 4.1 Codebase Refactor

* Introduce an OutputRenderer interface:

  ```go
  type OutputRenderer interface {
      Render(markdown string) (string, error)
      ShouldRender() bool
  }
  ```
* Implement GlamourRenderer using Glamour.
* Create a RawRenderer for compatibility mode.
* Replace all direct writes to stdout with renderer-managed writes.

### 4.2 Integration

* Inject renderer into the CLI command execution pipeline.
* Ensure Docker integration and remote execution still work.
* Ensure Windows terminal compatibility (ANSI support detection).
* Add build tags or toggles for environments where Glamour isn't available.

### 4.3 CI/CD Pipeline Updates

* Update build workflows to output `glowllama` binaries for:

  * Linux x86_64, ARM64
  * macOS x86_64, ARM64
  * Windows x86_64
* Update Dockerfile to use new binary and labels.
* Ensure backward compatibility environment variables continue functioning unless renamed.

### 4.4 Licensing

* Verify MIT license compatibility.
* Retain credit to Ollama per license guidance.
* Document new license block for Glowllama fork.

---

## 5. User Experience

### 5.1 CLI Interaction

Example output (rendered in terminal):

````markdown
## Summary
This function sorts a list using quicksort.

```python
def quicksort(items):
    if len(items) <= 1:
        return items
    pivot = items[0]
    left = [x for x in items[1:] if x < pivot]
    right = [x for x in items[1:] if x >= pivot]
    return quicksort(left) + [pivot] + quicksort(right)
```
````

### 5.2 Failure Modes

* If rendering fails: fallback to raw output.
* If terminal does not support ANSI: auto-disable styling.
* If streaming + rendering is incompatible: gracefully disable streaming for that run.

---

## 6. Risks and Mitigations

| Risk                                       | Mitigation                                        |
| ------------------------------------------ | ------------------------------------------------- |
| Glamour is slow for large outputs          | Provide `--raw` and `--no-style` options          |
| Streaming rendering creates artifacts      | Default to buffered mode                          |
| Breaking compatibility with Ollama scripts | Provide a symlink `ollama` in the repo optionally |
| Terminal incompatibility                   | Auto-detect color support                         |

---

## 7. Deliverables

* Fully functional Glowllama CLI with rendered Markdown.
* Updated docs, README, install scripts, packaging.
* Version 1.0.0 release binaries and Docker image.
* A migration guide for Ollama users.

---

## End-to-End Implementation Checklist

### 1. Repo Setup and Branding

[x] Fork Ollama repo
[x] Rename binary from `ollama` to `glowllama`
[x] Update all go module paths
[x] Global search and replace all branding strings
[x] Update README, docs, examples
[ ] Update GitHub metadata (repo name, description, topics)

### 2. Code Architecture

[x] Create OutputRenderer interface
[x] Implement GlamourRenderer
[x] Implement RawRenderer
[x] Add renderer selection logic
[ ] Refactor CLI output to central rendering pipeline (IN PROGRESS)
[x] Add flags: `--raw`, `--renderer=`, `--no-style`
[ ] Implement config loader for renderer settings
[x] Add auto-detection for ANSI support

### 3. Markdown Rendering Integration

[x] Add Glamour dependency
[x] Add default style presets
[x] Implement safe fallback for rendering errors
[ ] Handle buffered output for streaming mode (IN PROGRESS)
[ ] Add unit tests for renderer behavior
[ ] Add snapshot tests for formatting

### 4. CLI Commands Update

[ ] Update help menus
[ ] Update command docs
[ ] Update default examples
[ ] Ensure all subcommands use renderer

### 5. Build and Packaging

[ ] Update Makefile with new binary name
[ ] Update Go build script
[ ] Update Dockerfile to use glowllama
[ ] Update install scripts (bash, zsh, fish completions)
[ ] Build release artifacts for all platforms
[ ] Smoke test binaries on Linux, macOS, Windows

### 6. CI/CD Updates

[ ] Update GitHub Actions workflows
[ ] Add build matrix for all OS/architectures
[ ] Add tests for renderer
[ ] Add automated releases with tagged versions
[ ] Update versioning strategy

### 7. Licensing

[ ] Keep original Ollama MIT license
[ ] Add attribution section
[ ] Add Glowllama MIT license
[ ] Document license changes in README

### 8. QA and Validation

[ ] Validate model loading and execution
[ ] Validate streaming behavior
[ ] Validate ANSI rendering across terminals
[ ] Validate performance on large outputs
[ ] Validate GPU, CPU builds
[ ] Validate Docker execution

### 9. Release

[ ] Tag v1.0.0
[ ] Publish binaries
[ ] Publish Docker image
[ ] Publish documentation site
[ ] Announce release notes
