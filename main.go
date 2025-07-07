package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/gemini"
	"github.com/pezhmankasraee/pkneuron/help"
)

var isHelp bool

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)

	fmt.Println("-- Welcome to PKNeuron!")

	var continueLoop rune
	for {

		fmt.Println("\033[33;1mDo you want to continue (Y/n)?\033[0m")
		_, err := fmt.Scanf("%c", &continueLoop)
		if err != nil {
			pklog.CreateLog(pklog.FatalError, err.Error())
		}

		if continueLoop == 'n' {
			os.Exit(0)
		}

		gemini.Init()
	}

}

func init() {
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
}
