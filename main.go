/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/kishanrao/docuforge/chat"
	"github.com/kishanrao/docuforge/cmd"
	"github.com/kishanrao/docuforge/helpers"
	"github.com/kishanrao/docuforge/internal/llm"
	"github.com/kishanrao/docuforge/parser"
)

func main() {

	localMode := flag.Bool("local", true, "use local llm like mistral with ollama")
	dirPath := "/Users/kishan/Projects/docuforge/test-configs"
	const outputDocPath = "outputs/generated_doc.md"
	configFiles, err := parser.GetConfigFiles(dirPath)

	if err != nil {
		log.Fatal("No config files and some random errors")
	}

	log.Println("Found config files")

	for _, file := range configFiles {
		log.Printf("Path: %s, Type: %s\n", file.Path, file.Type)
	}

	infraDocs, err := parser.LoadInfraDocuments(configFiles)
	if err != nil {
		log.Fatalf("error loading infra documents: %v", err)
	}

	// 4. Now you have infraDocs available to print, send to LLM, etc
	// for _, doc := range infraDocs {
	// 	log.Println("========")
	// 	log.Printf("Path: %s\n", doc.Path)
	// 	log.Printf("Type: %s\n", doc.Type)
	// 	log.Printf("Content:\n%s\n", doc.Content)
	// 	log.Println("========\n")
	// }

	// Step 3: Build prompt
	prompt := llm.BuildPrompt(infraDocs)

	var provider llm.Provider

	if *localMode {
		provider = llm.NewLocalLLMProvider("mistral")
		log.Println("Using local llm with mistral")
	} else {
		provider = llm.NewOpenAiProvider()
		log.Print("Using open ai provider ")
	}

	if !helpers.IsDocRecent(outputDocPath, 24*time.Hour) {

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("red"))
		s.Suffix = "Generating knowledge base for your repo!"
		s.Start()

		// Step 4: Call LLM
		ctx := context.Background()
		output, err := provider.GenerateDocumentation(ctx, prompt)
		if err != nil {
			log.Fatalf("error generating documentation: %v", err)
		}

		// Step 4: Save output
		err = os.WriteFile("outputs/generated_doc.md", []byte(output), 0644)
		if err != nil {
			log.Fatalf("error saving documentation: %v", err)
		}

		s.Stop()
		log.Println("✅ Documentation written to outputs/generated_doc.md")
	} else {
		log.Println("Generated output file is less than 24 hours old, reusing for chat")
	}

	chat.ChatAgentWithDoc("/Users/kishan/Projects/docuforge/outputs/generated_doc.md", provider)

	// (Optional) invoke CLI framework if used
	cmd.Execute()
}
