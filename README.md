<div align="center">
  <h1>🦙✨ Glowllama</h1>
  <p><strong>Get up and running with large language models - beautifully rendered in your terminal.</strong></p>
  <p><em>A fork of <a href="https://github.com/ollama/ollama">Ollama</a> with integrated terminal-rendered Markdown output via <a href="https://github.com/charmbracelet/glamour">Glamour</a>.</em></p>
</div>

---

## ✨ What's Different?

Glowllama enhances the Ollama experience with:

- **🎨 Beautiful Markdown Rendering**: All model output is rendered with Glamour for improved readability
- **🌈 Syntax Highlighting**: Code blocks are automatically highlighted in your terminal
- **📝 Better Formatting**: Tables, lists, and headers are properly formatted
- **🎭 Multiple Themes**: Choose from various rendering styles
- **⚡ Performance**: Optional buffered streaming for optimal rendering

## 🚀 Installation

### macOS & Linux

```bash
# Coming soon
curl -fsSL https://glowllama.dev/install.sh | sh
```

### Build from Source

```bash
git clone https://github.com/glowllama/glowllama.git
cd glowllama
go build .
```

## 🎯 Quickstart

To run and chat with a model:

```bash
glowllama run gemma3
```

All standard Ollama commands work with Glowllama:

```bash
glowllama pull llama3.3     # Pull a model
glowllama list              # List installed models
glowllama run qwq           # Run a model
```

## 🎨 Rendering Options

Control how output is rendered:

```bash
# Use a specific rendering style
glowllama run gemma3 --renderer=dark

# Disable rendering (raw output)
glowllama run gemma3 --raw

# Disable styling but keep formatting
glowllama run gemma3 --no-style
```

## ⚙️ Configuration

Create `~/.config/glowllama/config.yaml`:

```yaml
renderer: glamour-dark
streaming_mode: buffered  # or passthrough
color: auto              # auto, force, or off
```

## 📚 Model Library

Glowllama is fully compatible with Ollama's model library. Visit [ollama.com/library](https://ollama.com/library) to browse available models.

Popular models:
- `glowllama run gemma3` - Google's Gemma 3
- `glowllama run llama3.3` - Meta's Llama 3.3
- `glowllama run qwq` - Alibaba's QwQ
- `glowllama run deepseek-r1` - DeepSeek R1

## 🔧 Development

```bash
# Run tests
go test ./...

# Build
go build .

# Install locally
go install .
```

## 🤝 Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📜 License

MIT License - see [LICENSE](LICENSE) for details.

This project is a fork of [Ollama](https://github.com/ollama/ollama) and maintains compatibility with the original project while adding enhanced terminal rendering capabilities.

## 🙏 Credits

- **Ollama Team** - For the amazing LLM runtime and CLI foundation
- **Charm** - For the beautiful Glamour Markdown renderer
- All contributors to both projects

---

<div align="center">
  <p>Made with 💜 by the Glowllama community</p>
</div>
