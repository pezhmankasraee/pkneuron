package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/command"
	"github.com/pezhmankasraee/pkneuron/command/utility/color"
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

			fmt.Print("\033[31m" + clean + color.GRAY)

			fmt.Println()
			fmt.Println("----------------------")

			response := schema.Convert(clean)

			fmt.Println("")
			fmt.Println(color.BLUE + "Q U E S T I O N" + color.GRAY)
			fmt.Println(response.Request)

			fmt.Println(color.BLUE + "I N T R O D U C T I O N" + color.GRAY)
			fmt.Println(response.Introduction)

			fmt.Println(color.BLUE + "D E S C R I P T I O N" + color.GRAY)
			fmt.Println(response.Description)

			fmt.Println(color.BLUE + "S U M M A R Y" + color.GRAY)
			fmt.Println(response.Summary)

			fmt.Println(color.BLUE + "K E Y W O R D S" + color.GRAY)
			for i := 0; i < len(response.Keywords); i++ {
				fmt.Print(response.Keywords[i])
				fmt.Print(", ")
			}

			fmt.Println()
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
