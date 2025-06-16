package functions

import (
	"strings"
)

// Check if the given file name has a valid extension from the provided list
func HasValidExtension(name string, validExtensions map[string]struct{}) bool {
	// Iterate through each valid extension
	for ext, _ := range validExtensions {
		// Check if the file name ends with the valid extension (case insensitive)
		if strings.HasSuffix(strings.ToLower(name), strings.ToLower(ext)) {
			return true
		}
	}
	// Return false if no valid extension is found
	return false
}
