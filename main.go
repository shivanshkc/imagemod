package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const (
	// giveThemPrompt defines the AI prompt for hairstyle generation
	giveThemPrompt = "Generate a photorealistic image of the same person from the reference photo, but give them a hairstyle that suits them the most. It is essential to preserve their exact facial identity, ensuring they remain fully recognizable."
)

func main() {
	ctx := context.Background()

	// Parse command line arguments
	inputFile := flag.String("input", "./input.jpeg", "Path to input image file")
	outputFile := flag.String("output", "./output.png", "Path to output image file")
	flag.Parse()

	// Validate command line arguments
	if *inputFile == "" {
		panic("input file path cannot be empty")
	}
	if *outputFile == "" {
		panic("output file path cannot be empty")
	}

	// Get API key from environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY environment variable is required")
	}

	// Read the input image file
	imageData, err := os.ReadFile(*inputFile)
	if err != nil {
		panic(fmt.Errorf(`failed to read input image file "%s": %w`, *inputFile, err))
	}

	// Initialize Gemini AI client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: apiKey})
	if err != nil {
		panic("failed to create Gemini AI client: " + err.Error())
	}

	// Prepare the request with prompt and image data
	parts := []*genai.Part{
		{Text: giveThemPrompt},
		{InlineData: &genai.Blob{Data: imageData, MIMEType: "image/jpeg"}},
	}

	// Generate content using Gemini AI
	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash-image-preview", []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		panic("failed to generate content with Gemini AI: " + err.Error())
	}

	// Validate that we received candidates from the AI
	if len(result.Candidates) == 0 {
		panic("no candidates returned from Gemini AI - the request may have been filtered or failed")
	}

	// Process the response parts (text and image data)
	for _, part := range result.Candidates[0].Content.Parts {
		// Handle text responses (descriptions, etc.)
		if part.Text != "" {
			fmt.Println(part.Text)
			continue
		}

		// Handle image data responses
		if part.InlineData == nil {
			continue
		}

		// Save the generated image to output file
		imageBytes := part.InlineData.Data
		if err := os.WriteFile(*outputFile, imageBytes, 0644); err != nil {
			panic(fmt.Errorf(`failed to write output image to "%s": %w`, *outputFile, err))
		}
	}
}
