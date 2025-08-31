package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const (
	giveThemPrompt = "Generate a photorealistic image of the same person from the reference photo, but give them a hairstyle that suits them the most. It is essential to preserve their exact facial identity, ensuring they remain fully recognizable."

	inputFile  = "./input/input.jpeg"
	outputFile = "./output/output.png"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY environment variable is required")
	}

	imageData, err := os.ReadFile(inputFile)
	if err != nil {
		panic("error while reading image file: " + err.Error())
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		panic("error in genai.NewClient call: " + err.Error())
	}

	parts := []*genai.Part{
		{Text: giveThemPrompt},
		{InlineData: &genai.Blob{Data: imageData, MIMEType: "image/jpeg"}},
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash-image-preview", []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		panic("error in GenerateContent call: " + err.Error())
	}

	for _, part := range result.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
			continue
		}

		if part.InlineData == nil {
			continue
		}

		imageBytes := part.InlineData.Data
		if err := os.WriteFile(outputFile, imageBytes, 0644); err != nil {
			panic("error while writing image: " + err.Error())
		}
	}
}
