package prompt

import (
	"github.com/google/uuid"
)

func Createprompt(userQuery string) string {

	jsonSchema := `
		req_id : ` + uuid.New().String() + `
		Act as structured data provider. Answer the following request: ` + userQuery +
		`
		Return a raw JSON object only with this schema:

		{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"title": "Response_v1",
		"type": "object",
		"properties": {
			"metadata": {
			"type": "object",
			"description": "Administrative headers for versioning and traceability.",
			"properties": {
				"version": {
				"type": "string",
				"const": "1.0.0",
				"description": "The fixed schema version. Must always be 1.0.0."
				},
				"request_id": {
				"type": "string",
				"pattern": "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$",
				"description": "The unique UUID provided in the user's request (req_id). Must be mirrored exactly."
				},
				"timestamp": {
				"type": "string",
				"format": "date-time",
				"description": "The current UTC time in ISO 8601 format."
				}
			},
			"required": ["version", "request_id", "timestamp"]
			},
			"content": {
			"type": "object",
			"description": "The generated knowledge payload.",
			"properties": {
				"introduction": {
				"type": "string",
				"maxLength": 2096,
				"description": "A high-level conceptual overview and paraphrased request, Maximum 2096 words."
				},
				"elaboration": {
				"type": "string",
				"maxLength": 7200,
				"description": "Detailed technical analysis and detailed explanation of your response including patterns and implementation details. Maximum 7200 words."
				},
				"summary": {
				"type": "string",
				"maxLength": 1500,
				"description": "Key takeaways and final conclusions based on elaboration section. Maximum 2096 words."
				},
				"keywords": {
				"type": "array",
				"items": {
					"type": "string",
					"maxLength": 30
				},
				"maxItems": 15,
				"description": "A list of up to 15 technical terms extracted from the elaboration section."
				}
			},
			"required": ["introduction", "elaboration", "summary", "keywords"]
			}
		},
		"required": ["metadata", "content"]
		}


		Constraint: No markdown, No conversational filler, Respond in plain text. Do not use backticks or other markup
		Quality: make sure the json is a valid json
		`

	return jsonSchema
}
