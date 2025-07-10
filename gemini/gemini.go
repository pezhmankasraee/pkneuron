package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
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
		fmt.Println("\033[31me:Nothing to process\033[0m")
	} else {
		fmt.Println("\033[36mPlease wait ... \033[0m")

		model := "gemini-2.5-flash"
		//model := "learnlm-2.0-flash-experimental"

		result, err := client.Models.GenerateContent(
			ctx,
			model,
			genai.Text(request),
			nil,
		)
		if err != nil {
			pklog.CreateLog(pklog.FatalError, err.Error())
		}

		clean := cleanResponse(result)
		fmt.Println(clean)
	}
}
