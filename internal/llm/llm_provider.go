package llm

import (
	"context"
)

type Provider interface {
	GenerateDocumentation(ctx context.Context, prompt string) (string, error)
}
