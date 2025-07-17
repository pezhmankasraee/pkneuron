package gemini

import (
	"regexp"

	"google.golang.org/genai"
)

func CleanResponse(response *genai.GenerateContentResponse) string {
	re := regexp.MustCompile("`{1,3}.*?`{1,3}")
	clean := re.ReplaceAllString(response.Text(), "")

	return clean
}

// extractResponse finds and returns the string content that exists between the specific
// delimiters within the input text.
// It uses regular expressions for pattern matching and returns the first captured string.
// If no matching content is found, an empty string is returned.
//
// example:
// FindStringSubmatch returns a slice of strings:
func ExtractPKNeuronTag(response string, regularExpresion string) string {

	re := regexp.MustCompile(regularExpresion)

	matches := re.FindStringSubmatch(response)

	// Check if a match was found and if the content capture group exists.
	if len(matches) > 1 {
		// Return the content from the first capturing group.
		return matches[1]
	}

	// If no match is found, return an empty string.
	return ""

}
