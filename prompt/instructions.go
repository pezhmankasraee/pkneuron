package prompt

import (
	"strings"

	"github.com/google/uuid"
)

func Createprompt(userQuery string) string {

	var sb strings.Builder

	sb.WriteString("Act as structured data provider. Answer the following: " + userQuery)
	sb.WriteString("Return a raw JSON object only with this schema:")
	sb.WriteString("1. \"schema_version\":\"1.0.0\"")
	id := uuid.New()
	uu := id.String()
	sb.WriteString("\"request_id\":\"" + uu + "\"")
	sb.WriteString("\"timestamp\":Current UTC ISO 8601 string")
	sb.WriteString("\"request\":(request should be rephrased, string, max 250 words)")
	sb.WriteString("\"introduction\":(introduction to your response, string, max 250 words)")
	sb.WriteString("\"description\":(elaboration of your response in details, string, max 1200 words)")
	sb.WriteString("\"summary\":(summary of your description part, string, max 250 words)")
	sb.WriteString("\"keywords\":(list of keywords used in description, array, max 15 words)")
	sb.WriteString("Constraint: No markdown, no conversational filler, just the JSON")

	return sb.String()
}
