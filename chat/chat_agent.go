package chat

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/kishanrao/docuforge/internal/llm"
)

func ChatAgentWithDoc(docPath string, provider llm.Provider) {

	ctx := context.Background()
	log.Println("INSIDE AGENT CODE")
	generatedDocContent, err := os.ReadFile(docPath)

	if err != nil {
		log.Printf("Error reading generated document file , returning: %v\n ", err)
		return
	}

	basePrompt := string(generatedDocContent)

	log.Println("Welcome to Docuforge. Ask me questions about your infrastructure in this repo!")

	reader := bufio.NewReader(os.Stdin)

	for {
		log.Println("\nYou: ")
		userInput, _ := reader.ReadString('\n')

		if userInput == "exit" {
			log.Println("See ya!")
			break
		}
		chatPrompt := basePrompt + "\n\n User supplied questions about the generated docs: " + userInput

		response, err := provider.GenerateDocumentation(ctx, chatPrompt)
		if err != nil {
			log.Println("Error generating a response in chat")
			continue
		}

		log.Printf("Agent: %s\n", response)

	}

}
