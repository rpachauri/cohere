package main

import (
	"context"
	_ "embed"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

var (
	//go:embed chtoken
	cohereAPIKey string
)

func main() {
	client := cohereclient.NewClient(cohereclient.WithToken(cohereAPIKey))
	response, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: "How are you today?",
		},
	)
	if err != nil {
		panic(err)
	}
	println(response.String())
}
