# DocuForge

DocuForge is a CLI tool for **automated infrastructure documentation and interactive Q&A** about your infrastructure-as-code (IaC) files. It supports local and cloud LLMs (like Ollama and OpenAI) to generate Markdown documentation and answer questions about your repo’s configuration files.

---

## Features

- **Automatic Markdown documentation** for Kubernetes, Terraform, and other config files.
- **Interactive chat agent**: Ask questions about your infrastructure from the CLI.
- **Supports local LLMs** (Ollama, Mistral, TinyLlama, etc.) for privacy and offline use.
- **Supports OpenAI** for richer answers (optional).
- **Smart regeneration**: Only regenerates docs if configs change or docs are outdated.
- **Extensible**: Easily add support for more file types or LLMs.

---

## Quick Start

### 1. Clone the Repo

```sh
git clone https://github.com/yourusername/docuforge.git
cd docuforge
```

### 2. Install Dependencies

- Go 1.20+
- (Optional) [Ollama](https://ollama.com/) for local LLMs

### 3. Run DocuForge

```sh
go run main.go --local --model mistral
```

- Use `--local` to use a local LLM (Ollama).
- Use `--model` to specify the model (e.g., `mistral`, `tinyllama`, `phi3:mini`).

### 4. Ask Questions

After documentation is generated, enter chat mode:

```
Welcome to Docuforge. Ask me questions about your infrastructure in this repo!
You: What does the deployment.yml file do?
Agent: It defines a Kubernetes Deployment named "dummy" with 3 replicas of the "dummy-container" running image "dummy:0.1" on port 8080.
```

---

## Example

Given this Kubernetes deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dummy
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: dummy-container
        image: dummy:0.1
        ports:
        - containerPort: 8080
```

DocuForge will generate:

> **deployment.yml**:  
> This file defines a Kubernetes Deployment named "dummy" with 3 replicas, running the "dummy-container" (image: dummy:0.1) on port 8080.

---

## Configuration

- Place your config files (YAML, HCL, TF, etc.) in a directory (e.g., `test-configs/`).
- Output documentation is saved to `outputs/generated_doc.md`.
- You can adjust the model and provider via CLI flags.

---

## Supported Providers

- **Local**: [Ollama](https://ollama.com/) (Mistral, TinyLlama, Phi-3, etc.)
- **Cloud**: OpenAI (set `OPENAI_API_KEY`)

---

## Contributing

Contributions are welcome! Please open issues or PRs for new features, bug fixes, or additional file type support.

---

## License

MIT License

---

## Acknowledgements

- [Ollama](https://ollama.com/)
- [OpenAI](https://openai.com/)
- [briandowns/spinner](https://github.com/briandowns/spinner)