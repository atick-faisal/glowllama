# Glowllama Documentation

Glowllama is a fork of Ollama with enhanced terminal rendering via Glamour.

### Getting Started
* [Examples](./examples.md)
* [Development guide](./development.md)

### Differences from Ollama

* **Beautiful Markdown Rendering**: All model output is rendered with Glamour
* **Enhanced CLI**: Additional flags for controlling output rendering
* **Same API**: Fully compatible with Ollama's API and model library

### Ollama Documentation

For general Ollama documentation (models, API, setup), refer to:
* [Ollama Quickstart](https://docs.ollama.com/quickstart)
* [Importing models](https://docs.ollama.com/import)
* [API Reference](https://docs.ollama.com/api)
* [Modelfile Reference](https://docs.ollama.com/modelfile)
* [OpenAI Compatibility](https://docs.ollama.com/api/openai-compatibility)
* [Troubleshooting Guide](https://docs.ollama.com/troubleshooting)

### Glowllama-Specific Features

**Rendering Flags**
```bash
glowllama run llama3.3 --raw           # Disable rendering
glowllama run llama3.3 --renderer=dark # Use dark theme
glowllama run llama3.3 --no-style      # Disable colors
```

**Configuration**

Create `~/.config/glowllama/config.yaml`:
```yaml
renderer: dark
streaming_mode: buffered
color: auto
```
