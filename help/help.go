package help

import (
	"os"
)

func ShowHelp(isHelp bool) {
	if isHelp {
		generalHelp()
		os.Exit(0)
	}
}

func generalHelp() {
	helpMenu := `
PKNeuron - An interactive LLM assistant for Linux terminal.

Usage: pkneuron [OPTIONS]

Options:
    -h, --help          Show this help message

Commands / Special Markers:
    All commands start and end with 'c:'.
    For example: c:s

    c:s, c:send               sends your request to AI model

For more information, visit: https://github.com/pezhmankasraee/pkneuron
`

	println(helpMenu)
}
