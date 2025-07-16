package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/command"
	"github.com/pezhmankasraee/pkneuron/errormessage"
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

	if strings.TrimSpace(request) == "" {
		fmt.Println("\033[31m" + errormessage.ERROR_00 + "\033[0m")
	} else {
		if request != command.LONG_HELP {

			fmt.Println("\033[36mPlease wait ... \033[0m")

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

			clean := cleanResponse(result)
			fmt.Println(clean)
		}
	}
}
