package help

import (
	"fmt"
	"os"
)

func ShowHelp(isHelp bool) {
	if isHelp {
		generalHelp()
		os.Exit(0)
	}
}

func generalHelp() {
	fmt.Println("")
	fmt.Println("PKNeuron - An interactive LLM assistant for Linux terminal.")
	fmt.Println("")
	fmt.Println("Usage: pkneuron [OPTIONS]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help          Show this help message")
	fmt.Println("")
	fmt.Println("For more information, visit: https://github.com/pezhmankasraee/pkneuron")
}
