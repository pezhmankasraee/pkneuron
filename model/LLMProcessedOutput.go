package model

// LLMProcessedOutput represents the structured output from an LLM process.
type LLMProcessedOutput struct {
	Response        string   // LLM full output without any manipulation
	RephrasedPrompt string   // Paraphrased prompt or question, relevant to the LLM output
	Keywords        []string // Crucial terms from response and rephrase_prompt (max 10)
	Summary         string   // Concise summary of LLM output (max 1000 chars)
}

var LLMProcessedOutputJsonSchema = `
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "LLM Processed Output Structure",
  "description": "Schema defining the structure for an LLM's full output, a rephrased prompt, extracted keywords, and a concise summary.",
  "type": "object",
  "properties": {
    "response": {
      "type": "string",
      "description": "LLM full output without any manipulation"
    },
    "rephrased_prompt": {
      "type": "string",
      "description": "Paraphrased prompt or question, made as relevant as possible to the LLM full output in the 'response' field."
    },
    "keywords": {
      "type": "array",
      "description": "Crucial terms selected carefully from the 'response' and 'rephrase_prompt'. Keywords should be precise, enabling reconstruction of the response's gist or providing a clear understanding of its core topics. Elements can be common terms, symbols, etc.",
      "items": {
        "type": "string"
      },
      "minItems": 1,
      "maxItems": 10
    },
    "summary": {
      "type": "string",
      "description": "A concise summary of the LLM's full output ('response'), with a maximum length of 1000 characters. The summary should be as close as possible to the original response and related keywords, capturing the core gist and remaining highly focused.",
      "maxLength": 1000
    }
  },
  "required": [
    "response",
    "rephrase_prompt",
    "keywords",
    "summary"
  ],
  "additionalProperties": false
}
`

var LLMProcessedOutputPKNeuronFormat = `
	[[
	{{PKNuron:response
	 PKNuron:response}}

	 {{PKNuron:rephrased_prompt
	 PKNuron:rephrased_prompt}}

	 {{PKNuron:keywords
	 PKNuron:keywords}}

	 {{PKNuron:summary
	 PKNuron:summary}}
	 ]]
`
