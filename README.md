# DocuForge

> ⚡️ Instantly generate clean, intelligent documentation and diagrams for your infrastructure-as-code with the power of LLMs.

![DocuForge Banner](https://raw.githubusercontent.com/kishanrao92/docuforge/main/assets/banner.png)

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.20+-brightgreen)](https://golang.org)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](CONTRIBUTING.md)

---

## What is DocuForge?

**DocuForge** is an open-source CLI tool for **automated infrastructure documentation and interactive Q&A**. It uses LLMs (local or cloud) to generate Markdown documentation and explain Terraform, Kubernetes YAML, and other config files. You can even ask it questions about your infra, directly from your terminal.

---

## ✨ Features

- 📄 **Automatic Markdown docs** for Kubernetes, Terraform, Docker, and more
- 🤖 **Interactive chat mode** to query your infra like an expert
- 🔐 **Local LLM support** via [Ollama](https://ollama.com/) for privacy/offline use
- ☁️ **OpenAI integration** for enhanced language output (optional)
- ⚙️ **Smart regeneration** — skip unchanged files
- 🧩 **Extensible** for other IaC formats and models

---

## 🚀 Quick Start

### 1. Clone the Repo

```bash
git clone https://github.com/kishanrao92/docuforge.git
cd docuforge
```

### 2. Install Requirements

- Go 1.20+
- (Optional) [Ollama](https://ollama.com/) for local LLMs

### 3. Run DocuForge

```bash
go run main.go --local --model mistral
```

- Use `--local` to use a local LLM via Ollama
- Use `--model` to specify model (e.g., `mistral`, `tinyllama`, `phi3:mini`)

### 4. Ask Questions

```bash
Welcome to DocuForge. Ask me questions about your infrastructure in this repo!
You: What does the deployment.yml file do?
Agent: It defines a Kubernetes Deployment named "dummy" with 3 replicas of the "dummy-container" running image "dummy:0.1" on port 8080.
```

---

## 🔍 Why DocuForge?

Infrastructure-as-code grows fast. Docs don’t.  
**DocuForge** helps you:

- Save hours of manual documentation time
- Accelerate onboarding and audits
- Understand infra better through Q&A
- Avoid vendor lock-in with local LLM support
- Stay compliant with up-to-date infra docs

---

## 📦 Example

Given this file:

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

DocuForge generates:

> **deployment.yml**:  
> This file defines a Kubernetes Deployment named "dummy" with 3 replicas, running the "dummy-container" (image: dummy:0.1) on port 8080.

---

## 🖼 Demo

![DocuForge Demo](https://raw.githubusercontent.com/kishanrao92/docuforge/main/assets/demo.gif)

---

## ⚙️ Configuration

- Place your config files (YAML, HCL, TF, etc.) in a folder like `test-configs/`
- Output saved to `outputs/generated_doc.md`
- Set model/provider with CLI flags

---

## 📚 Supported Providers

| Type    | Provider     | Notes                               |
|---------|--------------|--------------------------------------|
| Local   | Ollama       | Mistral, TinyLlama, Phi-3, etc.      |
| Cloud   | OpenAI       | GPT-3.5/4, set `OPENAI_API_KEY` env  |

---

## 🛠 Roadmap

- [x] Markdown generation
- [x] Interactive Q&A mode
- [x] Ollama + OpenAI support
- [ ] MermaidJS diagram generation
- [ ] GitHub Action integration
- [ ] CDK & Pulumi support

---

## 💡 Ideal For

- DevOps and SRE teams
- Infra/platform engineers
- Startups needing audit-ready docs
- New hires onboarding onto your codebase

---

## 👏 Contributing

Contributions welcome!  
Open issues for features, file type support, or bugs. Star the repo to support the project.

---

## 📄 License

MIT License

---

## 🙏 Acknowledgements

- [Ollama](https://ollama.com/)
- [OpenAI](https://openai.com/)
- [briandowns/spinner](https://github.com/briandowns/spinner)
