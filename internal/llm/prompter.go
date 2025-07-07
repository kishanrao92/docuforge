package llm

import (
	"fmt"
	"log"
	"os"

	"github.com/kishanrao/docuforge/parser"
)

// func BuildPrompt(documents []parser.InfraDocument) string {
// 	prompt := `You are a senior DevOps documentation generator.
// For each config file provided, analyze its purpose, extract key configurations, and suggest an architecture diagram in MermaidJS format.

// `

// 	for _, doc := range documents {
// 		contents, err := os.ReadFile(doc.Path)
// 		if err != nil {
// 			continue // or better: log warning
// 		}

// 		prompt += fmt.Sprintf(`
// File: %s (%s)

// ---
// `, doc.Path, doc.Type, string(contents))
// 	}

// 	prompt += `
// Please generate a detailed Markdown documentation for the above infrastructure configuration.
// Include:
// - Purpose of each file.
// - Key configuration options explained.
// - Any dependencies.
// - An architecture diagram using MermaidJS (if applicable).
// `
// 	return prompt
// }

func BuildPrompt(documents []parser.InfraDocument) string {
	prompt := "Summarize the main resources and their key settings from these configuration files:\n\n"

	for _, doc := range documents {
		contents, err := os.ReadFile(doc.Path)
		if err != nil {
			log.Printf("warning: failed to read file %s: %v", doc.Path, err)
			continue
		}
		prompt += fmt.Sprintf("File: %s (%s)\n%s\n\n", doc.Path, doc.Type, string(contents))
	}

	prompt += "List the resources and their main settings in plain English."
	return prompt
}
