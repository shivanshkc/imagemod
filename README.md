# Image Generation Tool

## TL;DR

A command-line tool that uses Google's Gemini AI to generate new images based on an input photo. Perfect for trying different hairstyles while keeping the same face.

**Quick Start:**
1. Get a Gemini API key from Google AI Studio
2. Set environment variable: `export GEMINI_API_KEY=your_api_key_here`
3. Run: `go run main.go -input your_photo.jpg -output new_image.png`

## Overview

This tool leverages Google's Gemini AI to transform images while preserving facial identity. It's designed primarily for generating different hairstyles, but can be used for other image transformations with custom prompts.

## Prerequisites

- Go 1.23.5 or later
- Google Gemini API key (free tier available)

## Setup

### 1. Get a Gemini API Key

1. Visit [Google AI Studio](https://aistudio.google.com/)
2. Sign in with your Google account
3. Create a new API key
4. Copy the API key for use in step 2

### 2. Set Environment Variable

```bash
export GEMINI_API_KEY=your_api_key_here
```

For permanent setup, add this line to your shell profile (`.bashrc`, `.zshrc`, etc.).

### 3. Install Dependencies

```bash
go mod download
```

## Usage

### Basic Usage

```bash
go run main.go -input path/to/your/photo.jpg -output path/to/output.png
```

This will generate a new hairstyle that suits the person in the photo.

### Command Line Options

| Flag | Description | Default | Required |
|------|-------------|---------|----------|
| `-input` | Path to input image file | `./input.jpeg` | Yes |
| `-output` | Path to output image file | `./output.png` | Yes |
| `-same-person-prompt` | Custom hairstyle prompt | "give them a hairstyle that suits them the most" | No |
| `-full-prompt` | Complete custom prompt (overrides hairstyle mode) | "" | No |

### Examples

**Generate a suitable hairstyle:**
```bash
go run main.go -input selfie.jpg -output new_look.png
```

**Try a specific hairstyle:**
```bash
go run main.go -input selfie.jpg -output bob_cut.png -same-person-prompt "give them a short bob haircut"
```

**Use a completely custom prompt:**
```bash
go run main.go -input selfie.jpg -output artistic.png -full-prompt "Transform this person into a Renaissance painting style portrait"
```

**Generate multiple variations:**
```bash
go run main.go -input selfie.jpg -output variation1.png -same-person-prompt "give them long curly hair"
go run main.go -input selfie.jpg -output variation2.png -same-person-prompt "give them a modern pixie cut"
go run main.go -input selfie.jpg -output variation3.png -same-person-prompt "give them a professional business hairstyle"
```

## How It Works

1. **Input Processing**: The tool reads your input image file
2. **Prompt Construction**: Creates an AI prompt that preserves facial identity while requesting changes
3. **AI Generation**: Sends the image and prompt to Google's Gemini AI model
4. **Output**: Saves the generated image to your specified output path

The tool automatically adds context to ensure the person's facial features remain recognizable in the generated image.

## Supported Image Formats

- **Input**: JPEG files (other formats may work but are not officially supported)
- **Output**: PNG format

## Troubleshooting

**"GEMINI_API_KEY environment variable is required"**
- Make sure you've set the API key environment variable correctly
- Verify the key is valid and active

**"failed to read input image file"**
- Check that the input file path is correct
- Ensure the file exists and is readable
- Verify the image format is supported

**"no candidates returned from Gemini AI"**
- The AI may have filtered the request due to content policies
- Try a different image or modify your prompt
- Ensure your API key has sufficient quota

**Image quality issues**
- Use high-resolution input images for better results
- Ensure good lighting and clear facial features in the input
- Try different prompt variations

## API Costs

This tool uses Google's Gemini AI API. Check [Google's pricing page](https://ai.google.dev/pricing) for current rates. The free tier includes generous limits for personal use.

## License

This project uses the Google Gemini AI API. Please review Google's terms of service for API usage.
