package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/command"
	"github.com/pezhmankasraee/pkneuron/errormessage"
	"github.com/pezhmankasraee/pkneuron/prompt"
	"github.com/pezhmankasraee/pkneuron/reader"
	"google.golang.org/genai"
)

func Init() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}

	request := reader.Read()
	request = prompt.Createprompt(request)

	if strings.TrimSpace(request) == "" {
		fmt.Println("\033[31m" + errormessage.ERROR_00 + "\033[0m")
	} else {
		if request != command.LONG_HELP {

			fmt.Println("Please wait ...")
			fmt.Println()

			model := "gemini-2.5-flash"
			//model := "learnlm-2.0-flash-experimental"

			config := &genai.GenerateContentConfig{
				ResponseMIMEType: "text/plain",
			}
			result, err := client.Models.GenerateContent(
				ctx,
				model,
				genai.Text(request),
				config,
			)
			if err != nil {
				pklog.CreateLog(pklog.FatalError, err.Error())
			}

			clean := CleanResponse(result)

			regexpResponse := createRegExp(prompt.ResponseStartSha512, prompt.ResponseEndSha512)
			response := ExtractPKNeuronTag(clean, regexpResponse)
			fmt.Println("\033[36mResponse> \033[0m")
			fmt.Println(response)

			regexpRephrased := createRegExp(prompt.RephrasedStartSha512, prompt.RephrasedEndSha512)
			rephrased := ExtractPKNeuronTag(clean, regexpRephrased)
			fmt.Println("\033[36mRephrase> \033[0m")
			fmt.Println(rephrased)

			regexpKeywords := createRegExp(prompt.KeywordsStartSha512, prompt.KeywordsEndSha512)
			keywords := ExtractPKNeuronTag(clean, regexpKeywords)
			fmt.Println("\033[36mKeywords> \033[0m")
			fmt.Println(keywords)

			regexpSummary := createRegExp(prompt.SummaryStartSha512, prompt.SummaryEndSha512)
			summary := ExtractPKNeuronTag(clean, regexpSummary)
			fmt.Println("\033[36mSummary> \033[0m")
			fmt.Println(summary)

			numberOfTokens := wordCount(clean)
			fmt.Println("\033[36mNumber of tokens> \033[0m")
			fmt.Println(numberOfTokens)
		}
	}
}

func wordCount(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func createRegExp(startTagSha512 string, endTagSha512 string) string {
	return `(?s)` + startTagSha512 + `(.*?)` + endTagSha512
}
