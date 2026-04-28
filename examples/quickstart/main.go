// Quick smoke test for the otari Go SDK.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mozilla-ai/otari-sdk-go/otari"
)

func main() {
	apiBase := os.Getenv("OTARI_API_BASE")
	if apiBase == "" {
		apiBase = "http://localhost:8100"
	}
	token := os.Getenv("OTARI_PLATFORM_TOKEN")
	if token == "" {
		log.Fatal("OTARI_PLATFORM_TOKEN is required")
	}

	client, err := otari.New(
		otari.WithBaseURL(apiBase),
		otari.WithAPIKey(token),
		otari.WithPlatformMode(),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// -- Chat completion --
	fmt.Println("=== Chat Completion ===")
	resp, err := client.Completion(ctx, otari.CompletionParams{
		Model:    "openai:gpt-4o-mini",
		Messages: []otari.Message{{Role: otari.RoleUser, Content: "Say hello in three languages, one per line."}},
	})
	if err != nil {
		log.Fatalf("Completion failed: %v", err)
	}
	fmt.Println(resp.Choices[0].Message.ContentString())
	fmt.Println()

	// -- Streaming --
	fmt.Println("=== Streaming ===")
	chunks, errs := client.CompletionStream(ctx, otari.CompletionParams{
		Model:    "openai:gpt-4o-mini",
		Messages: []otari.Message{{Role: otari.RoleUser, Content: "Count from 1 to 5."}},
	})
	for chunk := range chunks {
		if len(chunk.Choices) > 0 {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
	if streamErr := <-errs; streamErr != nil {
		log.Fatalf("Streaming error: %v", streamErr)
	}
	fmt.Println()
	fmt.Println()

	// -- Embeddings --
	fmt.Println("=== Embeddings ===")
	emb, err := client.Embedding(ctx, otari.EmbeddingParams{
		Model: "openai:text-embedding-3-small",
		Input: "The quick brown fox jumps over the lazy dog.",
	})
	if err != nil {
		fmt.Printf("Skipped (%v)\n", err)
	} else {
		vec := emb.Data[0].Embedding
		fmt.Printf("Dimension: %d, first 5 values: %v\n", len(vec), vec[:5])
	}
	fmt.Println()

	// -- List models --
	fmt.Println("=== Models ===")
	models, err := client.ListModels(ctx)
	if err != nil {
		fmt.Printf("Skipped (%v)\n", err)
	} else {
		limit := 10
		if len(models.Data) < limit {
			limit = len(models.Data)
		}
		for _, m := range models.Data[:limit] {
			fmt.Printf("  %s\n", m.ID)
		}
		if len(models.Data) > 10 {
			fmt.Printf("  ... and %d more\n", len(models.Data)-10)
		}
	}

	// -- Rerank --
	fmt.Println("\n=== Rerank ===")
	rerank, err := client.Rerank(ctx, otari.RerankParams{
		Model:     "cohere:rerank-v3.5",
		Query:     "What is machine learning?",
		Documents: []string{"ML is a subset of AI", "The weather is nice", "Neural networks learn from data"},
	})
	if err != nil {
		fmt.Printf("Skipped (%v)\n", err)
	} else {
		for _, r := range rerank.Results {
			fmt.Printf("  doc[%d] score=%.3f\n", r.Index, r.RelevanceScore)
		}
	}

	fmt.Println()
	fmt.Println("Done.")
}
