package functions

import (
	"strings"
	"errors"
	"regexp"
	"sync"
	"unicode"
	"unicode/utf8"
	"html"
)

// Define constants for chunk size and maximum concurrent goroutines
const chunkSize = 4096
const maxConcurrent = 100

// Sanitize a chunk of text by removing control characters and specified characters
func sanitizeChunk(chunk string, resultChan chan<- string, wg *sync.WaitGroup, limiter chan struct{}, charsToRemove []rune) {
	defer wg.Done()
	// Limit the number of concurrent goroutines
	limiter <- struct{}{}
	defer func() { <-limiter }()
	var sanitized strings.Builder
	// Iterate over each character in the chunk
	chunk = html.EscapeString(chunk)
	for _, char := range chunk {
		// Skip control characters (except newlines), specified characters, and invalid runes
		if (unicode.IsControl(char) && char != '\n' && char != '\t') || contains(charsToRemove, char) || !utf8.ValidRune(char) {
			continue
		}
		// Append valid characters to the sanitized builder
		sanitized.WriteRune(char)
	}
	// Send the sanitized string to the result channel
	resultChan <- sanitized.String()
}

// Check if the given character is in the list of characters to remove
func contains(list []rune, char rune) bool {
	for _, c := range list {
		if c == char {
			return true
		}
	}
	return false
}

// Sanitize the entire input string by processing it in chunks
func SafeString(input string, charsToRemove []rune) string {
	// Trim whitespace from the input
	input = strings.TrimSpace(input)
	lenInput := len(input)
	// Calculate the number of chunks needed
	numChunks := (lenInput + chunkSize - 1) / chunkSize

	resultChan := make(chan string, numChunks)
	var wg sync.WaitGroup
	// Create a limiter for concurrent goroutines
	limiter := make(chan struct{}, maxConcurrent)

	// Process the input string in chunks
	for i := 0; i < lenInput; i += chunkSize {
		end := i + chunkSize
		if end > lenInput {
			end = lenInput
		}
		wg.Add(1)
		// Launch a goroutine to sanitize each chunk
		go sanitizeChunk(input[i:end], resultChan, &wg, limiter, charsToRemove)
	}

	// Wait for all sanitization goroutines to finish and close the result channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var sanitized strings.Builder
	// Combine sanitized chunks into a single string
	for chunk := range resultChan {
		sanitized.WriteString(chunk)
	}

	// Return the final sanitized string
	return sanitized.String()
}

var SanitizeInputRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9@._-]+$`) ;
func SanitizeInput(input string) (string, error) {
    //only chars , numbers , @ , .
    isValid := SanitizeInputRegex.MatchString(input) ;
    if !isValid{
        return "", errors.New("invalid input detected")
    }
    return input, nil
}

func CompressString(s string)string{
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", ""), "\t", ""), "\r", ""), "\r\n" , "") ;
}
