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
	"github.com/pezhmankasraee/pkneuron/schema"
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

			fmt.Print(clean)

			fmt.Println("----------------------")

			response := schema.Convert(clean)

			fmt.Println("")
			fmt.Println("\033[96m" + "Q U E S T I O N" + "\033[0m")
			fmt.Println(response.Request)

			fmt.Println("\033[96m" + "I N T R O D U C T I O N" + "\033[0m")
			fmt.Println(response.Introduction)

			fmt.Println("\033[96m" + "D E S C R I P T I O N" + "\033[0m")
			fmt.Println(response.Description)

			fmt.Println("\033[96m" + "S U M M A R Y" + "\033[0m")
			fmt.Println(response.Summary)

			fmt.Println("\033[96m]" + "K E Y W O R D S" + "\033[0m]")
			for i := 0; i < len(response.Keywords); i++ {
				fmt.Print(response.Keywords[i])
				fmt.Print(", ")
			}
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
