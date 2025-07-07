package main

import (
	"flag"

	"github.com/pezhmankasraee/pkneuron/help"
)

var isHelp bool

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)
}

func init() {
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
}
