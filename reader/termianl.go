package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[33mPaste your text (press Ctrl+D to finish):\033[0m")

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil { // This will be true on EOF (Ctrl+D)
			break
		}
		line = strings.TrimRight(line, "\r\n")
		lines = append(lines, line)
	}
	input := strings.Join(lines, "\n")
	return input
}
