package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const (
	// defaultSamePersonPrompt defines the default AI prompt for hairstyle generation
	defaultSamePersonPrompt = "give them a hairstyle that suits them the most"

	// samePersonPrefix defines the prefix added to same-person prompts
	samePersonPrefix = "Generate a photorealistic image of the same person from the reference photo, but "

	// samePersonSuffix defines the suffix added to same-person prompts
	samePersonSuffix = ". It is essential to preserve their exact facial identity, ensuring they remain fully recognizable."
)

func main() {
	ctx := context.Background()

	// Input-output file flags.
	inputFile := flag.String("input", "./input.jpeg", "Path to input image file")
	outputFile := flag.String("output", "./output.png", "Path to output image file")

	// Prompt flags.
	samePersonPrompt := flag.String("same-person-prompt", defaultSamePersonPrompt, "Custom prompt for same-person image generation")
	fullPrompt := flag.String("full-prompt", "", "Full custom prompt (overrides -same-person-prompt)")

	// Load flags.
	flag.Parse()

	// Validate command line arguments
	if *inputFile == "" {
		panic("input file path cannot be empty")
	}
	if *outputFile == "" {
		panic("output file path cannot be empty")
	}

	// Determine the final prompt to use
	var finalPrompt string
	if *fullPrompt != "" {
		// Use full prompt as-is if provided
		finalPrompt = *fullPrompt
	} else {
		// Construct same-person prompt with prefix and suffix
		finalPrompt = samePersonPrefix + *samePersonPrompt + samePersonSuffix
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
		{Text: finalPrompt},
		{InlineData: &genai.Blob{Data: imageData, MIMEType: "image/jpeg"}},
	}

	fmt.Println("YOU:", finalPrompt)
	// Generate content using Gemini AI
	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash-image-preview", []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		panic("failed to generate content with Gemini AI: " + err.Error())
	}

	// Validate that we received candidates from the AI
	if len(result.Candidates) == 0 {
		panic("no candidates returned from Gemini AI - the request may have been filtered or failed")
	}

	fmt.Print("AGENT: ")
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
