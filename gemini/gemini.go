package gemini

import (
	"context"
	"fmt"

	"github.com/pezhmankasraee/pklog/v2"
	"google.golang.org/genai"
)

func Init() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}

	//model := "gemini-2.5-flash"
	model := "learnlm-2.0-flash-experimental"
	result, _ := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text("Explain how AI works in a few words"),
		nil,
	)

	fmt.Println(result.Text())
}
