package help

import (
	"os"
)

var commandHelp string = `
Commands / Special Markers:
    All commands start with 'cc::'.
    For example: >>c:s

    cc::e, >>cc::exit               exits the application (it can be done with Ctrl-c, too)
    cc::h, >>cc::help               shows commands in the request prompt
    cc::s, cc::send               sends your request to AI model

`

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
`

	visit := "For more information, visit: https://github.com/pezhmankasraee/pkneuron"

	println(helpMenu + commandHelp + visit)
}

func ShowCommands() {

	println(commandHelp)
}
