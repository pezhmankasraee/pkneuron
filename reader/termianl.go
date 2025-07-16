package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pkneuron/command"
	"github.com/pezhmankasraee/pkneuron/help"
)

func Read() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[33;1mNOTE: \033[0m")
	fmt.Println("\033[33;1m-  use " + command.SHORT_SEND + " / " + command.LONG_SEND + " at the end of the request\033[0m")
	fmt.Println("\033[33;1m-  use " + command.SHORT_HELP + " / " + command.LONG_HELP + " to see list of commmands\033[0m")
	fmt.Println("\033[33mPKNeuron> \033[0m")

	var sb strings.Builder // Use strings.Builder for efficient concatenation
	// sb.Grow(initialCapacity) // Optional: If you can estimate total size, pre-allocate for even more efficiency

	firstLine := true
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			pklog.CreateLog(pklog.FatalError, err.Error())
			break
		}

		line = strings.TrimRight(line, "\r\n") // Trim both \r and \n

		if strings.EqualFold(line, command.SHORT_EXIT) || strings.EqualFold(line, command.LONG_EXIT) {
			fmt.Println("Exiting ...")
			os.Exit(0)
		}

		if strings.EqualFold(line, command.SHORT_SEND) || strings.EqualFold(line, command.LONG_SEND) {
			break
		}

		if strings.EqualFold(line, command.SHORT_HELP) || strings.EqualFold(line, command.LONG_HELP) {
			help.ShowCommands()
			sb.WriteString(command.LONG_HELP)
			break
		}

		if !firstLine {
			sb.WriteString("\n") // Add newline separator *before* appending subsequent lines
		}
		sb.WriteString(line)
		firstLine = false
	}

	return sb.String() // Get the final string from the builder
}
