package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
)

func Read() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[33;1mNOTE: use c:s / c:send at the end of the request)\033[0m")
	fmt.Println("\033[33mrequest > \033[0m")

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

		if strings.EqualFold(line, "c:e") || strings.EqualFold(line, "c:exit") {
			fmt.Println("Exiting ...")
			os.Exit(0)
		}

		if strings.EqualFold(line, "c:send") || strings.EqualFold(line, "c:s") {
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
