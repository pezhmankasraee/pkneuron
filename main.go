package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/gemini"
	"github.com/pezhmankasraee/pkneuron/help"
)

var isHelp bool

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)

	fmt.Println("-- Welcome to PKNeuron!")

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("\033[33;1mDo you want to continue (Y/n)?\033[0m")
		continueLoop, err := reader.ReadString('\n')
		if err != nil {
			pklog.CreateLog(pklog.FatalError, err.Error())
			continue
		}

		continueLoop = strings.TrimSpace(continueLoop) // remove spaces/newlines
		if len(continueLoop) > 0 && continueLoop[0] == 'n' {
			fmt.Println("Exiting ...")
			os.Exit(0)
		}

		gemini.Init()
	}

}

func init() {
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
}
