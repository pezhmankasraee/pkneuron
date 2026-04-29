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
