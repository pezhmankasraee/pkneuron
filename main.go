package main

import (
	"flag"
	"fmt"

	"github.com/pezhmankasraee/pkneuron/gemini"
	"github.com/pezhmankasraee/pkneuron/help"
)

var isHelp bool

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)

	fmt.Println("-- Welcome to PKNeuron!")

	for {
		gemini.Init()
	}

}

func init() {
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
}
